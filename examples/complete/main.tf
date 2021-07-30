resource "aws_kms_key" "this" {
  deletion_window_in_days = 30
}

module "my-bucket" {
  source                  = "../.."
  create                  = var.create_bucket
  bucket                  = var.bucket
  block_public_acls       = var.block_public_acls
  block_public_policy     = var.block_public_policy
  ignore_public_acls      = var.ignore_public_acls
  restrict_public_buckets = var.restrict_public_buckets
  tags                    = var.tags

  sse_configuration = {
    rule = {
      apply_server_side_encryption_by_default = {
        sse_algorithm     = var.sse_algorithm
        kms_master_key_id = aws_kms_key.this.arn
      }
    }
  }
}
