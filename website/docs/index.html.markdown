---
layout: "pagerduty"
page_title: "Provider: PagerDuty"
sidebar_current: "docs-pagerduty-index"
description: |-
  PagerDuty is an alarm aggregation and dispatching service
---

# PagerDuty Provider

[PagerDuty](https://www.pagerduty.com/) is an alarm aggregation and dispatching service for system administrators and support teams. It collects alerts from your monitoring tools, gives you an overall view of all of your monitoring alarms, and alerts an on duty engineer if there’s a problem.

## Nordcloud's fork

This fork is intended to introduce many improvements over the official provider, including faster bug fixing time and multiple performance improvements. It's used internally by Nordcloud in a big-scale environment and produces drastically faster results compared to the upstream one.

Please keep in mind that some resources are not compatible with the implementation in the official provider and may require code or Terraform state changes. This documentation is always up to date with the current resource implementation. We don't, however, guarantee backwards compatibility even between minor releases. If you're using this provider in a production environment, make sure to define a specific version requirement in your provider definition so that our updates don't break your workflow.

## Example Usage

```hcl
# Configure the PagerDuty provider
provider "pagerduty" {
  token = var.pagerduty_token
}

# Create a PagerDuty team
resource "pagerduty_team" "engineering" {
  name        = "Engineering"
  description = "All engineering"
}

# Create a PagerDuty user
resource "pagerduty_user" "earline" {
  name  = "Earline Greenholt"
  email = "125.greenholt.earline@graham.name"
  teams = [pagerduty_team.engineering.id]
}
```

Use the navigation to the left to read about the available resources.

## Argument Reference

The following arguments are supported:

* `token` - (Required) The v2 authorization token. It can also be sourced from the PAGERDUTY_TOKEN environment variable. See [API Documentation](https://v2.developer.pagerduty.com/docs/authentication) for more information.
* `skip_credentials_validation` - (Optional) Skip validation of the token against the PagerDuty API.
