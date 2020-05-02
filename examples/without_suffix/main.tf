module "default" {
  source         = "../../"
  naming_options = {
    prefix        = "prefix"
    resource_name = "resource_name"
  }
}
