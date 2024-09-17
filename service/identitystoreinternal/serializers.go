// Code generated by smithy-go-codegen DO NOT EDIT.

package identitystoreinternal

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/identitystoreinternal/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/encoding/httpbinding"
	smithyjson "github.com/aws/smithy-go/encoding/json"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"path"
)

type awsAwsjson11_serializeOpListBearerTokens struct {
}

func (*awsAwsjson11_serializeOpListBearerTokens) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpListBearerTokens) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*ListBearerTokensInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	operationPath := "/"
	if len(request.Request.URL.Path) == 0 {
		request.Request.URL.Path = operationPath
	} else {
		request.Request.URL.Path = path.Join(request.Request.URL.Path, operationPath)
		if request.Request.URL.Path != "/" && operationPath[len(operationPath)-1] == '/' {
			request.Request.URL.Path += "/"
		}
	}
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.1")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("AWSIdentityStoreService.ListBearerTokens")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeOpDocumentListBearerTokensInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson11_serializeOpListProvisioningTenants struct {
}

func (*awsAwsjson11_serializeOpListProvisioningTenants) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpListProvisioningTenants) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*ListProvisioningTenantsInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	operationPath := "/"
	if len(request.Request.URL.Path) == 0 {
		request.Request.URL.Path = operationPath
	} else {
		request.Request.URL.Path = path.Join(request.Request.URL.Path, operationPath)
		if request.Request.URL.Path != "/" && operationPath[len(operationPath)-1] == '/' {
			request.Request.URL.Path += "/"
		}
	}
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.1")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("AWSIdentityStoreService.ListProvisioningTenants")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeOpDocumentListProvisioningTenantsInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson11_serializeOpSearchUsers struct {
}

func (*awsAwsjson11_serializeOpSearchUsers) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpSearchUsers) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*SearchUsersInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	operationPath := "/"
	if len(request.Request.URL.Path) == 0 {
		request.Request.URL.Path = operationPath
	} else {
		request.Request.URL.Path = path.Join(request.Request.URL.Path, operationPath)
		if request.Request.URL.Path != "/" && operationPath[len(operationPath)-1] == '/' {
			request.Request.URL.Path += "/"
		}
	}
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.1")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("AWSIdentityStoreService.SearchUsers")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeOpDocumentSearchUsersInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsAwsjson11_serializeDocumentFilter(v *types.Filter, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.AttributePath != nil {
		ok := object.Key("AttributePath")
		ok.String(*v.AttributePath)
	}

	if v.AttributeValue != nil {
		ok := object.Key("AttributeValue")
		ok.String(*v.AttributeValue)
	}

	return nil
}

func awsAwsjson11_serializeDocumentFilters(v []types.Filter, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		if err := awsAwsjson11_serializeDocumentFilter(&v[i], av); err != nil {
			return err
		}
	}
	return nil
}

func awsAwsjson11_serializeOpDocumentListBearerTokensInput(v *ListBearerTokensInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.IdentityStoreId != nil {
		ok := object.Key("IdentityStoreId")
		ok.String(*v.IdentityStoreId)
	}

	if v.TenantId != nil {
		ok := object.Key("TenantId")
		ok.String(*v.TenantId)
	}

	return nil
}

func awsAwsjson11_serializeOpDocumentListProvisioningTenantsInput(v *ListProvisioningTenantsInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.IdentityStoreId != nil {
		ok := object.Key("IdentityStoreId")
		ok.String(*v.IdentityStoreId)
	}

	return nil
}

func awsAwsjson11_serializeOpDocumentSearchUsersInput(v *SearchUsersInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Filters != nil {
		ok := object.Key("Filters")
		if err := awsAwsjson11_serializeDocumentFilters(v.Filters, ok); err != nil {
			return err
		}
	}

	if v.IdentityStoreId != nil {
		ok := object.Key("IdentityStoreId")
		ok.String(*v.IdentityStoreId)
	}

	if v.MaxResults != nil {
		ok := object.Key("MaxResults")
		ok.Integer(*v.MaxResults)
	}

	if v.NextToken != nil {
		ok := object.Key("NextToken")
		ok.String(*v.NextToken)
	}

	return nil
}
