output "bucket" {
  description = "The name of the bucket"
  value       = concat(aws_s3_bucket.this.*.bucket, [""])[0]
}

output "bucket_arn" {
  description = "The ARN of the bucket."
  value = concat(aws_s3_bucket.this.*.arn, [""])[0]
}

output "bucket_region" {
  description = "The AWS region where this bucket exists."
  value = concat(aws_s3_bucket.this.*.region, [""])[0]
}
