---
page_title: "procurator_datastore_folder Resource - Procurator Provider"
description: "Creates or manages a folder inside a datastore."
---

# procurator_datastore_folder Resource

Creates or manages a folder inside a datastore.

## Example Usage

```terraform
resource "procurator_datastore_folder" "images" {
  path = "DATASTORE_ID:/images"
}
```

## Argument Reference

- `path` - (Required) Datastore path in the form `DATASTORE_ID:/path`.

## Attribute Reference

- `id` - Resource ID.
- `path` - Full datastore path.

## Notes

- This resource assumes the backend supports folder creation inside datastore inventory paths.
- Use datastore ID prefixes rather than friendly names for consistent behavior.
