---
page_title: "procurator_datastore Resource - Procurator Provider"
description: "Manages a datastore through Procurator core."
---

# procurator_datastore Resource

Manages a datastore through Procurator core.

## Example Usage

### LVM-like Datastore

```terraform
resource "procurator_datastore" "example" {
  name      = "data-01"
  type_code = 2

  devices = ["sdb"]
}
```

### NFS Datastore

```terraform
resource "procurator_datastore" "example" {
  name      = "nfs-01"
  type_code = 4
  server    = "10.10.10.20"
  folder    = "/exports/vmstore"
}
```

## Argument Reference

- `name` - (Required) Datastore name.
- `type_code` - (Required) Datastore type code.
- `devices` - (Optional) Device names for block-backed datastore creation.
- `server` - (Optional) NFS server.
- `folder` - (Optional) NFS export path or folder.

## Attribute Reference

- `id` - Datastore ID.
- `pool_name` - Backend pool name.
- `state` - Datastore state returned by backend.
- `status` - Datastore status returned by backend.
- `drive_type` - Device class if available.
- `capacity_mb` - Total capacity in MB.
- `free_mb` - Free capacity in MB.
- `used_mb` - Used capacity in MB.
- `provisioned_mb` - Provisioned capacity in MB.
- `thin_provisioning` - Whether thin provisioning is enabled.
- `access_mode` - Datastore access mode.

## Notes

- Datastore creation is performed through Procurator core.
- For block-backed creation, use device names returned by the storage inventory data source.
