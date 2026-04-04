---
page_title: "procurator_vm Resource - Procurator Provider"
description: "Manages a virtual machine."
---

# procurator_vm Resource

Manages a virtual machine.

## Example Usage

### Create from Scratch

```terraform
resource "procurator_vm" "example" {
  name            = "vm-example"
  storage_id      = "DATASTORE_ID"
  power_state     = "stopped"
  vcpus           = 2
  max_vcpus       = 2
  core_per_socket = 1
  memory_size_mb  = 4096
  cpu_model       = "host-model"
  machine_type    = "pc-q35-6.2"

  disk_devices {
    bus            = "virtio"
    target         = "vda"
    size           = 30
    create         = true
    boot_order     = 1
    storage_id     = "DATASTORE_ID"
    provision_type = "thin"
    read_only      = false
    disk_mode      = "dependent"
    device_type    = "disk"
  }

  network_devices {
    network    = "VLAN106"
    model      = "virtio"
    boot_order = 0
    vlan       = 0
  }

  boot_options {
    firmware      = "efi"
    boot_delay_ms = 1000
    boot_menu     = true
  }

  guest_tools {
    enabled           = true
    synchronized_time = true
  }
}
```

### Deploy from Template

```terraform
resource "procurator_vm" "from_template" {
  name        = "vm-from-template"
  template_id = "TEMPLATE_ID"
  storage_id  = "DATASTORE_ID"
  power_state = "stopped"
}
```

## Argument Reference

- `name` - (Required) VM name.
- `storage_id` - (Optional) Target datastore ID.
- `template_id` - (Optional) Template ID for deployment from template.
- `power_state` - (Optional) Desired power state.
- `start` - (Optional) Legacy create-and-start flag if supported by backend.
- `vcpus` - (Optional) Number of vCPUs.
- `max_vcpus` - (Optional) Maximum number of vCPUs.
- `core_per_socket` - (Optional) Cores per socket.
- `memory_size_mb` - (Optional) Memory size in MB.
- `cpu_model` - (Optional) CPU model.
- `machine_type` - (Optional) Machine type.
- `cpu_hotplug` - (Optional) CPU hotplug flag.
- `memory_hotplug` - (Optional) Memory hotplug flag.
- `disk_devices` - (Optional) VM disk blocks.
- `network_devices` - (Optional) VM NIC blocks.
- `boot_options` - (Optional) Boot options block.
- `guest_tools` - (Optional) Guest tools block.

## Attribute Reference

- `id` - VM ID.
- `uuid` - VM UUID.
- `is_template` - Whether backend reports the object as a template.
- `guest_os_family` - Guest OS family.
- `guest_os_version` - Guest OS version.
- `compatibility` - Compatibility mode.
- `storage_folder` - VM storage folder.

## Notes

- VM creation and reconciliation depend on asynchronous backend task completion.
- Template deployments may inherit devices from the template even when Terraform config does not declare them explicitly.
