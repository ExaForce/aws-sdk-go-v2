module github.com/aws/aws-sdk-go-v2/service/internal/checksum

go 1.20

require (
	github.com/aws/aws-sdk-go-v2 v1.27.2
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.11.11
	github.com/aws/smithy-go v1.20.2
)

replace github.com/aws/aws-sdk-go-v2 => ../../../

replace github.com/aws/aws-sdk-go-v2/service/internal/presigned-url => ../../../service/internal/presigned-url/
