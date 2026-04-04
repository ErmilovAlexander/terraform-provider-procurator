---
page_title: "procurator_storage_devices Data Source - Procurator Provider"
description: "Returns storage device inventory."
---

# procurator_storage_devices Data Source

Returns storage device inventory.

## Example Usage

```terraform
data "procurator_storage_devices" "all" {}
```

## Attribute Reference

- `items` - List of storage device objects including:
  - `id`
  - `name`
  - `identifier`
  - `adapter`
  - `lun`
  - `capacity_mb`
  - `drive_type`
  - `transport`
  - `datastore_id`
  - `datastore_name`
  - `datastore_type`
  - `sector_format`
  - `storage_interface`
  - `operational_state`
  - `owner`
  - `perennially_reserved`
