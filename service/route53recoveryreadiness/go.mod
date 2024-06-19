module github.com/aws/aws-sdk-go-v2/service/route53recoveryreadiness

go 1.20

require (
	github.com/aws/aws-sdk-go-v2 v1.30.0
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.3.12
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.6.12
	github.com/aws/smithy-go v1.20.2
)

replace github.com/aws/aws-sdk-go-v2 => ../../

replace github.com/aws/aws-sdk-go-v2/internal/configsources => ../../internal/configsources/

replace github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 => ../../internal/endpoints/v2/
