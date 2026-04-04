
---

## `docs/resources/network.md`

```md
---
page_title: "procurator_network Resource - Procurator Provider"
description: "Manages a network in Umbra."
---

# procurator_network Resource

Manages a network in Umbra.

## Example Usage

```terraform
resource "procurator_network" "example" {
  name      = "prod-net"
  vlan      = 120
  switch_id = "uSwitch0"
}