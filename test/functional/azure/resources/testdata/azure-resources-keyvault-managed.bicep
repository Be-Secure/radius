resource app 'radius.dev/Applications@v1alpha1' = {
  name: 'azure-resources-keyvault-managed'

  resource kv 'Components' = {
    name: 'kv'
    kind: 'azure.com/KeyVault@v1alpha1'
    properties: {
      config: {
        managed: true
      }
    }
  }

  resource kvaccessor 'Components' = {
    name: 'kvaccessor'
    kind: 'radius.dev/Container@v1alpha1'
    properties: {
      run: {
        container: {
          image: 'radius.azurecr.io/magpie:latest'
        }
      }
      uses: [
        {
          binding: kv.properties.bindings.default
          env: {
            BINDING_KEYVAULT_URI: kv.properties.bindings.default.uri
          }
        }
      ]
    }
  }
}