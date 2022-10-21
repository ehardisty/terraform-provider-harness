---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "harness_platform_infrastructure Data Source - terraform-provider-harness"
subcategory: "Next Gen"
description: |-
  Data source for retrieving a Harness Infrastructure.
---

# harness_platform_infrastructure (Data Source)

Data source for retrieving a Harness Infrastructure.

## Example Usage

```terraform
data "harness_platform_infrastructure" "test" {
  identifier = "identifier"
  name       = "name"
  org_id     = "org_id"
  project_id = "project_id"
  env_id     = "env_id"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `env_id` (String) environment identifier.
- `org_id` (String) Unique identifier of the Organization.
- `project_id` (String) Unique identifier of the Project.

### Optional

- `identifier` (String) Unique identifier of the resource.
- `name` (String) Name of the resource.

### Read-Only

- `deployment_type` (String) Infrastructure deployment type. Valid values are KUBERNETES_DIRECT, KUBERNETES_GCP, SERVERLESS_AWS_LAMBDA, PDC, KUBERNETES_AZURE, SSH_WINRM_AZURE, SSH_WINRM_AWS, AZURE_WEB_APP, ECS, GITOPS, CUSTOM_DEPLOYMENT.
- `description` (String) Description of the resource.
- `id` (String) The ID of this resource.
- `tags` (Set of String) Tags to associate with the resource. Tags should be in the form `name:value`.
- `type` (String) Type of Infrastructure. Valid values are KUBERNETES_DIRECT, KUBERNETES_GCP, SERVERLESS_AWS_LAMBDA, PDC, KUBERNETES_AZURE, SSH_WINRM_AZURE, SSH_WINRM_AWS, AZURE_WEB_APP, ECS, GITOPS, CUSTOM_DEPLOYMENT.
- `yaml` (String) Infrastructure YAML

