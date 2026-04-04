
---

# RESOURCES

## `docs/resources/datastore.md`

```md
---
page_title: "procurator_datastore Resource - Procurator Provider"
description: "Manages a datastore through Procurator core."
---

# procurator_datastore Resource

Manages a datastore through Procurator core.

## Example Usage

### LVM-like datastore

```terraform
resource "procurator_datastore" "example" {
  name      = "data-01"
  type_code = 2

  devices = ["sdb"]
}