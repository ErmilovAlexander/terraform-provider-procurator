---
page_title: "Authentication"
subcategory: "Guides"
description: "Authentication and TLS configuration for the Procurator provider."
---

# Authentication

The Procurator provider communicates with Procurator gRPC endpoints over TLS.

## Token Authentication

```terraform
provider "procurator" {
  endpoint  = "10.10.102.22:3641"
  token     = var.token
  ca_file   = var.ca_file
  authority = var.authority
}
```

## Username/Password Authentication

```terraform
provider "procurator" {
  endpoint  = "10.10.102.22:3641"
  username  = var.username
  password  = var.password
  ca_file   = var.ca_file
  authority = var.authority
}
```

## Insecure Mode

```terraform
provider "procurator" {
  endpoint = "10.10.102.22:3641"
  token    = var.token
  insecure = true
}
```

`insecure = true` should only be used in development or controlled test environments.

## CA Resolution Behavior

The provider resolves trust configuration in this order:

1. `ca_file`, if specified.
2. Embedded CA, if the binary was built with embedded CA support.
3. `insecure = true`, if explicit TLS verification bypass is enabled.

If no CA source is available and `insecure = false`, provider initialization fails.
