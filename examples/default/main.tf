module "default" {
  source         = "../../"
  naming_options = {
    prefix        = "prefix"
    suffix        = "suffix"
    resource_name = "resource_name"
  }
}
