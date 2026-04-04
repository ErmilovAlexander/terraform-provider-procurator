
---

## `docs/data-sources/vm.md`

```md
---
page_title: "procurator_vm Data Source - Procurator Provider"
description: "Finds a virtual machine by ID, deployment name, UUID, or name."
---

# procurator_vm Data Source

Finds a virtual machine by ID, deployment name, UUID, or name.

## Example Usage

```terraform
data "procurator_vm" "example" {
  name = "vm-example"
}