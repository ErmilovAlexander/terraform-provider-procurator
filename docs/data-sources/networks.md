---
page_title: "procurator_networks Data Source - Procurator Provider"
description: "Returns all visible networks."
---

# procurator_networks Data Source

Returns all visible networks.

## Example Usage

```terraform
data "procurator_networks" "all" {}
```

## Attribute Reference

- `networks` - List of networks with fields such as:
  - `id`
  - `name`
  - `vlan`
  - `switch_id`
  - `state`
  - `kind`
  - `net_bridge`
