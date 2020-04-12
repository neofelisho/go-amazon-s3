# Amazon S3 implementation with Go

This project implements the basic AWS S3 services for Go.

## Install the AWS SDK for GO

[Reference](https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/setting-up.html)

We don't need to install it again if we just want to run this project.

## Setup AWS configurations

```shell script
$ export AWS_REGION=YOUR_AWS_REGION
$ export AWS_ACCESS_KEY_ID=YOUR_AKID
$ export AWS_SECRET_ACCESS_KEY=YOUR_SECRET_KEY 
```

For GoLand user can set up in `File -> Settings -> Go/Go Modules`:

![GoLand env variables](
https://user-images.githubusercontent.com/13026209/79066546-f00c1c00-7ce2-11ea-81d0-4124a764e666.png)

## AWS Session

Before using S3 service, we need to create a session of AWS. We can use `CreateSession` method in the `session.go` to 
create a new session.

## S3 Implementation

The implementation about S3 bucket is in the `bucket.go`, and about S3 object is in the `object.go`.

For now there are `CreateBucket`, `ListBuckets`, `DeleteBucket`, `ListObjects`, `UploadObject`, `DownloadObject`, 
`CopyObject`, `DeleteObject`.

## S3 Integration Test

After the implementation of S3 services, we implement the integration tests in `bucket_test.go` and `object_test.go`.