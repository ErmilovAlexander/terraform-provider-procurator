---
page_title: "procurator_vm_network_attachment Resource - Procurator Provider"
description: "Attaches an additional NIC to an existing VM."
---

# procurator_vm_network_attachment Resource

Attaches an additional NIC to an existing VM.

## Example Usage

```terraform
resource "procurator_vm_network_attachment" "example" {
  vm_id      = "VM_ID"
  network    = "VLAN106"
  target     = "eth1"
  model      = "virtio"
  boot_order = 10
  vlan       = 0
}
```

## Argument Reference

- `vm_id` - (Required) VM ID.
- `network` - (Required) Network name.
- `target` - (Optional) NIC target label.
- `model` - (Optional) NIC model.
- `boot_order` - (Optional) Boot order.
- `vlan` - (Optional) VLAN value.

## Attribute Reference

- `id` - Attachment ID.
- `mac` - Generated MAC address if returned by backend.
