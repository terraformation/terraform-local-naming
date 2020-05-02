locals {
  naming_options = merge({
    resource_name = ""
    suffix        = ""
    prefix        = ""
    separator     = "-"
    lower         = true
    length        = 0
  }, var.naming_options)
  suffix         = local.naming_options.suffix == "" ? random_id.default[0].hex : local.naming_options.suffix
  items          = [
    local.naming_options.prefix,
    local.naming_options.resource_name,
    local.suffix
  ]
  join           = join(local.naming_options.separator, compact(local.items))
  case           = local.naming_options.lower ? lower(local.join) : upper(local.join)
  rendered       = local.naming_options.length == 0 ? local.case: substr(local.case, 0, local.naming_options.length)
}
resource "random_id" "default" {
  depends_on  = [
    null_resource.module_depends_on
  ]
  count       = local.naming_options.suffix == "" ? 1 : 0
  byte_length = 4
}

resource "null_resource" "module_depends_on" {
  triggers = {
    value = length(var.module_depends_on)
  }
}
