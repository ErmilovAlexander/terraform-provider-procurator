---
page_title: "procurator_host Data Source - Procurator Provider"
description: "Returns information about the current Procurator host."
---

# procurator_host Data Source

Returns information about the current Procurator host.

## Example Usage

```terraform
data "procurator_host" "current" {}
```

## Attribute Reference

- `id` - Host ID.
- `name` - Host name.
- `hostname` - Hostname.
- `uuid` - Host UUID.
- `vendor` - Vendor or CPU/vendor string returned by backend.
- `model` - Platform model.
- `version` - Version or BIOS/build string.
