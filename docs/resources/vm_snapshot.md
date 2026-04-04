
---

## `docs/resources/vm_snapshot.md`

```md
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