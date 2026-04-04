---
page_title: "procurator_vm_datastore_migration Resource - Procurator Provider"
description: "Migrates a VM to another datastore."
---

# procurator_vm_datastore_migration Resource

Migrates a VM to another datastore.

## Example Usage

### Migrate All VM Disks

```terraform
resource "procurator_vm_datastore_migration" "example" {
  vm_id               = "VM_ID"
  target_datastore_id = "DATASTORE_ID"
}
```

### Migrate Selected Disks

```terraform
resource "procurator_vm_datastore_migration" "example" {
  vm_id               = "VM_ID"
  target_datastore_id = "DATASTORE_ID"
  include_meta        = true

  disk_source_paths = [
    "old_ds:/vm-folder/disk1.sdk",
    "old_ds:/vm-folder/disk2.sdk",
  ]
}
```

## Argument Reference

- `vm_id` - (Required) VM ID.
- `target_datastore_id` - (Required) Target datastore ID.
- `include_meta` - (Optional) Include VM metadata in migration.
- `disk_source_paths` - (Optional) Explicit disk source paths. If omitted, provider migrates all VM disks.

## Attribute Reference

- `id` - Resource ID.
- `task_id` - Migration task ID.
- `source_datastore_id` - Source datastore ID.
- `final_datastore_id` - Final datastore ID after migration.
