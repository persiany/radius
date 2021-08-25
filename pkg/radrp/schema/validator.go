// ------------------------------------------------------------
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.
// ------------------------------------------------------------

package schema

import (
	"embed"
	"fmt"
	"log"
	"strings"

	"github.com/xeipuuv/gojsonschema"
)

// ValidationError represents a validation error.
type ValidationError struct {

	// Position contains the field position, e.g. (root),
	// (root).location, (root).properties.components.0
	//
	// It could be unset, in case the object was not valid JSON.
	Position string

	// Message contains the error message, e.g. "location is required".
	Message string

	// JSONError contains the parsing error if the provided document
	// wasn't valid JSON.
	JSONError error
}

// Validator validates a JSON blob.
type Validator interface {

	// ValidateJSON validates a JSON blob and returns all the errors.
	ValidateJSON(json []byte) []ValidationError
}

var (
	// The listing of files below has an ordering to them, because
	// each file may depend on one or more files on the preceding
	// lines.
	//go:embed common-types.json
	//go:embed traits.json
	//go:embed basic-component.json
	//go:embed azure-cosmosdb-sql.json
	//go:embed azure-cosmosdb-mongo.json
	//go:embed azure-keyvault.json
	//go:embed azure-servicebus.json
	//go:embed container.json
	//go:embed dapr-pubsub.json
	//go:embed dapr-state.json
	//go:embed mongodb.json
	//go:embed redis.json
	//go:embed rabbitmq.json
	//go:embed components.json
	//go:embed radius.json
	jsonFiles embed.FS

	applicationValidator = newValidator("ApplicationResource")
	componentValidator   = newValidator("/components.json#/definitions/ComponentResource")
	deploymentValidator  = newValidator("DeploymentResource")
	scopeValidator       = newValidator("ScopeResource")
)

type validator struct {
	TypeName string
	schema   *gojsonschema.Schema
}

// ValidateJSON implements Validator.
func (v *validator) ValidateJSON(json []byte) []ValidationError {
	documentLoader := gojsonschema.NewBytesLoader(json)
	result, err := v.schema.Validate(documentLoader)
	if err != nil {
		return invalidJSONError(err)
	}
	if result.Valid() {
		return nil
	}
	errSet := make(map[ValidationError]struct{})
	errs := []ValidationError{}
	for _, err := range result.Errors() {
		if isAggregateError(err) {
			// Aggregate errors (OneOf, AllOf, AnyOf, Not) are usually
			// derived from other errors, and only make sense when the
			// users understand the details of JSON schema file. For
			// general error messages we probably want to avoid
			// displaying these.
			continue
		}
		v := ValidationError{
			Position: err.Context().String(),
			Message:  err.Description(),
		}
		if _, existed := errSet[v]; !existed {
			errSet[v] = struct{}{}
			errs = append(errs, v)
		}
	}
	return errs
}

// ValidatorFor returns a Validator for the given type, based on the
// type name.
func ValidatorFor(obj interface{}) (Validator, error) {
	objT := fmt.Sprintf("%T", obj)

	for suffix, validator := range map[string]*validator{
		".Application": applicationValidator,
		".Component":   componentValidator,
		".Deployment":  deploymentValidator,
		".Scope":       scopeValidator,
	} {
		if strings.HasSuffix(objT, suffix) {
			return validator, nil
		}
	}
	return nil, fmt.Errorf("Can't find a JSON validator for type %s", objT)
}

func GetComponentValidator() Validator {
	return componentValidator
}

func newValidator(typeName string) *validator {
	loader := gojsonschema.NewSchemaLoader()
	files, _ := jsonFiles.ReadDir(".")
	for _, f := range files {
		data, err := jsonFiles.ReadFile(f.Name())
		if err != nil {
			log.Fatalf("Cannot read embedded file %s: %v", f.Name(), err)
		}
		fileLoader := gojsonschema.NewBytesLoader(data)
		if err = loader.AddSchema( /* url */ "/"+f.Name(), fileLoader); err != nil {
			log.Fatalf("Failed to parse JSON Schema from %s: %s", f.Name(), err)
		}
	}
	ref := fmt.Sprintf("/radius.json#/definitions/%s", typeName)
	if strings.HasPrefix(typeName, "/") { // Allowing absolute path.
		ref = typeName
	}
	schema, err := loader.Compile(gojsonschema.NewStringLoader(fmt.Sprintf(`{
	  "$schema": "http://json-schema.org/draft-04/schema#",
	  "type":    "object",
	  "$ref":    "%s"
	}`, ref)))
	if err != nil {
		log.Fatalf("Failed to parse JSON Schema %s", err)
	}
	return &validator{
		schema:   schema,
		TypeName: typeName,
	}
}

func isAggregateError(err gojsonschema.ResultError) bool {
	switch err.(type) {
	case *gojsonschema.NumberAnyOfError, *gojsonschema.NumberOneOfError, *gojsonschema.NumberAllOfError:
		return true
	}
	return false
}

func invalidJSONError(err error) []ValidationError {
	return []ValidationError{{
		Message:   "invalid JSON error",
		JSONError: err,
	}}
}

type AggregateValidationError struct {
	Details []ValidationError
}

func (v *AggregateValidationError) Error() string {
	var message strings.Builder
	fmt.Fprintln(&message, "failed validation(s):")
	for _, err := range v.Details {
		if err.JSONError != nil {
			// The given document isn't even JSON.
			fmt.Fprintf(&message, "- %s: %v\n", err.Message, err.JSONError)
		} else {
			fmt.Fprintf(&message, "- %s: %s\n", err.Position, err.Message)
		}
	}
	return message.String()
}
