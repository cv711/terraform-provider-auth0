resource "auth0_resource_server" "my_resource_server" {
  name        = "My Resource Server (Managed by Terraform)"
  identifier  = "my-resource-server-identifier"
  signing_alg = "RS256"

  token_lifetime                                  = 86400
  skip_consent_for_verifiable_first_party_clients = true
  enforce_policies                                = true

  scopes {
    value       = "read:something"
    description = "read something"
  }
}

resource "auth0_user" "my_user" {
  connection_name = "Username-Password-Authentication"
  user_id         = "auth0|1234567890"
  email           = "test@test.com"
  password        = "passpass$12$12"
  nickname        = "testnick"
  username        = "testnick"
  roles           = [auth0_role.my_role.id]
}

resource "auth0_role" "my_role" {
  name        = "My Role - (Managed by Terraform)"
  description = "Role Description..."

  permissions {
    resource_server_identifier = auth0_resource_server.my_resource_server.identifier
    name                       = "read:something"
  }
}
