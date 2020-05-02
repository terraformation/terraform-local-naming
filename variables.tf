variable "naming_options" {
  type    = map(any)
  default = {}
  description = "(Optional) Options to pass to th namee generator"
}
variable "module_depends_on" {
  default = [""]
  description = "(Optional) "
}
