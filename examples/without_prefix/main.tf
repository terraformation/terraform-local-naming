module "default" {
  source         = "../../"
  naming_options = {
    suffix        = "suffix"
    resource_name = "resource_name"
  }
}
