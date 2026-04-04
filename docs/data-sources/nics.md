---
page_title: "procurator_nics Data Source - Procurator Provider"
description: "Returns physical or logical NIC inventory from Umbra."
---

# procurator_nics Data Source

Returns physical or logical NIC inventory from Umbra.

## Example Usage

```terraform
data "procurator_nics" "all" {}
```

## Attribute Reference

- `nics` - List of NIC objects including:
  - `id`
  - `name`
  - `adapter`
  - `pci_addr`
  - `driver`
  - `carrier`
  - `speed`
  - `duplex`
  - `networks`
  - `sr_iov`
  - `cdp`
  - `lldp`
  - `managed`
  - `switch_id`
  - `mac`
  - `state`
  - `errors`
