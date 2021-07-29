resource "aws_s3_bucket" "this" {
  count = var.create ? 1 : 0

  bucket = var.bucket
  acl    = var.acl
  tags   = var.tags

  versioning {
    enabled = var.versioning_enabled
  }

  dynamic "server_side_encryption_configuration" {
    for_each = var.sse_enabled ? [var.sse_configuration] : []
    content {

      dynamic "rule" {
        for_each = length(keys(lookup(server_side_encryption_configuration.value, "rule", {}))) > 0 ? [lookup(server_side_encryption_configuration.value, "rule", {})] : []
        content {

          dynamic "apply_server_side_encryption_by_default" {
            for_each = length(keys(lookup(rule.value, "apply_server_side_encryption_by_default", {}))) > 0 ? [lookup(rule.value, "apply_server_side_encryption_by_default", {})] : []
            content {
              sse_algorithm     = apply_server_side_encryption_by_default.value.sse_algorithm
              kms_master_key_id = lookup(apply_server_side_encryption_by_default.value, "kms_master_key_id", null)
            }
          }
        }
      }
    }
  }

  dynamic "logging" {
    for_each = var.logging_enabled ? [var.logging_configuration] : []
    content {
      target_bucket = logging.value.target_bucket
      target_prefix = lookup(looking.value, "target_prefix", null)
    }
  }
}

resource "aws_s3_bucket_public_access_block" "this" {
  count = var.create ? 1 : 0

  bucket                  = aws_s3_bucket.this[0].id
  block_public_acls       = var.block_public_acls
  block_public_policy     = var.block_public_policy
  ignore_public_acls      = var.ignore_public_acls
  restrict_public_buckets = var.restrict_public_buckets
}
