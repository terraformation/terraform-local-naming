module "default" {
  source         = "../../"
  naming_options = {
    prefix        = "Prefix"
    suffix        = "Suffix"
    resource_name = "resource_name"
  }
}
