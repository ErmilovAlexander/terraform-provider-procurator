---
page_title: "procurator_template Data Source - Procurator Provider"
description: "Finds a template by ID, UUID, or name."
---

# procurator_template Data Source

Finds a template by ID, UUID, or name.

## Example Usage

```terraform
data "procurator_template" "example" {
  name = "base-template"
}
```

## Argument Reference

- `id` - (Optional) Template ID.
- `name` - (Optional) Template name.
- `uuid` - (Optional) Template UUID.

## Attribute Reference

- `id` - Template ID.
- `uuid` - Template UUID.
- `name` - Template name.
- `is_template` - Template flag.
- `storage_id` - Datastore ID.
- `storage_folder` - Template storage folder.
- `machine_type` - Machine type.
- `memory_size_mb` - Template memory.
- `vcpus` - Template vCPU count.
- `guest_os_family` - Guest OS family.
- `guest_os_version` - Guest OS version.
- `compatibility` - Compatibility mode.
