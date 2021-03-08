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

package stage

import (
	"context"
	"strings"

	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
	ackerr "github.com/aws-controllers-k8s/runtime/pkg/errors"
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
	_ = &svcapitypes.Stage{}
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

	resp, respErr := rm.sdkapi.GetStageWithContext(ctx, input)
	rm.metrics.RecordAPICall("READ_ONE", "GetStage", respErr)
	if respErr != nil {
		if awsErr, ok := ackerr.AWSError(respErr); ok && awsErr.Code() == "NotFoundException" {
			return nil, ackerr.NotFound
		}
		return nil, respErr
	}

	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := r.ko.DeepCopy()

	if resp.AccessLogSettings != nil {
		f0 := &svcapitypes.AccessLogSettings{}
		if resp.AccessLogSettings.DestinationArn != nil {
			f0.DestinationARN = resp.AccessLogSettings.DestinationArn
		}
		if resp.AccessLogSettings.Format != nil {
			f0.Format = resp.AccessLogSettings.Format
		}
		ko.Spec.AccessLogSettings = f0
	}
	if resp.ApiGatewayManaged != nil {
		ko.Status.APIGatewayManaged = resp.ApiGatewayManaged
	}
	if resp.AutoDeploy != nil {
		ko.Spec.AutoDeploy = resp.AutoDeploy
	}
	if resp.ClientCertificateId != nil {
		ko.Spec.ClientCertificateID = resp.ClientCertificateId
	}
	if resp.CreatedDate != nil {
		ko.Status.CreatedDate = &metav1.Time{*resp.CreatedDate}
	}
	if resp.DefaultRouteSettings != nil {
		f5 := &svcapitypes.RouteSettings{}
		if resp.DefaultRouteSettings.DataTraceEnabled != nil {
			f5.DataTraceEnabled = resp.DefaultRouteSettings.DataTraceEnabled
		}
		if resp.DefaultRouteSettings.DetailedMetricsEnabled != nil {
			f5.DetailedMetricsEnabled = resp.DefaultRouteSettings.DetailedMetricsEnabled
		}
		if resp.DefaultRouteSettings.LoggingLevel != nil {
			f5.LoggingLevel = resp.DefaultRouteSettings.LoggingLevel
		}
		if resp.DefaultRouteSettings.ThrottlingBurstLimit != nil {
			f5.ThrottlingBurstLimit = resp.DefaultRouteSettings.ThrottlingBurstLimit
		}
		if resp.DefaultRouteSettings.ThrottlingRateLimit != nil {
			f5.ThrottlingRateLimit = resp.DefaultRouteSettings.ThrottlingRateLimit
		}
		ko.Spec.DefaultRouteSettings = f5
	}
	if resp.DeploymentId != nil {
		ko.Spec.DeploymentID = resp.DeploymentId
	}
	if resp.Description != nil {
		ko.Spec.Description = resp.Description
	}
	if resp.LastDeploymentStatusMessage != nil {
		ko.Status.LastDeploymentStatusMessage = resp.LastDeploymentStatusMessage
	}
	if resp.LastUpdatedDate != nil {
		ko.Status.LastUpdatedDate = &metav1.Time{*resp.LastUpdatedDate}
	}
	if resp.RouteSettings != nil {
		f10 := map[string]*svcapitypes.RouteSettings{}
		for f10key, f10valiter := range resp.RouteSettings {
			f10val := &svcapitypes.RouteSettings{}
			if f10valiter.DataTraceEnabled != nil {
				f10val.DataTraceEnabled = f10valiter.DataTraceEnabled
			}
			if f10valiter.DetailedMetricsEnabled != nil {
				f10val.DetailedMetricsEnabled = f10valiter.DetailedMetricsEnabled
			}
			if f10valiter.LoggingLevel != nil {
				f10val.LoggingLevel = f10valiter.LoggingLevel
			}
			if f10valiter.ThrottlingBurstLimit != nil {
				f10val.ThrottlingBurstLimit = f10valiter.ThrottlingBurstLimit
			}
			if f10valiter.ThrottlingRateLimit != nil {
				f10val.ThrottlingRateLimit = f10valiter.ThrottlingRateLimit
			}
			f10[f10key] = f10val
		}
		ko.Spec.RouteSettings = f10
	}
	if resp.StageName != nil {
		ko.Spec.StageName = resp.StageName
	}
	if resp.StageVariables != nil {
		f12 := map[string]*string{}
		for f12key, f12valiter := range resp.StageVariables {
			var f12val string
			f12val = *f12valiter
			f12[f12key] = &f12val
		}
		ko.Spec.StageVariables = f12
	}
	if resp.Tags != nil {
		f13 := map[string]*string{}
		for f13key, f13valiter := range resp.Tags {
			var f13val string
			f13val = *f13valiter
			f13[f13key] = &f13val
		}
		ko.Spec.Tags = f13
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
	return r.ko.Spec.StageName == nil || r.ko.Spec.APIID == nil

}

// newDescribeRequestPayload returns SDK-specific struct for the HTTP request
// payload of the Describe API call for the resource
func (rm *resourceManager) newDescribeRequestPayload(
	r *resource,
) (*svcsdk.GetStageInput, error) {
	res := &svcsdk.GetStageInput{}

	if r.ko.Spec.APIID != nil {
		res.SetApiId(*r.ko.Spec.APIID)
	}
	if r.ko.Spec.StageName != nil {
		res.SetStageName(*r.ko.Spec.StageName)
	}

	return res, nil
}

// sdkCreate creates the supplied resource in the backend AWS service API and
// returns a new resource with any fields in the Status field filled in
func (rm *resourceManager) sdkCreate(
	ctx context.Context,
	r *resource,
) (*resource, error) {
	input, err := rm.newCreateRequestPayload(ctx, r)
	if err != nil {
		return nil, err
	}

	resp, respErr := rm.sdkapi.CreateStageWithContext(ctx, input)
	rm.metrics.RecordAPICall("CREATE", "CreateStage", respErr)
	if respErr != nil {
		return nil, respErr
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := r.ko.DeepCopy()

	if resp.ApiGatewayManaged != nil {
		ko.Status.APIGatewayManaged = resp.ApiGatewayManaged
	}
	if resp.CreatedDate != nil {
		ko.Status.CreatedDate = &metav1.Time{*resp.CreatedDate}
	}
	if resp.LastDeploymentStatusMessage != nil {
		ko.Status.LastDeploymentStatusMessage = resp.LastDeploymentStatusMessage
	}
	if resp.LastUpdatedDate != nil {
		ko.Status.LastUpdatedDate = &metav1.Time{*resp.LastUpdatedDate}
	}

	rm.setStatusDefaults(ko)

	return &resource{ko}, nil
}

// newCreateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Create API call for the resource
func (rm *resourceManager) newCreateRequestPayload(
	ctx context.Context,
	r *resource,
) (*svcsdk.CreateStageInput, error) {
	res := &svcsdk.CreateStageInput{}

	if r.ko.Spec.AccessLogSettings != nil {
		f0 := &svcsdk.AccessLogSettings{}
		if r.ko.Spec.AccessLogSettings.DestinationARN != nil {
			f0.SetDestinationArn(*r.ko.Spec.AccessLogSettings.DestinationARN)
		}
		if r.ko.Spec.AccessLogSettings.Format != nil {
			f0.SetFormat(*r.ko.Spec.AccessLogSettings.Format)
		}
		res.SetAccessLogSettings(f0)
	}
	if r.ko.Spec.APIID != nil {
		res.SetApiId(*r.ko.Spec.APIID)
	}
	if r.ko.Spec.AutoDeploy != nil {
		res.SetAutoDeploy(*r.ko.Spec.AutoDeploy)
	}
	if r.ko.Spec.ClientCertificateID != nil {
		res.SetClientCertificateId(*r.ko.Spec.ClientCertificateID)
	}
	if r.ko.Spec.DefaultRouteSettings != nil {
		f4 := &svcsdk.RouteSettings{}
		if r.ko.Spec.DefaultRouteSettings.DataTraceEnabled != nil {
			f4.SetDataTraceEnabled(*r.ko.Spec.DefaultRouteSettings.DataTraceEnabled)
		}
		if r.ko.Spec.DefaultRouteSettings.DetailedMetricsEnabled != nil {
			f4.SetDetailedMetricsEnabled(*r.ko.Spec.DefaultRouteSettings.DetailedMetricsEnabled)
		}
		if r.ko.Spec.DefaultRouteSettings.LoggingLevel != nil {
			f4.SetLoggingLevel(*r.ko.Spec.DefaultRouteSettings.LoggingLevel)
		}
		if r.ko.Spec.DefaultRouteSettings.ThrottlingBurstLimit != nil {
			f4.SetThrottlingBurstLimit(*r.ko.Spec.DefaultRouteSettings.ThrottlingBurstLimit)
		}
		if r.ko.Spec.DefaultRouteSettings.ThrottlingRateLimit != nil {
			f4.SetThrottlingRateLimit(*r.ko.Spec.DefaultRouteSettings.ThrottlingRateLimit)
		}
		res.SetDefaultRouteSettings(f4)
	}
	if r.ko.Spec.DeploymentID != nil {
		res.SetDeploymentId(*r.ko.Spec.DeploymentID)
	}
	if r.ko.Spec.Description != nil {
		res.SetDescription(*r.ko.Spec.Description)
	}
	if r.ko.Spec.RouteSettings != nil {
		f7 := map[string]*svcsdk.RouteSettings{}
		for f7key, f7valiter := range r.ko.Spec.RouteSettings {
			f7val := &svcsdk.RouteSettings{}
			if f7valiter.DataTraceEnabled != nil {
				f7val.SetDataTraceEnabled(*f7valiter.DataTraceEnabled)
			}
			if f7valiter.DetailedMetricsEnabled != nil {
				f7val.SetDetailedMetricsEnabled(*f7valiter.DetailedMetricsEnabled)
			}
			if f7valiter.LoggingLevel != nil {
				f7val.SetLoggingLevel(*f7valiter.LoggingLevel)
			}
			if f7valiter.ThrottlingBurstLimit != nil {
				f7val.SetThrottlingBurstLimit(*f7valiter.ThrottlingBurstLimit)
			}
			if f7valiter.ThrottlingRateLimit != nil {
				f7val.SetThrottlingRateLimit(*f7valiter.ThrottlingRateLimit)
			}
			f7[f7key] = f7val
		}
		res.SetRouteSettings(f7)
	}
	if r.ko.Spec.StageName != nil {
		res.SetStageName(*r.ko.Spec.StageName)
	}
	if r.ko.Spec.StageVariables != nil {
		f9 := map[string]*string{}
		for f9key, f9valiter := range r.ko.Spec.StageVariables {
			var f9val string
			f9val = *f9valiter
			f9[f9key] = &f9val
		}
		res.SetStageVariables(f9)
	}
	if r.ko.Spec.Tags != nil {
		f10 := map[string]*string{}
		for f10key, f10valiter := range r.ko.Spec.Tags {
			var f10val string
			f10val = *f10valiter
			f10[f10key] = &f10val
		}
		res.SetTags(f10)
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

	input, err := rm.newUpdateRequestPayload(ctx, desired)
	if err != nil {
		return nil, err
	}

	resp, respErr := rm.sdkapi.UpdateStageWithContext(ctx, input)
	rm.metrics.RecordAPICall("UPDATE", "UpdateStage", respErr)
	if respErr != nil {
		return nil, respErr
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := desired.ko.DeepCopy()

	if resp.ApiGatewayManaged != nil {
		ko.Status.APIGatewayManaged = resp.ApiGatewayManaged
	}
	if resp.CreatedDate != nil {
		ko.Status.CreatedDate = &metav1.Time{*resp.CreatedDate}
	}
	if resp.LastDeploymentStatusMessage != nil {
		ko.Status.LastDeploymentStatusMessage = resp.LastDeploymentStatusMessage
	}
	if resp.LastUpdatedDate != nil {
		ko.Status.LastUpdatedDate = &metav1.Time{*resp.LastUpdatedDate}
	}

	rm.setStatusDefaults(ko)

	return &resource{ko}, nil
}

// newUpdateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Update API call for the resource
func (rm *resourceManager) newUpdateRequestPayload(
	ctx context.Context,
	r *resource,
) (*svcsdk.UpdateStageInput, error) {
	res := &svcsdk.UpdateStageInput{}

	if r.ko.Spec.AccessLogSettings != nil {
		f0 := &svcsdk.AccessLogSettings{}
		if r.ko.Spec.AccessLogSettings.DestinationARN != nil {
			f0.SetDestinationArn(*r.ko.Spec.AccessLogSettings.DestinationARN)
		}
		if r.ko.Spec.AccessLogSettings.Format != nil {
			f0.SetFormat(*r.ko.Spec.AccessLogSettings.Format)
		}
		res.SetAccessLogSettings(f0)
	}
	if r.ko.Spec.APIID != nil {
		res.SetApiId(*r.ko.Spec.APIID)
	}
	if r.ko.Spec.AutoDeploy != nil {
		res.SetAutoDeploy(*r.ko.Spec.AutoDeploy)
	}
	if r.ko.Spec.ClientCertificateID != nil {
		res.SetClientCertificateId(*r.ko.Spec.ClientCertificateID)
	}
	if r.ko.Spec.DefaultRouteSettings != nil {
		f4 := &svcsdk.RouteSettings{}
		if r.ko.Spec.DefaultRouteSettings.DataTraceEnabled != nil {
			f4.SetDataTraceEnabled(*r.ko.Spec.DefaultRouteSettings.DataTraceEnabled)
		}
		if r.ko.Spec.DefaultRouteSettings.DetailedMetricsEnabled != nil {
			f4.SetDetailedMetricsEnabled(*r.ko.Spec.DefaultRouteSettings.DetailedMetricsEnabled)
		}
		if r.ko.Spec.DefaultRouteSettings.LoggingLevel != nil {
			f4.SetLoggingLevel(*r.ko.Spec.DefaultRouteSettings.LoggingLevel)
		}
		if r.ko.Spec.DefaultRouteSettings.ThrottlingBurstLimit != nil {
			f4.SetThrottlingBurstLimit(*r.ko.Spec.DefaultRouteSettings.ThrottlingBurstLimit)
		}
		if r.ko.Spec.DefaultRouteSettings.ThrottlingRateLimit != nil {
			f4.SetThrottlingRateLimit(*r.ko.Spec.DefaultRouteSettings.ThrottlingRateLimit)
		}
		res.SetDefaultRouteSettings(f4)
	}
	if r.ko.Spec.DeploymentID != nil {
		res.SetDeploymentId(*r.ko.Spec.DeploymentID)
	}
	if r.ko.Spec.Description != nil {
		res.SetDescription(*r.ko.Spec.Description)
	}
	if r.ko.Spec.RouteSettings != nil {
		f7 := map[string]*svcsdk.RouteSettings{}
		for f7key, f7valiter := range r.ko.Spec.RouteSettings {
			f7val := &svcsdk.RouteSettings{}
			if f7valiter.DataTraceEnabled != nil {
				f7val.SetDataTraceEnabled(*f7valiter.DataTraceEnabled)
			}
			if f7valiter.DetailedMetricsEnabled != nil {
				f7val.SetDetailedMetricsEnabled(*f7valiter.DetailedMetricsEnabled)
			}
			if f7valiter.LoggingLevel != nil {
				f7val.SetLoggingLevel(*f7valiter.LoggingLevel)
			}
			if f7valiter.ThrottlingBurstLimit != nil {
				f7val.SetThrottlingBurstLimit(*f7valiter.ThrottlingBurstLimit)
			}
			if f7valiter.ThrottlingRateLimit != nil {
				f7val.SetThrottlingRateLimit(*f7valiter.ThrottlingRateLimit)
			}
			f7[f7key] = f7val
		}
		res.SetRouteSettings(f7)
	}
	if r.ko.Spec.StageName != nil {
		res.SetStageName(*r.ko.Spec.StageName)
	}
	if r.ko.Spec.StageVariables != nil {
		f9 := map[string]*string{}
		for f9key, f9valiter := range r.ko.Spec.StageVariables {
			var f9val string
			f9val = *f9valiter
			f9[f9key] = &f9val
		}
		res.SetStageVariables(f9)
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
	_, respErr := rm.sdkapi.DeleteStageWithContext(ctx, input)
	rm.metrics.RecordAPICall("DELETE", "DeleteStage", respErr)
	return respErr
}

// newDeleteRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Delete API call for the resource
func (rm *resourceManager) newDeleteRequestPayload(
	r *resource,
) (*svcsdk.DeleteStageInput, error) {
	res := &svcsdk.DeleteStageInput{}

	if r.ko.Spec.APIID != nil {
		res.SetApiId(*r.ko.Spec.APIID)
	}
	if r.ko.Spec.StageName != nil {
		res.SetStageName(*r.ko.Spec.StageName)
	}

	return res, nil
}

// setStatusDefaults sets default properties into supplied custom resource
func (rm *resourceManager) setStatusDefaults(
	ko *svcapitypes.Stage,
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
