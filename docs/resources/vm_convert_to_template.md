
---

## `docs/resources/vm_convert_to_template.md`

```md
---
page_title: "procurator_vm_convert_to_template Resource - Procurator Provider"
description: "Converts a VM into a template."
---

# procurator_vm_convert_to_template Resource

Converts a VM into a template.

## Example Usage

```terraform
resource "procurator_vm_convert_to_template" "example" {
  vm_id = "VM_ID"
}