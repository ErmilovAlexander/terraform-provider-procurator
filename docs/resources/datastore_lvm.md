
---

## `docs/resources/datastore_lvm.md`

```md
---
page_title: "procurator_datastore_lvm Resource - Procurator Provider"
description: "Creates a datastore backed by block devices through Procurator core."
---

# procurator_datastore_lvm Resource

Creates a datastore backed by block devices through Procurator core.

## Example Usage

```terraform
resource "procurator_datastore_lvm" "example" {
  name = "fast-ssd"

  devices = [
    "sdb"
  ]
}