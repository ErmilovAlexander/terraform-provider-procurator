---
page_title: "procurator_vm_convert_to_template Resource - Procurator Provider"
description: "Converts a VM into a template."
---

# procurator_vm_convert_to_template Resource

Converts a VM into a template.

## Example Usage

```terraform
resource "procurator_vm_convert_to_template" "example" {
  vm_id = "VM_ID"
}
```

## Argument Reference

- `vm_id` - (Required) VM ID.

## Attribute Reference

- `id` - Template ID or resulting object ID.
- `name` - Template name.
- `uuid` - Template UUID.
- `storage_id` - Template datastore ID.
- `storage_folder` - Template storage folder.
- `machine_type` - Machine type.
- `memory_size_mb` - Template memory size.
- `vcpus` - Template vCPU count.
- `is_template` - True for successful conversion.
