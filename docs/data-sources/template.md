
---

## `docs/data-sources/template.md`

```md
---
page_title: "procurator_template Data Source - Procurator Provider"
description: "Finds a template by ID, UUID, or name."
---

# procurator_template Data Source

Finds a template by ID, UUID, or name.

## Example Usage

```terraform
data "procurator_template" "example" {
  name = "base-template"
}