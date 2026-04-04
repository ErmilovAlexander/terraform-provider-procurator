
---

## `docs/data-sources/datastore.md`

```md
---
page_title: "procurator_datastore Data Source - Procurator Provider"
description: "Finds a datastore by name or ID."
---

# procurator_datastore Data Source

Finds a datastore by name or ID.

## Example Usage

```terraform
data "procurator_datastore" "example" {
  name = "DEV-STOR-0"
}