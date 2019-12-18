# s3pre

A tool to generate a pre-signed URL for an Amazon S3 object

With this tool, you can:
* generate pre-signed URL for S3 object easily
* set expiration date easily and readable
 
Its benefits over 'aws s3 presign' are:
* more readable format for expiration date
* you will type less than 'aws cli'
* (COMING) it will detect the region, so you don't need to specify the region for the object
 
### USAGE

* **AWS CLI:** aws s3 presign s3://test_bucket/test_object.obj --region ap-northeast-2 --expires-in 36000 (in seconds)
* **S3PRE:** s3pre s3://test_bucket/test_object.obj -r=ap-northeast-2 -e=1h (in readable format, 10m, 1h, 2d, etc.)
 
> **Note:** Don't forget to add quotations around the key when it contains any spaces (e.g. s3://test_bucket/aws korea.pdf).