---
page_title: "procurator_vm_disk_attachment Resource - Procurator Provider"
description: "Attaches an additional virtual disk to an existing VM."
---

# procurator_vm_disk_attachment Resource

Attaches an additional virtual disk to an existing VM.

## Example Usage

```terraform
resource "procurator_vm_disk_attachment" "example" {
  vm_id            = "VM_ID"
  size_gb          = 10
  storage_id       = "DATASTORE_ID"
  device_type      = "disk"
  bus              = "virtio"
  target           = "vdb"
  boot_order       = 2
  provision_type   = "thin"
  disk_mode        = "dependent"
  read_only        = false
  remove_on_detach = true
}
```

## Argument Reference

- `vm_id` - (Required) VM ID.
- `size_gb` - (Required) Disk size in GB.
- `storage_id` - (Required) Datastore ID.
- `device_type` - (Optional) Device type.
- `bus` - (Optional) Bus type.
- `target` - (Optional) Target device name.
- `boot_order` - (Optional) Boot order.
- `provision_type` - (Optional) Provisioning type.
- `disk_mode` - (Optional) Disk mode.
- `read_only` - (Optional) Read-only flag.
- `remove_on_detach` - (Optional) Remove disk when resource is destroyed.

## Attribute Reference

- `id` - Attachment ID.
