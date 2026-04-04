
---

## `docs/resources/vm_network_attachment.md`

```md
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