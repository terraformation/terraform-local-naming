locals {
  naming_options          = merge({
    resource_name = ""
    suffix        = ""
    prefix        = ""
    hyphen        = true
    lower         = true
    length        = 0
  }, var.naming_options)
  base                    = local.naming_options.prefix
  base_with_resource_name = local.naming_options.resource_name =="" ? local.base : (local.base =="" ? local.naming_options.resource_name: format("%s-%s",  local.base, local.naming_options.resource_name))
  base_with_prefix        = local.naming_options.suffix =="" ? local.base_with_resource_name : (local.base_with_resource_name =="" ? local.naming_options.suffix: format("%s-%s", local.base_with_resource_name, local.naming_options.suffix))
  base_without_hyphen     = local.naming_options.hyphen ? local.base_with_prefix : replace(local.base_with_prefix, "-", "")
  base_without_case       = local.naming_options.lower ? lower(local.base_without_hyphen) : upper(local.base_without_hyphen)
  //  trimmed_result          = trimspace(local.name_without_case)
  //  result                  = local.naming_options.length != 0 ?substr(local.trimmed_result, 0, local.naming_options.length): local.trimmed_result
  result                  = local.base_without_case
}
