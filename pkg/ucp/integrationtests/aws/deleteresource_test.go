// ------------------------------------------------------------
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.
// ------------------------------------------------------------

package aws

// Tests that test with Mock RP functionality and UCP Server

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/Azure/go-autorest/autorest/to"
	"github.com/aws/aws-sdk-go-v2/service/cloudcontrol"
	"github.com/aws/aws-sdk-go-v2/service/cloudcontrol/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_DeleteAWSResource(t *testing.T) {
	ucp, ucpClient, cloudcontrolClient, _ := initializeTest(t)

	getResponseBody := map[string]any{
		"RetentionPeriodHours": 178,
		"ShardCount":           3,
	}
	getResponseBodyBytes, err := json.Marshal(getResponseBody)
	require.NoError(t, err)

	cloudcontrolClient.EXPECT().GetResource(gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, params *cloudcontrol.GetResourceInput, optFns ...func(*cloudcontrol.Options)) (*cloudcontrol.GetResourceOutput, error) {
		output := cloudcontrol.GetResourceOutput{
			ResourceDescription: &types.ResourceDescription{
				Identifier: to.StringPtr(testAWSResourceName),
				Properties: to.StringPtr(string(getResponseBodyBytes)),
			},
		}
		return &output, nil
	})

	cloudcontrolClient.EXPECT().DeleteResource(gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, params *cloudcontrol.DeleteResourceInput, optFns ...func(*cloudcontrol.Options)) (*cloudcontrol.DeleteResourceOutput, error) {
		output := cloudcontrol.DeleteResourceOutput{
			ProgressEvent: &types.ProgressEvent{
				OperationStatus: types.OperationStatusSuccess,
				RequestToken:    to.StringPtr(testAWSRequestToken),
			},
		}
		return &output, nil
	})

	deleteRequest, err := http.NewRequest(http.MethodDelete, ucp.URL+basePath+testProxyRequestAWSPath, nil)
	require.NoError(t, err)
	deleteResponse, err := ucpClient.httpClient.Do(deleteRequest)
	require.NoError(t, err)

	assert.Equal(t, http.StatusAccepted, deleteResponse.StatusCode)
}
