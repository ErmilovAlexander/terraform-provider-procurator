---
page_title: "Getting Started"
subcategory: "Guides"
description: "Basic provider configuration and first steps with the Procurator provider."
---

# Getting Started

This guide shows the basic steps required to configure the Procurator provider and validate connectivity.

## Example Provider Configuration

```terraform
terraform {
  required_providers {
    procurator = {
      source  = "ErmilovAlexander/procurator"
      version = "0.1.0"
    }
  }
}

provider "procurator" {
  endpoint         = "10.10.102.22:3641"
  umbra_endpoint   = "10.10.102.22:50051"
  storage_endpoint = "10.10.102.22:3642"

  token     = var.token
  ca_file   = var.ca_file
  authority = var.authority
}
```

## Example Inventory Discovery

```terraform
data "procurator_host" "current" {}

data "procurator_datastore" "main" {
  name = "DEV-STOR-0"
}

data "procurator_networks" "all" {}
```

## Example VM Creation

```terraform
resource "procurator_vm" "example" {
  name            = "vm-example"
  storage_id      = data.procurator_datastore.main.id
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
    storage_id     = data.procurator_datastore.main.id
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
}
```

## Suggested Validation Sequence

1. Confirm provider authentication with `data "procurator_host"`.
2. Resolve target datastore with `data "procurator_datastore"`.
3. Resolve target network with `data "procurator_network"` or `data "procurator_networks"`.
4. Apply a small VM configuration first.
5. Add optional resources such as snapshots, disk attachments, network attachments, or migrations afterwards.
