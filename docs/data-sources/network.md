
---

## `docs/data-sources/network.md`

```md
---
page_title: "procurator_network Data Source - Procurator Provider"
description: "Finds a network by ID or name."
---

# procurator_network Data Source

Finds a network by ID or name.

## Example Usage

```terraform
data "procurator_network" "example" {
  name = "VLAN106"
}