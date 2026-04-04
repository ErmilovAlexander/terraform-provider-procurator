---
page_title: "procurator_switch Data Source - Procurator Provider"
description: "Finds a switch by ID."
---

# procurator_switch Data Source

Finds a switch by ID.

## Example Usage

```terraform
data "procurator_switch" "example" {
  id = "uSwitch0"
}
```

## Argument Reference

- `id` - (Required) Switch ID.

## Attribute Reference

- `id` - Switch ID.
- `mtu` - Switch MTU.
- `state` - Switch state.
- `networks` - Attached networks.
- `errors` - Backend errors.
- `nics` - NIC configuration object.
