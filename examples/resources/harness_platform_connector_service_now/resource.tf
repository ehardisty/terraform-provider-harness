# Credential type UsernamePassword
resource "harness_platform_connector_service_now" "test" {
  identifier  = "identifier"
  name        = "name"
  description = "test"
  tags        = ["foo:bar"]

  service_now_url    = "https://servicenow.com"
  delegate_selectors = ["harness-delegate"]
  auth {
    auth_type = "UsernamePassword"
    username_password {
      username     = "admin"
      password_ref = "account.password_ref"
    }
  }
}

# Credential type AdfsClientCredentialsWithCertificate
resource "harness_platform_connector_service_now" "test" {
  identifier  = "identifier"
  name        = "name"
  description = "test"
  tags        = ["foo:bar"]

  service_now_url    = "https://servicenow.com"
  delegate_selectors = ["harness-delegate"]
  auth {
    auth_type = "AdfsClientCredentialsWithCertificate"
    adfs {
      certificate_ref = "account.certificate_ref"
      private_key_ref = "account.private_key_ref}"
      client_id_ref   = "account.client_id_ref"
      resource_id_ref = "account.resource_id_ref"
      adfs_url        = "https://adfs_url.com"
    }
  }
}
