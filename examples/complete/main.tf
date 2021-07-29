resource "aws_kms_key" "this" {
  deletion_window_in_days = 30
}

module "bucket" {
  source = "../.."

  bucket                  = "my-bucket-123"
  block_public_acls       = true
  block_public_policy     = true
  ignore_public_acls      = true
  restrict_public_buckets = true

  tags = {
    "managed-by": "terraform",
    "env": "staging"
  }

  sse_configuration = {
    rule = {
      apply_server_side_encryption_by_default = {
        sse_algorithm     = "aws:kms"
        kms_master_key_id = aws_kms_key.this.arn
      }
    }
  }
}
