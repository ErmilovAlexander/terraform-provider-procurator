
---

## `docs/resources/switch.md`

```md
---
page_title: "procurator_switch Resource - Procurator Provider"
description: "Manages a virtual switch in Umbra."
---

# procurator_switch Resource

Manages a virtual switch in Umbra.

## Example Usage

```terraform
resource "procurator_switch" "example" {
  mtu = 1500

  nics = {
    active  = ["enp1s0"]
    standby = []
    unused  = []
    inherit = false
  }
}