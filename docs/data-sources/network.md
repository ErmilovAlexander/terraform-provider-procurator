---
page_title: "procurator_network Data Source - Procurator Provider"
description: "Finds a network by ID or name."
---

# procurator_network Data Source

Finds a network by ID or name.

## Example Usage

```terraform
data "procurator_network" "example" {
  name = "VLAN106"
}
```

## Argument Reference

- `id` - (Optional) Network ID.
- `name` - (Optional) Network name.

## Attribute Reference

- `id` - Network ID.
- `name` - Network name.
- `vlan` - VLAN ID.
- `switch_id` - Switch ID.
- `state` - Network state.
- `kind` - Network kind.
- `net_bridge` - Backing bridge.
- `vms_count` - VM count.
- `active_ports` - Active port count.
- `errors` - Backend errors.
