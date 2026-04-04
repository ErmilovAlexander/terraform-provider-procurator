---
page_title: "procurator_network Resource - Procurator Provider"
description: "Manages a network in Umbra."
---

# procurator_network Resource

Manages a network in Umbra.

## Example Usage

```terraform
resource "procurator_network" "example" {
  name      = "prod-net"
  vlan      = 120
  switch_id = "uSwitch0"
}
```

## Argument Reference

- `name` - (Required) Network name.
- `vlan` - (Required) VLAN ID.
- `switch_id` - (Required) Switch ID where the network is created.

## Attribute Reference

- `id` - Network ID.
- `state` - Network state.
- `kind` - Network kind.
- `vms_count` - Number of attached VMs.
- `active_ports` - Number of active ports.
- `net_bridge` - Backing bridge.
- `errors` - Backend-reported errors.

## Notes

- The resource is managed through Umbra rather than Procurator core.
- Read-after-create behavior depends on backend visibility through Umbra network inventory.
