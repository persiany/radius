resource app 'radius.dev/Application@v1alpha3' = {
  name: 'azure-resources-container-httproute'

  resource frontend 'ContainerComponent' = {
    name: 'frontend'
    properties: {
      connections: {
        backend: {
          kind: 'Http'
          source: backend_http.id
        }
      }
      container: {
        image: 'rynowak/frontend:0.5.0-dev'
        env: {
          SERVICE__BACKEND__HOST: backend_http.properties.host
          SERVICE__BACKEND__PORT: string(backend_http.properties.port)
        }
      }
    }
  }

  resource backend_http 'HttpRoute' = {
    name: 'backend'
  }

  resource backend 'ContainerComponent' = {
    name: 'backend'
    properties: {
      container: {
        image: 'rynowak/backend:0.5.0-dev'
        ports: {
          web: {
            containerPort: 80
            provides: backend_http.id
          }
        }
        volumes:{
          'my-volume':{
            kind: 'ephemeral'
            mountPath:'/tmpfs'
            managedStore:'memory'
          }
        }
      }
    }
  }
}
