<h1 align="center">terraform-local-naming</h1> 
<div align="center">

[![Build Status](https://dev.azure.com/terraformation/Terraformation/_apis/build/status/terraformation.terraform-local-naming?branchName=master)](https://dev.azure.com/terraformation/Terraformation/_build/latest?definitionId=2&branchName=master)
[![Licence](https://img.shields.io/github/license/terraformation/terraform-local-naming)](https://github.com/terraformation/terraform-local-naming/blob/master/LICENSE)
![Contributors](https://img.shields.io/github/contributors/terraformation/terraform-local-naming)
![Last commit](https://img.shields.io/github/last-commit/terraformation/terraform-local-naming)
![Last commit](https://img.shields.io/github/commit-activity/m/terraformation/terraform-local-naming)
[![Maintained by Terraformation](https://img.shields.io/badge/maintained%20by-Terraformation-%235849a6.svg)](https://github.com/terraformation)
![Terraform Version](https://img.shields.io/badge/tf-%3E%3D0.12.0-blue.svg)
![Chat](https://img.shields.io/gitter/room/terraformation/community)

</div>


## What does this Module ?
This module provides a mechanism to create consistent naming across multiple resources. The origin of the creation of this module is linked to the fact that most of company around the world commonly use standardized naming to describe their it resources.

Behind this standardize naming, their is often a meaning to give some information about of the context of the resource. (Like the zone where is hosted the resource).

## How to use this Module ?
```hcl
module "naming_module" {
  source = "terraformation/naming/local"
  naming_options = {
    suffix = "test"
    prefix = "test2"
  }
}
```

### Inputs
| Name | Description |Type| Default | Required |
|------|-------------|:-----:|:-----:|:-----:|
|naming_options|The options to pass to the generator|object/any|{}|no
### Outputs
| Name | Description |
|------|-------------|
|rendered|The value of the generated name.|

## What's a Module?

A Module is a canonical, reusable, best-practices definition for how to run a single piece of infrastructure, such 
as a database or server cluster. Each Module is created using [Terraform](https://www.terraform.io/), and
includes automated tests, examples, and documentation. It is maintained both by the open source community and 
companies that provide commercial support. 

Instead of figuring out the details of how to run a piece of infrastructure from scratch, you can reuse 
existing code that has been proven in production. And instead of maintaining all that infrastructure code yourself, 
you can leverage the work of the Module community to pick up infrastructure improvements through
a version number bump.

## Who maintains this Module?

This Module is maintained by Terraformation. If you're looking for help please contact us on the following channels.

Also, you can find the complete list of contributors of this project [here](https://github.com/terraformation/terraform-local-naming/tree/master/CONTRIBUTORS.md).

## How do I contribute to this Module?

Contributions are very welcome! Check out the [Contribution Guidelines](https://github.com/terraformation/terraform-local-naming/tree/master/CONTRIBUTING.md) for instructions.

## How is this Module versioned?

This Module follows the principles of [Semantic Versioning](http://semver.org/). You can find each new release, 
along with the changelog, in the [Releases Page](../../releases). 

During initial development, the major version will be 0 (e.g., `0.x.y`), which indicates the code does not yet have a 
stable API. Once we hit `1.0.0`, we will make every effort to maintain a backwards compatible API and use the MAJOR, 
MINOR, and PATCH versions on each release to indicate any incompatibilities. 

You will find a synthesis of each published version in the [changelog](https://github.com/terraformation/terraform-local-naming/tree/master/CHANGELOG.md) file.

## License

This code is released under the MIT License. Please see [LICENSE](https://github.com/terraformation/terraform-local-naming/tree/master/LICENSE) and [NOTICE](https://github.com/terraformation/terraform-local-naming/tree/master/NOTICE) for more 
details.

Copyright (c) 2019 Terraformation
