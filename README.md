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

## AWS session

## S3 Implementation

The implementation about S3 bucket is in the `bucket.go`, and about S3 object is in the `object.go`.

For now there are `CreateBucket`, `ListBuckets`, `DeleteBucket`, `ListObjects`, `UploadObject`, `DownloadObject`, 
`CopyObject`, `DeleteObject`.
