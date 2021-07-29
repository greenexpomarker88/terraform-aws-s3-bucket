variable "create" {
  description = "Determines if resources in this module should be created. Defaults to true."
  type        = bool
  default     = true
}

variable "bucket" {
  description = "The name of the bucket."
  type        = string
}

variable "acl" {
  description = "The canned access control list (ACL) to be applied to the bucket. Defaults to `private`."
  type        = string
  default     = "private"
}

variable "tags" {
  description = "Key value pairs to be attached to all resources in the module."
  type        = map(string)
  default     = {}
}

variable "block_public_acls" {
  description = "Determines if the bucket should block public ACLs."
  type        = bool
  default     = true
}

variable "block_public_policy" {
  description = "Determines if the bucket should block public bucket policies."
  type        = bool
  default     = true
}

variable "ignore_public_acls" {
  description = "Determines if the bucket should ignore public ACLs."
  type        = bool
  default     = true
}

variable "restrict_public_buckets" {
  description = "Determines if the bucket should restrict bucket policies."
  type        = bool
  default     = true
}

variable "versioning_enabled" {
  description = "Determines if versioning should be enabled on this bucket."
  type        = bool
  default     = true
}

variable "sse_enabled" {
  description = "Determines if server side encryption should be enabled for this bucket."
  type        = bool
  default     = true
}

variable "sse_configuration" {
  description = "Map of server side encryption configuration properties."
  type        = any
  default     = {}
}

variable "logging_enabled" {
  description = "Determines if logging should be enabled for this bucket."
  type        = bool
  default     = false
}

variable "logging_configuration" {
  description = "Map of logging configuration properties."
  type        = any
  default     = {}
}
