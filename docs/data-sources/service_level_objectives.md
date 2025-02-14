---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "datadog_service_level_objectives Data Source - terraform-provider-datadog"
subcategory: ""
description: |-
  Use this data source to retrieve information about multiple SLOs for use in other resources.
---

# datadog_service_level_objectives (Data Source)

Use this data source to retrieve information about multiple SLOs for use in other resources.

## Example Usage

```terraform
data "datadog_service_level_objectives" "ft_foo_slos" {
  tags_query = "owner:ft-foo"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `ids` (List of String) An array of SLO IDs to limit the search.
- `metrics_query` (String) Filter results based on SLO numerator and denominator.
- `name_query` (String) Filter results based on SLO names.
- `tags_query` (String) Filter results based on a single SLO tag.

### Read-Only

- `id` (String) The ID of this resource.
- `slos` (List of Object) List of SLOs (see [below for nested schema](#nestedatt--slos))

<a id="nestedatt--slos"></a>
### Nested Schema for `slos`

Read-Only:

- `id` (String)
- `name` (String)
- `type` (String)
