# AWS S3 Bucket Terraform Module
Terraform Module that can be used to configure an S3 bucket.

### Requirements
* aws >= 3.51
* terraform >= 1.0.0

### Inputs
| Input                   	| Description                                                                              	| Type   	| Default   	| Required 	|
|-------------------------	|------------------------------------------------------------------------------------------	|--------	|-----------	|----------	|
| create                  	| Determines if resources in this module should be created. Defaults to true.              	| bool   	| true      	| yes      	|
| bucket                  	| The name of the bucket.                                                                  	| string 	|           	| yes      	|
| acl                     	| The canned access control list (ACL) to be applied to the bucket. Defaults to `private`. 	| string 	| "private" 	| no       	|
| tags                    	| Key value pairs to be attached to all resources in the module.                           	| map    	| {}        	| no       	|
| block_public_acls       	| Determines if the bucket should block public ACLs.                                       	| bool   	| true      	| no       	|
| block_public_policy     	| Determines if the bucket should block public bucket policies.                            	| bool   	| true      	| no       	|
| ignore_public_acls      	| Determines if the bucket should block public bucket policies.                            	| bool   	| true      	| no       	|
| restrict_public_buckets 	| Determines if the bucket should restrict bucket policies.                                	| bool   	| true      	| no       	|
| versioning_enabled      	| Determines if versioning should be enabled on this bucket.                               	| bool   	| true      	| no       	|
| sse_enabled             	| Determines if server side encryption should be enabled for this bucket.                  	| bool   	| true      	| no       	|
| sse_configuration       	| Map of server side encryption configuration properties.                                  	| any    	| {}        	| no       	|
| logging_enabled         	| Determines if logging should be enabled for this bucket.                                 	| bool   	| false     	| no       	|
| logging_configuration   	| Map of logging configuration properties.                                                 	| any    	| {}        	| no       	|

### Outputs
| Output        	| Description                              	|
|---------------	|------------------------------------------	|
| bucket        	| The name of the bucket.                  	|
| bucket_arn    	| The ARN of the bucket.                   	|
| bucket_region 	| The AWS region where this bucket exists. 	|

### Local Development and Testing
Testing during local development can be achieved by following these steps:

You should have the following locally:
* go >= 1.16 installed
* aws cli configured with an aws account

Execute the following commands to run the test suite.
```shell
cd test
go mod download
go test
```
