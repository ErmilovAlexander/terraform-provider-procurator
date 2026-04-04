---
page_title: "procurator_vm_snapshot Resource - Procurator Provider"
description: "Creates and manages a VM snapshot."
---

# procurator_vm_snapshot Resource

Creates and manages a VM snapshot.

## Example Usage

```terraform
resource "procurator_vm_snapshot" "example" {
  vm_id          = "VM_ID"
  name           = "snap-01"
  description    = "snapshot created by terraform"
  include_memory = false
  quiesce_fs     = false
}
```

## Argument Reference

- `vm_id` - (Required) VM ID.
- `name` - (Required) Snapshot name.
- `description` - (Optional) Snapshot description.
- `include_memory` - (Optional) Include VM memory.
- `quiesce_fs` - (Optional) Quiesce guest filesystem.

## Attribute Reference

- `id` - Terraform resource ID.
- `snapshot_id` - Backend snapshot ID.
