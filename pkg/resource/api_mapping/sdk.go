// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package api_mapping

import (
	"context"
	"strings"

	ackv1alpha1 "github.com/aws/aws-controllers-k8s/apis/core/v1alpha1"
	ackcompare "github.com/aws/aws-controllers-k8s/pkg/compare"
	ackerr "github.com/aws/aws-controllers-k8s/pkg/errors"
	"github.com/aws/aws-sdk-go/aws"
	svcsdk "github.com/aws/aws-sdk-go/service/apigatewayv2"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	svcapitypes "github.com/aws-controllers-k8s/apigatewayv2-controller/apis/v1alpha1"
)

// Hack to avoid import errors during build...
var (
	_ = &metav1.Time{}
	_ = strings.ToLower("")
	_ = &aws.JSONValue{}
	_ = &svcsdk.ApiGatewayV2{}
	_ = &svcapitypes.APIMapping{}
	_ = ackv1alpha1.AWSAccountID("")
	_ = &ackerr.NotFound
)

// sdkFind returns SDK-specific information about a supplied resource
func (rm *resourceManager) sdkFind(
	ctx context.Context,
	r *resource,
) (*resource, error) {
	// If any required fields in the input shape are missing, AWS resource is
	// not created yet. Return NotFound here to indicate to callers that the
	// resource isn't yet created.
	if rm.requiredFieldsMissingFromReadOneInput(r) {
		return nil, ackerr.NotFound
	}

	input, err := rm.newDescribeRequestPayload(r)
	if err != nil {
		return nil, err
	}

	resp, respErr := rm.sdkapi.GetApiMappingWithContext(ctx, input)
	rm.metrics.RecordAPICall("READ_ONE", "GetApiMapping", respErr)
	if respErr != nil {
		if awsErr, ok := ackerr.AWSError(respErr); ok && awsErr.Code() == "NotFoundException" {
			return nil, ackerr.NotFound
		}
		return nil, respErr
	}

	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := r.ko.DeepCopy()

	if resp.ApiId != nil {
		ko.Spec.APIID = resp.ApiId
	}
	if resp.ApiMappingId != nil {
		ko.Status.APIMappingID = resp.ApiMappingId
	}
	if resp.ApiMappingKey != nil {
		ko.Spec.APIMappingKey = resp.ApiMappingKey
	}
	if resp.Stage != nil {
		ko.Spec.Stage = resp.Stage
	}

	rm.setStatusDefaults(ko)
	return &resource{ko}, nil
}

// requiredFieldsMissingFromReadOneInput returns true if there are any fields
// for the ReadOne Input shape that are required by not present in the
// resource's Spec or Status
func (rm *resourceManager) requiredFieldsMissingFromReadOneInput(
	r *resource,
) bool {
	return r.ko.Status.APIMappingID == nil || r.ko.Spec.DomainName == nil

}

// newDescribeRequestPayload returns SDK-specific struct for the HTTP request
// payload of the Describe API call for the resource
func (rm *resourceManager) newDescribeRequestPayload(
	r *resource,
) (*svcsdk.GetApiMappingInput, error) {
	res := &svcsdk.GetApiMappingInput{}

	if r.ko.Status.APIMappingID != nil {
		res.SetApiMappingId(*r.ko.Status.APIMappingID)
	}
	if r.ko.Spec.DomainName != nil {
		res.SetDomainName(*r.ko.Spec.DomainName)
	}

	return res, nil
}

// sdkCreate creates the supplied resource in the backend AWS service API and
// returns a new resource with any fields in the Status field filled in
func (rm *resourceManager) sdkCreate(
	ctx context.Context,
	r *resource,
) (*resource, error) {
	input, err := rm.newCreateRequestPayload(r)
	if err != nil {
		return nil, err
	}

	resp, respErr := rm.sdkapi.CreateApiMappingWithContext(ctx, input)
	rm.metrics.RecordAPICall("CREATE", "CreateApiMapping", respErr)
	if respErr != nil {
		return nil, respErr
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := r.ko.DeepCopy()

	if resp.ApiMappingId != nil {
		ko.Status.APIMappingID = resp.ApiMappingId
	}

	rm.setStatusDefaults(ko)

	return &resource{ko}, nil
}

// newCreateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Create API call for the resource
func (rm *resourceManager) newCreateRequestPayload(
	r *resource,
) (*svcsdk.CreateApiMappingInput, error) {
	res := &svcsdk.CreateApiMappingInput{}

	if r.ko.Spec.APIID != nil {
		res.SetApiId(*r.ko.Spec.APIID)
	}
	if r.ko.Spec.APIMappingKey != nil {
		res.SetApiMappingKey(*r.ko.Spec.APIMappingKey)
	}
	if r.ko.Spec.DomainName != nil {
		res.SetDomainName(*r.ko.Spec.DomainName)
	}
	if r.ko.Spec.Stage != nil {
		res.SetStage(*r.ko.Spec.Stage)
	}

	return res, nil
}

// sdkUpdate patches the supplied resource in the backend AWS service API and
// returns a new resource with updated fields.
func (rm *resourceManager) sdkUpdate(
	ctx context.Context,
	desired *resource,
	latest *resource,
	diffReporter *ackcompare.Reporter,
) (*resource, error) {

	input, err := rm.newUpdateRequestPayload(desired)
	if err != nil {
		return nil, err
	}

	resp, respErr := rm.sdkapi.UpdateApiMappingWithContext(ctx, input)
	rm.metrics.RecordAPICall("UPDATE", "UpdateApiMapping", respErr)
	if respErr != nil {
		return nil, respErr
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := desired.ko.DeepCopy()

	if resp.ApiMappingId != nil {
		ko.Status.APIMappingID = resp.ApiMappingId
	}

	rm.setStatusDefaults(ko)

	return &resource{ko}, nil
}

// newUpdateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Update API call for the resource
func (rm *resourceManager) newUpdateRequestPayload(
	r *resource,
) (*svcsdk.UpdateApiMappingInput, error) {
	res := &svcsdk.UpdateApiMappingInput{}

	if r.ko.Spec.APIID != nil {
		res.SetApiId(*r.ko.Spec.APIID)
	}
	if r.ko.Status.APIMappingID != nil {
		res.SetApiMappingId(*r.ko.Status.APIMappingID)
	}
	if r.ko.Spec.APIMappingKey != nil {
		res.SetApiMappingKey(*r.ko.Spec.APIMappingKey)
	}
	if r.ko.Spec.DomainName != nil {
		res.SetDomainName(*r.ko.Spec.DomainName)
	}
	if r.ko.Spec.Stage != nil {
		res.SetStage(*r.ko.Spec.Stage)
	}

	return res, nil
}

// sdkDelete deletes the supplied resource in the backend AWS service API
func (rm *resourceManager) sdkDelete(
	ctx context.Context,
	r *resource,
) error {
	input, err := rm.newDeleteRequestPayload(r)
	if err != nil {
		return err
	}
	_, respErr := rm.sdkapi.DeleteApiMappingWithContext(ctx, input)
	rm.metrics.RecordAPICall("DELETE", "DeleteApiMapping", respErr)
	return respErr
}

// newDeleteRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Delete API call for the resource
func (rm *resourceManager) newDeleteRequestPayload(
	r *resource,
) (*svcsdk.DeleteApiMappingInput, error) {
	res := &svcsdk.DeleteApiMappingInput{}

	if r.ko.Status.APIMappingID != nil {
		res.SetApiMappingId(*r.ko.Status.APIMappingID)
	}
	if r.ko.Spec.DomainName != nil {
		res.SetDomainName(*r.ko.Spec.DomainName)
	}

	return res, nil
}

// setStatusDefaults sets default properties into supplied custom resource
func (rm *resourceManager) setStatusDefaults(
	ko *svcapitypes.APIMapping,
) {
	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if ko.Status.ACKResourceMetadata.OwnerAccountID == nil {
		ko.Status.ACKResourceMetadata.OwnerAccountID = &rm.awsAccountID
	}
	if ko.Status.Conditions == nil {
		ko.Status.Conditions = []*ackv1alpha1.Condition{}
	}
}

// updateConditions returns updated resource, true; if conditions were updated
// else it returns nil, false
func (rm *resourceManager) updateConditions(
	r *resource,
	err error,
) (*resource, bool) {
	ko := r.ko.DeepCopy()
	rm.setStatusDefaults(ko)

	// Terminal condition
	var terminalCondition *ackv1alpha1.Condition = nil
	for _, condition := range ko.Status.Conditions {
		if condition.Type == ackv1alpha1.ConditionTypeTerminal {
			terminalCondition = condition
			break
		}
	}

	if rm.terminalAWSError(err) {
		if terminalCondition == nil {
			terminalCondition = &ackv1alpha1.Condition{
				Type: ackv1alpha1.ConditionTypeTerminal,
			}
			ko.Status.Conditions = append(ko.Status.Conditions, terminalCondition)
		}
		terminalCondition.Status = corev1.ConditionTrue
		awsErr, _ := ackerr.AWSError(err)
		errorMessage := awsErr.Message()
		terminalCondition.Message = &errorMessage
	} else if terminalCondition != nil {
		terminalCondition.Status = corev1.ConditionFalse
		terminalCondition.Message = nil
	}
	if terminalCondition != nil {
		return &resource{ko}, true // updated
	}
	return nil, false // not updated
}

// terminalAWSError returns awserr, true; if the supplied error is an aws Error type
// and if the exception indicates that it is a Terminal exception
// 'Terminal' exception are specified in generator configuration
func (rm *resourceManager) terminalAWSError(err error) bool {
	// No terminal_errors specified for this resource in generator config
	return false
}
