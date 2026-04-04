---
page_title: "procurator_switch Resource - Procurator Provider"
description: "Manages a virtual switch in Umbra."
---

# procurator_switch Resource

Manages a virtual switch in Umbra.

## Example Usage

```terraform
resource "procurator_switch" "example" {
  mtu = 1500

  nics = {
    active  = ["enp1s0"]
    standby = []
    unused  = []
    inherit = false
  }
}
```

## Argument Reference

- `mtu` - (Required) Switch MTU.
- `nics` - (Optional) NIC configuration object.
- `nics.active` - (Optional) Active NICs.
- `nics.standby` - (Optional) Standby NICs.
- `nics.unused` - (Optional) Unused NICs.
- `nics.inherit` - (Optional) Whether configuration is inherited.

## Attribute Reference

- `id` - Switch ID.
- `state` - Switch state.
- `networks` - Networks attached to the switch.
- `errors` - Backend-reported errors.
- `nics.connected` - Connected NICs reported by backend.

## Notes

- `nics` is configured as an object attribute, not as nested repeated blocks.
- Empty lists and null values are semantically different for Terraform state reconciliation.
