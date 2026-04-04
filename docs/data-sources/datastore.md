---
page_title: "procurator_datastore Data Source - Procurator Provider"
description: "Finds a datastore by name or ID."
---

# procurator_datastore Data Source

Finds a datastore by name or ID.

## Example Usage

```terraform
data "procurator_datastore" "example" {
  name = "DEV-STOR-0"
}
```

## Argument Reference

- `id` - (Optional) Datastore ID.
- `name` - (Optional) Datastore name.

## Attribute Reference

- `id` - Datastore ID.
- `name` - Datastore name.
- `pool_name` - Backend pool name.
- `type_code` - Datastore type code.
- `state` - Datastore state.
- `status` - Datastore status.
- `drive_type` - Drive type.
- `capacity_mb` - Total capacity.
- `free_mb` - Free capacity.
- `used_mb` - Used capacity.
- `provisioned_mb` - Provisioned capacity.
- `thin_provisioning` - Thin provisioning flag.
- `access_mode` - Access mode.
