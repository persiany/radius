// ------------------------------------------------------------
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.
// ------------------------------------------------------------

package db

import (
	"context"
	"errors"
	"fmt"

	"github.com/Azure/radius/pkg/radlogger"
	"github.com/Azure/radius/pkg/radrp/resources"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ErrNotFound is an error returned when an item is not found in the database.
var ErrNotFound = errors.New("the item was not found")

// ErrConcurrency is an error returned when the item contains stale data and cannot be modified.
var ErrConcurrency = errors.New("the item has been changed")

// applicationsCollection represents the collection used to store applications in the db.
const applicationsCollection string = "applications"

// operationsCollection represents the collection used to store operations in the db.
const operationsCollection string = "operations"

// NewRadrpDB creates a new RadrpDB.
func NewRadrpDB(m *mongo.Database) RadrpDB {
	return radrpDB{
		db: m,
	}
}

//go:generate mockgen -destination=../../../mocks/mock_db.go -package=mocks github.com/Azure/radius/pkg/radrp/db RadrpDB

// RadrpDB is our database abstraction.
//
// Patch operations are an upsert operation. It creates or updates the entry. `true` will be returned for a new record.
type RadrpDB interface {
	ListApplicationsByResourceGroup(ctx context.Context, id resources.ResourceID) ([]Application, error)
	GetApplicationByID(ctx context.Context, id resources.ApplicationID) (*Application, error)
	PatchApplication(ctx context.Context, patch *ApplicationPatch) (bool, error)
	UpdateApplication(ctx context.Context, app *Application) (bool, error)
	DeleteApplicationByID(ctx context.Context, id resources.ApplicationID) error

	ListComponentsByApplicationID(ctx context.Context, id resources.ApplicationID) ([]Component, error)
	GetComponentByApplicationID(ctx context.Context, id resources.ApplicationID, name string) (*Component, error)
	PatchComponentByApplicationID(ctx context.Context, id resources.ApplicationID, name string, patch *Component) (bool, error)
	DeleteComponentByApplicationID(ctx context.Context, id resources.ApplicationID, name string) error

	ListDeploymentsByApplicationID(ctx context.Context, id resources.ApplicationID) ([]Deployment, error)
	GetDeploymentByApplicationID(ctx context.Context, id resources.ApplicationID, name string) (*Deployment, error)
	PatchDeploymentByApplicationID(ctx context.Context, id resources.ApplicationID, name string, patch *Deployment) (bool, error)
	DeleteDeploymentByApplicationID(ctx context.Context, id resources.ApplicationID, name string) error

	ListScopesByApplicationID(ctx context.Context, id resources.ApplicationID) ([]Scope, error)
	GetScopeByApplicationID(ctx context.Context, id resources.ApplicationID, name string) (*Scope, error)
	PatchScopeByApplicationID(ctx context.Context, id resources.ApplicationID, name string, patch *Scope) (bool, error)
	DeleteScopeByApplicationID(ctx context.Context, id resources.ApplicationID, name string) error

	GetOperationByID(ctx context.Context, id resources.ResourceID) (*Operation, error)
	PatchOperationByID(ctx context.Context, id resources.ResourceID, patch *Operation) (bool, error)
	DeleteOperationByID(ctx context.Context, id resources.ResourceID) error
}

type radrpDB struct {
	db *mongo.Database
}

// ListApplicationsByResourceGroup lists applications by (subscription, resource group).
func (d radrpDB) ListApplicationsByResourceGroup(ctx context.Context, id resources.ResourceID) ([]Application, error) {
	items := make([]Application, 0)

	filter := bson.D{{Key: "subscriptionId", Value: id.SubscriptionID}, {Key: "resourceGroup", Value: id.ResourceGroup}}
	logger := radlogger.GetLogger(ctx)
	logger.Info(fmt.Sprintf("Listing applications from DB with operation filter: %s", filter))
	col := d.db.Collection(applicationsCollection)
	cursor, err := col.Find(ctx, filter)
	if err != nil {
		return items, fmt.Errorf("error querying Applications: %w", err)
	}

	err = cursor.All(ctx, &items)
	if err != nil {
		return items, fmt.Errorf("error reading Applications: %w", err)
	}

	logger.Info(fmt.Sprintf("Found %d Applications", len(items)))
	return items, nil
}

// GetApplicationByID finds applications by fully-qualified resource id.
func (d radrpDB) GetApplicationByID(ctx context.Context, id resources.ApplicationID) (*Application, error) {
	item := &Application{}

	filter := bson.D{{Key: "_id", Value: id.ID}}
	logger := radlogger.GetLogger(ctx)
	logger.Info(fmt.Sprintf("Getting application from DB with operation filter: %s", filter))
	col := d.db.Collection(applicationsCollection)
	result := col.FindOne(ctx, filter)
	err := result.Err()
	if err == mongo.ErrNoDocuments {
		logger.Info("Application was not found.")
		return nil, ErrNotFound
	} else if err != nil {
		return nil, fmt.Errorf("error querying %v: %w", id, err)
	}

	logger.Info("Found application in DB")
	err = result.Decode(item)
	if err != nil {
		return nil, fmt.Errorf("error reading %v: %w", id, err)
	}

	return item, nil
}

func (d radrpDB) PatchApplication(ctx context.Context, patch *ApplicationPatch) (bool, error) {
	options := options.Update().SetUpsert(true)
	filter := bson.D{{Key: "_id", Value: patch.ResourceBase.ID}}
	logger := radlogger.GetLogger(ctx)
	update := bson.D{{Key: "$set", Value: patch}}

	logger.Info(fmt.Sprintf("Updating Application in DB with operation filter: %s", filter))
	col := d.db.Collection(applicationsCollection)
	result, err := col.UpdateOne(ctx, filter, update, options)
	if err != nil {
		return false, fmt.Errorf("error updating Application: %s", err)
	}

	logger.Info(fmt.Sprintf("Successfully updated Application in DB with result - %+v", result))
	return result.UpsertedCount > 0, nil
}

func (d radrpDB) UpdateApplication(ctx context.Context, app *Application) (bool, error) {
	options := options.Update().SetUpsert(true)
	filter := bson.D{{Key: "_id", Value: app.ResourceBase.ID}}
	logger := radlogger.GetLogger(ctx).WithValues(radlogger.LogFieldAppID, app.ResourceBase.ID)
	update := bson.D{{Key: "$set", Value: app}}

	logger.Info(fmt.Sprintf("Updating Application in DB with operation filter: %s", filter))
	col := d.db.Collection(applicationsCollection)
	result, err := col.UpdateOne(ctx, filter, update, options)
	if err != nil {
		return false, fmt.Errorf("error updating Application: %s", err)
	}
	logger.Info(fmt.Sprintf("Updated Application in DB - %+v", result))
	return (result.UpsertedCount > 0 || result.ModifiedCount > 0), nil
}

func (d radrpDB) DeleteApplicationByID(ctx context.Context, id resources.ApplicationID) error {
	filter := bson.D{{Key: "_id", Value: id.ID}}
	logger := radlogger.GetLogger(ctx).WithValues(radlogger.LogFieldAppID, id)
	logger.Info(fmt.Sprintf("Deleting Application from DB with operation filter: %s", filter))
	col := d.db.Collection(applicationsCollection)
	result := col.FindOneAndDelete(ctx, filter)
	err := result.Err()
	if err == mongo.ErrNoDocuments {
		return nil
	} else if err != nil {
		return fmt.Errorf("error deleting Application with _id: '%s': %w", id, err)
	}

	logger.Info("Deleted Application from DB")
	return nil
}

func (d radrpDB) ListComponentsByApplicationID(ctx context.Context, id resources.ApplicationID) ([]Component, error) {
	logger := radlogger.GetLogger(ctx).WithValues(radlogger.LogFieldAppID, id)
	logger.Info("Listing Components from DB")
	application, err := d.GetApplicationByID(ctx, id)
	if err != nil {
		return nil, err
	}

	items := make([]Component, 0, len(application.Components))
	for _, item := range application.Components {
		items = append(items, item)
	}

	logger.Info(fmt.Sprintf("Found %d components in DB", len(application.Components)))
	return items, nil
}

func (d radrpDB) GetComponentByApplicationID(ctx context.Context, id resources.ApplicationID, name string) (*Component, error) {
	logger := radlogger.GetLogger(ctx).WithValues(
		radlogger.LogFieldAppID, id,
		radlogger.LogFieldComponentName, name)
	logger.Info("Getting components")
	application, err := d.GetApplicationByID(ctx, id)
	if err != nil {
		return nil, err
	}

	item, ok := application.Components[name]
	if !ok {
		logger.Info("Failed to find components in DB")
		return nil, ErrNotFound
	}

	logger.Info(fmt.Sprintf("Found component with revision: %s in DB", item.Revision))
	return &item, nil
}

func (d radrpDB) PatchComponentByApplicationID(ctx context.Context, id resources.ApplicationID, name string, patch *Component) (bool, error) {
	options := options.Update().SetUpsert(true)
	key := fmt.Sprintf("components.%s", name)
	filter := bson.D{{Key: "_id", Value: id.ID}}
	logger := radlogger.GetLogger(ctx).WithValues(
		radlogger.LogFieldAppID, id,
		radlogger.LogFieldAppName, name)
	update := bson.D{{Key: "$set", Value: bson.D{{Key: key, Value: patch}}}}

	logger.Info(fmt.Sprintf("Updating component in DB with operation filter: %s", filter))
	col := d.db.Collection(applicationsCollection)
	result, err := col.UpdateOne(ctx, filter, update, options)
	if err != nil {
		return false, fmt.Errorf("error updating Component: %s", err)
	}

	logger.Info("Updated component in DB")

	return result.UpsertedCount > 1, nil
}

func (d radrpDB) DeleteComponentByApplicationID(ctx context.Context, id resources.ApplicationID, name string) error {
	options := options.Update().SetUpsert(true)
	key := fmt.Sprintf("components.%s", name)
	filter := bson.D{{Key: "_id", Value: id.ID}}
	logger := radlogger.GetLogger(ctx).WithValues(
		radlogger.LogFieldAppID, id,
		radlogger.LogFieldAppName, name)
	update := bson.D{{Key: "$unset", Value: bson.D{{Key: key, Value: ""}}}}

	logger.Info(fmt.Sprintf("Deleting component from DB with operation filter: %s", filter))
	col := d.db.Collection(applicationsCollection)
	result, err := col.UpdateOne(ctx, filter, update, options)
	if err != nil {
		return fmt.Errorf("error deleting Application: %s", err)
	}

	logger.Info(fmt.Sprintf("Deleted component in DB - %+v", result))
	return nil
}

func (d radrpDB) ListDeploymentsByApplicationID(ctx context.Context, id resources.ApplicationID) ([]Deployment, error) {
	application, err := d.GetApplicationByID(ctx, id)
	logger := radlogger.GetLogger(ctx).WithValues(
		radlogger.LogFieldAppID, id,
		radlogger.LogFieldAppName, application.Name)
	logger.Info("Getting deployments from DB")
	if err != nil {
		return nil, err
	}

	items := make([]Deployment, 0, len(application.Deployments))
	for _, v := range application.Deployments {
		items = append(items, v)
	}

	logger.Info(fmt.Sprintf("Found %d Deployments in DB", len(items)))
	return items, nil
}

func (d radrpDB) GetDeploymentByApplicationID(ctx context.Context, id resources.ApplicationID, name string) (*Deployment, error) {
	logger := radlogger.GetLogger(ctx).WithValues(
		radlogger.LogFieldAppID, id,
		radlogger.LogFieldDeploymentName, name)
	logger.Info("Getting deployment from DB")
	application, err := d.GetApplicationByID(ctx, id)
	if err != nil {
		return nil, err
	}

	item, ok := application.Deployments[name]
	if !ok {
		logger.Info("Failed to find deployment in DB")
		return nil, ErrNotFound
	}

	logger.Info("Found Deployment in DB")
	return &item, nil
}

func (d radrpDB) PatchDeploymentByApplicationID(ctx context.Context, id resources.ApplicationID, name string, patch *Deployment) (bool, error) {
	options := options.Update().SetUpsert(true)
	key := fmt.Sprintf("deployments.%s", name)
	filter := bson.D{{Key: "_id", Value: id.ID}}
	logger := radlogger.GetLogger(ctx).WithValues(
		radlogger.LogFieldAppID, id,
		radlogger.LogFieldDeploymentName, name)
	update := bson.D{{Key: "$set", Value: bson.D{{Key: key, Value: patch}}}}

	logger.Info(fmt.Sprintf("Updating deployment in DB with operation filter: %s", filter))
	col := d.db.Collection(applicationsCollection)
	result, err := col.UpdateOne(ctx, filter, update, options)
	if err != nil {
		return false, fmt.Errorf("error updating Application: %s", err)
	}

	logger.Info(fmt.Sprintf("Updated application in DB - %+v", result))
	return result.UpsertedCount > 1, nil
}

func (d radrpDB) DeleteDeploymentByApplicationID(ctx context.Context, id resources.ApplicationID, name string) error {
	options := options.Update().SetUpsert(true)
	key := fmt.Sprintf("deployments.%s", name)
	filter := bson.D{{Key: "_id", Value: id.ID}}
	logger := radlogger.GetLogger(ctx).WithValues(
		radlogger.LogFieldAppID, id,
		radlogger.LogFieldDeploymentName, name)
	update := bson.D{{Key: "$unset", Value: bson.D{{Key: key}}}}

	logger.Info(fmt.Sprintf("Deleting deployment from DB: %s with operation filter: %s", update, filter))
	col := d.db.Collection(applicationsCollection)
	result, err := col.UpdateOne(ctx, filter, update, options)
	if err != nil {
		return fmt.Errorf("error deleting Application: %s", err)
	}

	logger.Info(fmt.Sprintf("Deleted deployment from DB- %+v", result))
	return nil
}

func (d radrpDB) ListScopesByApplicationID(ctx context.Context, id resources.ApplicationID) ([]Scope, error) {
	application, err := d.GetApplicationByID(ctx, id)
	logger := radlogger.GetLogger(ctx).WithValues(
		radlogger.LogFieldAppID, id,
		radlogger.LogFieldAppName, application.Name)
	logger.Info("Getting Scopes from DB")
	if err != nil {
		return nil, err
	}

	items := make([]Scope, 0, len(application.Scopes))
	for _, v := range application.Scopes {
		items = append(items, v)
	}

	logger.Info(fmt.Sprintf("Found %d Scopes in DB", len(items)))
	return items, nil
}

func (d radrpDB) GetScopeByApplicationID(ctx context.Context, id resources.ApplicationID, name string) (*Scope, error) {
	application, err := d.GetApplicationByID(ctx, id)
	logger := radlogger.GetLogger(ctx).WithValues(
		radlogger.LogFieldAppID, id,
		radlogger.LogFieldAppName, application.Name,
		radlogger.LogFieldScopeName, name)
	logger.Info("Getting scope from DB")
	if err != nil {
		return nil, err
	}

	item, ok := application.Scopes[name]
	if !ok {
		logger.Info("Failed to find scope in DB")
		return nil, ErrNotFound
	}

	logger.Info("Found scope in DB")
	return &item, nil
}

func (d radrpDB) PatchScopeByApplicationID(ctx context.Context, id resources.ApplicationID, name string, patch *Scope) (bool, error) {
	options := options.Update().SetUpsert(true)
	key := fmt.Sprintf("scopes.%s", name)
	filter := bson.D{{Key: "_id", Value: id.ID}}
	logger := radlogger.GetLogger(ctx).WithValues(
		radlogger.LogFieldAppID, id,
		radlogger.LogFieldScopeName, name)
	update := bson.D{{Key: "$set", Value: bson.D{{Key: key, Value: patch}}}}

	logger.Info(fmt.Sprintf("Updating scope in DB with operation filter: %s", filter))
	col := d.db.Collection(applicationsCollection)
	result, err := col.UpdateOne(ctx, filter, update, options)
	if err != nil {
		return false, fmt.Errorf("error updating Scope: %s", err)
	}

	logger.Info(fmt.Sprintf("Updated scope in DB - %+v", result))
	return result.UpsertedCount > 1, nil
}

func (d radrpDB) DeleteScopeByApplicationID(ctx context.Context, id resources.ApplicationID, name string) error {
	options := options.Update().SetUpsert(true)
	key := fmt.Sprintf("scopes.%s", name)
	filter := bson.D{{Key: "_id", Value: id.ID}}
	logger := radlogger.GetLogger(ctx).WithValues(
		radlogger.LogFieldAppID, id,
		radlogger.LogFieldScopeName, name)
	update := bson.D{{Key: "$unset", Value: bson.D{{Key: key, Value: ""}}}}

	logger.Info(fmt.Sprintf("Deleting scope from DB with operation filter: %s", filter))
	col := d.db.Collection(applicationsCollection)
	result, err := col.UpdateOne(ctx, filter, update, options)
	if err != nil {
		return fmt.Errorf("error deleting Application: %s", err)
	}

	logger.Info(fmt.Sprintf("Deleted scope from DB - %+v", result))
	return nil
}

func (d radrpDB) GetOperationByID(ctx context.Context, id resources.ResourceID) (*Operation, error) {
	item := &Operation{}

	filter := bson.D{{Key: "_id", Value: id.ID}}
	logger := radlogger.GetLogger(ctx).WithValues(
		radlogger.LogFieldOperationID, id)
	logger.Info(fmt.Sprintf("Getting operation from DB with operation filter: %s", filter))
	col := d.db.Collection(operationsCollection)
	result := col.FindOne(ctx, filter)
	err := result.Err()
	if err == mongo.ErrNoDocuments {
		logger.Info("operation was not found.")
		return nil, ErrNotFound
	} else if err != nil {
		return nil, fmt.Errorf("error querying %v: %w", id, err)
	}

	logger.Info("Found operation in DB")
	err = result.Decode(item)
	if err != nil {
		return nil, fmt.Errorf("error reading %v: %w", id, err)
	}

	return item, nil
}

func (d radrpDB) PatchOperationByID(ctx context.Context, id resources.ResourceID, patch *Operation) (bool, error) {
	options := options.Update().SetUpsert(true)
	filter := bson.D{{Key: "_id", Value: id.ID}}
	logger := radlogger.GetLogger(ctx).WithValues(
		radlogger.LogFieldOperationID, id)
	update := bson.D{{Key: "$set", Value: patch}}

	logger.Info(fmt.Sprintf("Updating operation in DB with operation filter: %s", filter))
	col := d.db.Collection(operationsCollection)
	result, err := col.UpdateOne(ctx, filter, update, options)
	if err != nil {
		return false, fmt.Errorf("error updating Operation: %s", err)
	}

	logger.Info(fmt.Sprintf("Updated operation in DB - %+v", result))
	return result.UpsertedCount > 1, nil
}

func (d radrpDB) DeleteOperationByID(ctx context.Context, id resources.ResourceID) error {
	filter := bson.D{{Key: "_id", Value: id.ID}}
	logger := radlogger.GetLogger(ctx).WithValues(
		radlogger.LogFieldOperationID, id)
	logger.Info(fmt.Sprintf("Deleting operation from DB with operation filter: %s", filter))
	col := d.db.Collection(operationsCollection)
	result := col.FindOneAndDelete(ctx, filter)
	err := result.Err()
	if err == mongo.ErrNoDocuments {
		return nil
	} else if err != nil {
		return fmt.Errorf("error deleting Operation with _id: '%s': %w", id, err)
	}

	logger.Info("Deleted operation from DB")
	return nil
}