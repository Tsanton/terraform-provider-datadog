---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "datadog_restriction_policy Resource - terraform-provider-datadog"
subcategory: ""
description: |-
  Provides a Datadog RestrictionPolicy resource. This can be used to create and manage Datadog restriction policies.
---

# datadog_restriction_policy (Resource)

Provides a Datadog RestrictionPolicy resource. This can be used to create and manage Datadog restriction policies.

## Example Usage

```terraform
# Create new restriction_policy resource


resource "datadog_restriction_policy" "foo" {
  resource_id = "security-rule:abc-def-ghi"
  bindings {
    principals = ["role:00000000-0000-1111-0000-000000000000"]
    relation   = "editor"
  }
  bindings {
    principals = ["org:10000000-0000-1111-0000-000000000000"]
    relation   = "viewer"
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `resource_id` (String) Identifier for the resource, formatted as resource_type:resource_id.

Note: Dashboards support is in private beta. Reach out to your Datadog contact or support to enable this.

### Optional

- `bindings` (Block Set) (see [below for nested schema](#nestedblock--bindings))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--bindings"></a>
### Nested Schema for `bindings`

Required:

- `principals` (Set of String) An array of principals. A principal is a subject or group of subjects. Each principal is formatted as `type:id`. Supported types: `role` and `org`. The org ID can be obtained through the api/v2/users API.
- `relation` (String) The role/level of access. See this page for more details https://docs.datadoghq.com/api/latest/restriction-policies/#supported-relations-for-resources

## Import

Import is supported using the following syntax:

```shell
terraform import datadog_restriction_policy.new_list ""
```