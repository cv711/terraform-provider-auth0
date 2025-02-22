---
page_title: "Resource: auth0_attack_protection"
description: |-
  Auth0 can detect attacks and stop malicious attempts to access your application such as blocking traffic from certain IPs and displaying CAPTCHAs.
---

# Resource: auth0_attack_protection

Auth0 can detect attacks and stop malicious attempts to access your application such as blocking traffic from certain IPs and displaying CAPTCHAs.

## Example Usage

```terraform
resource "auth0_attack_protection" "my_protection" {
  suspicious_ip_throttling {
    enabled   = true
    shields   = ["admin_notification", "block"]
    allowlist = ["192.168.1.1"]

    pre_login {
      max_attempts = 100
      rate         = 864000
    }

    pre_user_registration {
      max_attempts = 50
      rate         = 1200
    }
  }

  brute_force_protection {
    allowlist    = ["127.0.0.1"]
    enabled      = true
    max_attempts = 5
    mode         = "count_per_identifier_and_ip"
    shields      = ["block", "user_notification"]
  }

  breached_password_detection {
    admin_notification_frequency = ["daily"]
    enabled                      = true
    method                       = "standard"
    shields                      = ["admin_notification", "block"]
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `breached_password_detection` (Block List, Max: 1) Breached password detection protects your applications from bad actors logging in with stolen credentials. (see [below for nested schema](#nestedblock--breached_password_detection))
- `brute_force_protection` (Block List, Max: 1) Brute-force protection safeguards against a single IP address attacking a single user account. (see [below for nested schema](#nestedblock--brute_force_protection))
- `suspicious_ip_throttling` (Block List, Max: 1) Suspicious IP throttling blocks traffic from any IP address that rapidly attempts too many logins or signups. (see [below for nested schema](#nestedblock--suspicious_ip_throttling))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--breached_password_detection"></a>
### Nested Schema for `breached_password_detection`

Optional:

- `admin_notification_frequency` (Set of String) When "admin_notification" is enabled, determines how often email notifications are sent. Possible values: `immediately`, `daily`, `weekly`, `monthly`.
- `enabled` (Boolean) Whether breached password detection is active.
- `method` (String) The subscription level for breached password detection methods. Use "enhanced" to enable Credential Guard. Possible values: `standard`, `enhanced`.
- `shields` (Set of String) Action to take when a breached password is detected.


<a id="nestedblock--brute_force_protection"></a>
### Nested Schema for `brute_force_protection`

Optional:

- `allowlist` (Set of String) List of trusted IP addresses that will not have attack protection enforced against them.
- `enabled` (Boolean) Whether brute force attack protections are active.
- `max_attempts` (Number) Maximum number of unsuccessful attempts. Only available on public tenants.
- `mode` (String) Determines whether the IP address is used when counting failed attempts. Possible values: `count_per_identifier_and_ip` or `count_per_identifier`.
- `shields` (Set of String) Action to take when a brute force protection threshold is violated. Possible values: `block`, `user_notification`


<a id="nestedblock--suspicious_ip_throttling"></a>
### Nested Schema for `suspicious_ip_throttling`

Optional:

- `allowlist` (Set of String) List of trusted IP addresses that will not have attack protection enforced against them.
- `enabled` (Boolean) Whether suspicious IP throttling attack protections are active.
- `pre_login` (Block List, Max: 1) Configuration options that apply before every login attempt. Only available on public tenants. (see [below for nested schema](#nestedblock--suspicious_ip_throttling--pre_login))
- `pre_user_registration` (Block List, Max: 1) Configuration options that apply before every user registration attempt. Only available on public tenants. (see [below for nested schema](#nestedblock--suspicious_ip_throttling--pre_user_registration))
- `shields` (Set of String) Action to take when a suspicious IP throttling threshold is violated. Possible values: `block`, `admin_notification`

<a id="nestedblock--suspicious_ip_throttling--pre_login"></a>
### Nested Schema for `suspicious_ip_throttling.pre_login`

Optional:

- `max_attempts` (Number) Total number of attempts allowed per day.
- `rate` (Number) Interval of time, given in milliseconds, at which new attempts are granted.


<a id="nestedblock--suspicious_ip_throttling--pre_user_registration"></a>
### Nested Schema for `suspicious_ip_throttling.pre_user_registration`

Optional:

- `max_attempts` (Number) Total number of attempts allowed.
- `rate` (Number) Interval of time, given in milliseconds, at which new attempts are granted.

## Import

Import is supported using the following syntax:

```shell
# As this is not a resource identifiable by an ID within the Auth0 Management API,
# attack_protection can be imported using a random string.
#
# We recommend [Version 4 UUID](https://www.uuidgenerator.net/version4)
#
# Example:
terraform import auth0_attack_protection.my_protection 24940d4b-4bd4-44e7-894e-f92e4de36a40
```
