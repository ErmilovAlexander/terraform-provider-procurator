---
page_title: "Procurator Provider"
description: "The Procurator provider manages virtualization resources in Procurator over gRPC APIs."
---

# Procurator Provider

The Procurator provider manages virtualization resources in Procurator over gRPC APIs.

The provider currently supports:
- virtual machines
- templates
- datastores
- datastore folders
- switches
- networks
- VM snapshots
- VM disk attachments
- VM network attachments
- VM datastore migration
- inventory data sources for host, VM, template, datastore, network, switch, NIC, and storage

## Example Usage

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

## Argument Reference

- `endpoint` - (Required) gRPC endpoint for Procurator core.
- `umbra_endpoint` - (Optional) gRPC endpoint for Umbra.
- `storage_endpoint` - (Optional) gRPC endpoint for storage service.
- `token` - (Optional) Bearer token used for authentication.
- `username` - (Optional) Username used for login flow if supported by backend.
- `password` - (Optional, Sensitive) Password used for login flow if supported by backend.
- `ca_file` - (Optional) Path to a PEM CA certificate file.
- `authority` - (Optional) TLS authority / server name override.
- `insecure` - (Optional) Disable TLS verification. Use only for development or controlled test environments.

## Authentication

The provider supports:
- bearer token authentication
- username/password login flow if implemented in backend
- TLS with `ca_file`
- embedded CA if the provider binary is built with it
- `insecure = true` for test environments

## Supported Resources

- `procurator_datastore`
- `procurator_datastore_lvm`
- `procurator_datastore_folder`
- `procurator_network`
- `procurator_switch`
- `procurator_vm`
- `procurator_vm_snapshot`
- `procurator_vm_disk_attachment`
- `procurator_vm_network_attachment`
- `procurator_vm_convert_to_template`
- `procurator_vm_datastore_migration`

## Supported Data Sources

- `procurator_host`
- `procurator_datastore`
- `procurator_vm`
- `procurator_template`
- `procurator_network`
- `procurator_networks`
- `procurator_switch`
- `procurator_switches`
- `procurator_nics`
- `procurator_storage_devices`
- `procurator_storage_adapters`
- `procurator_vm_snapshots`

## Notes

- Datastore creation is performed through Procurator core.
- Network and switch resources are managed through Umbra.
- Storage inventory is exposed through storage data sources.
- Asynchronous operations are handled through backend task flows and then reconciled into Terraform state.
