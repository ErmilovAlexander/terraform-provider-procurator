---
page_title: "procurator_switches Data Source - Procurator Provider"
description: "Returns all visible switches."
---

# procurator_switches Data Source

Returns all visible switches.

## Example Usage

```terraform
data "procurator_switches" "all" {}
```

## Attribute Reference

- `switches` - List of switch objects including:
  - `id`
  - `mtu`
  - `state`
  - `networks`
  - `errors`
  - `nics`
