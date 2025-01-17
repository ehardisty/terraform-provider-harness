---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "harness_platform_policy Resource - terraform-provider-harness"
subcategory: "Next Gen"
description: |-
  Resource for creating a Harness Policy.
---

# harness_platform_policy (Resource)

Resource for creating a Harness Policy.

## Example Usage

```terraform
resource "harness_platform_policy" "test" {
  identifier = harness_platform_policy.test.identifier
  name       = harness_platform_policy.test.name
  rego       = "package test"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `identifier` (String) Unique identifier of the resource.
- `name` (String) Name of the resource.
- `rego` (String) Rego code for the policy.

### Optional

- `description` (String) Description of the resource.
- `org_id` (String) Unique identifier of the organization.
- `project_id` (String) Unique identifier of the project.
- `tags` (Set of String) Tags to associate with the resource.

### Read-Only

- `id` (String) The ID of this resource.

## Import

Import is supported using the following syntax:

```shell
# Import using the organization id
terraform import harness_platform_policy.example <organization_id>
```
