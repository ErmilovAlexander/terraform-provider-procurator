---
page_title: "procurator_vm Data Source - Procurator Provider"
description: "Finds a virtual machine by ID, deployment name, UUID, or name."
---

# procurator_vm Data Source

Finds a virtual machine by ID, deployment name, UUID, or name.

## Example Usage

```terraform
data "procurator_vm" "example" {
  name = "vm-example"
}
```

## Argument Reference

- `id` - (Optional) VM ID.
- `deployment_name` - (Optional) Deployment name.
- `name` - (Optional) VM name.
- `uuid` - (Optional) VM UUID.

## Attribute Reference

- `id` - VM ID.
- `uuid` - VM UUID.
- `name` - VM name.
- `deployment_name` - Deployment name.
- `storage_id` - Datastore ID.
- `storage_folder` - VM storage folder.
- `power_state` - Power state.
- `guest_os_family` - Guest OS family.
- `guest_os_version` - Guest OS version.
- `machine_type` - Machine type.
- `vcpus` - vCPU count.
- `memory_size_mb` - Memory size.
- `disk_devices` - VM disks.
- `network_devices` - VM NICs.
