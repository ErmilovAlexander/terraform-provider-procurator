---
page_title: "procurator_storage_adapters Data Source - Procurator Provider"
description: "Returns storage adapter inventory."
---

# procurator_storage_adapters Data Source

Returns storage adapter inventory.

## Example Usage

```terraform
data "procurator_storage_adapters" "all" {}
```

## Attribute Reference

- `items` - List of storage adapter objects including:
  - `id`
  - `adapter`
  - `identifier`
  - `model`
  - `type`
  - `status_text`
  - `status_value`
  - `targets`
  - `devices`
  - `rescan_storage_adapter`
  - `scan_storage_device`
  - `scan_datastore`
  - `protect_datastore`
  - `response`
