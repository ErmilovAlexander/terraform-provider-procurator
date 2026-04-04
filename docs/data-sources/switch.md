
---

## `docs/data-sources/switch.md`

```md
---
page_title: "procurator_switch Data Source - Procurator Provider"
description: "Finds a switch by ID."
---

# procurator_switch Data Source

Finds a switch by ID.

## Example Usage

```terraform
data "procurator_switch" "example" {
  id = "uSwitch0"
}