---
page_title: "procurator_vm_snapshots Data Source - Procurator Provider"
description: "Returns snapshots for a VM."
---

# procurator_vm_snapshots Data Source

Returns snapshots for a VM.

## Example Usage

```terraform
data "procurator_vm_snapshots" "example" {
  vm_id = "VM_ID"
}
```

## Argument Reference

- `vm_id` - (Required) VM ID.

## Attribute Reference

- `current_id` - Current snapshot ID.
- `items` - Snapshot list. Each item may include:
  - `id`
  - `name`
  - `description`
  - `timestamp`
  - `size`
  - `parent_id`
  - `quiesce_fs`
  - `vm_description`
  - `memory`
  - `disks`
