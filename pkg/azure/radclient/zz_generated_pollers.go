// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package radclient

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"time"
)

// ApplicationResourcePoller provides polling facilities until the operation reaches a terminal state.
type ApplicationResourcePoller interface {
	azcore.Poller
	// FinalResponse performs a final GET to the service and returns the final response
	// for the polling operation. If there is an error performing the final GET then an error is returned.
	// If the final GET succeeded then the final ApplicationResourceResponse will be returned.
	FinalResponse(ctx context.Context) (ApplicationResourceResponse, error)
}

type applicationResourcePoller struct {
	pt *armcore.LROPoller
}

func (p *applicationResourcePoller) Done() bool {
	return p.pt.Done()
}

func (p *applicationResourcePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

func (p *applicationResourcePoller) FinalResponse(ctx context.Context) (ApplicationResourceResponse, error) {
	respType := ApplicationResourceResponse{ApplicationResource: &ApplicationResource{}}
	resp, err := p.pt.FinalResponse(ctx, respType.ApplicationResource)
	if err != nil {
		return ApplicationResourceResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

func (p *applicationResourcePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

func (p *applicationResourcePoller) pollUntilDone(ctx context.Context, freq time.Duration) (ApplicationResourceResponse, error) {
	respType := ApplicationResourceResponse{ApplicationResource: &ApplicationResource{}}
	resp, err := p.pt.PollUntilDone(ctx, freq, respType.ApplicationResource)
	if err != nil {
		return ApplicationResourceResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// AzureKeyVaultComponentResourcePoller provides polling facilities until the operation reaches a terminal state.
type AzureKeyVaultComponentResourcePoller interface {
	azcore.Poller
	// FinalResponse performs a final GET to the service and returns the final response
	// for the polling operation. If there is an error performing the final GET then an error is returned.
	// If the final GET succeeded then the final AzureKeyVaultComponentResourceResponse will be returned.
	FinalResponse(ctx context.Context) (AzureKeyVaultComponentResourceResponse, error)
}

type azureKeyVaultComponentResourcePoller struct {
	pt *armcore.LROPoller
}

func (p *azureKeyVaultComponentResourcePoller) Done() bool {
	return p.pt.Done()
}

func (p *azureKeyVaultComponentResourcePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

func (p *azureKeyVaultComponentResourcePoller) FinalResponse(ctx context.Context) (AzureKeyVaultComponentResourceResponse, error) {
	respType := AzureKeyVaultComponentResourceResponse{AzureKeyVaultComponentResource: &AzureKeyVaultComponentResource{}}
	resp, err := p.pt.FinalResponse(ctx, respType.AzureKeyVaultComponentResource)
	if err != nil {
		return AzureKeyVaultComponentResourceResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

func (p *azureKeyVaultComponentResourcePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

func (p *azureKeyVaultComponentResourcePoller) pollUntilDone(ctx context.Context, freq time.Duration) (AzureKeyVaultComponentResourceResponse, error) {
	respType := AzureKeyVaultComponentResourceResponse{AzureKeyVaultComponentResource: &AzureKeyVaultComponentResource{}}
	resp, err := p.pt.PollUntilDone(ctx, freq, respType.AzureKeyVaultComponentResource)
	if err != nil {
		return AzureKeyVaultComponentResourceResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// AzureServiceBusComponentResourcePoller provides polling facilities until the operation reaches a terminal state.
type AzureServiceBusComponentResourcePoller interface {
	azcore.Poller
	// FinalResponse performs a final GET to the service and returns the final response
	// for the polling operation. If there is an error performing the final GET then an error is returned.
	// If the final GET succeeded then the final AzureServiceBusComponentResourceResponse will be returned.
	FinalResponse(ctx context.Context) (AzureServiceBusComponentResourceResponse, error)
}

type azureServiceBusComponentResourcePoller struct {
	pt *armcore.LROPoller
}

func (p *azureServiceBusComponentResourcePoller) Done() bool {
	return p.pt.Done()
}

func (p *azureServiceBusComponentResourcePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

func (p *azureServiceBusComponentResourcePoller) FinalResponse(ctx context.Context) (AzureServiceBusComponentResourceResponse, error) {
	respType := AzureServiceBusComponentResourceResponse{AzureServiceBusComponentResource: &AzureServiceBusComponentResource{}}
	resp, err := p.pt.FinalResponse(ctx, respType.AzureServiceBusComponentResource)
	if err != nil {
		return AzureServiceBusComponentResourceResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

func (p *azureServiceBusComponentResourcePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

func (p *azureServiceBusComponentResourcePoller) pollUntilDone(ctx context.Context, freq time.Duration) (AzureServiceBusComponentResourceResponse, error) {
	respType := AzureServiceBusComponentResourceResponse{AzureServiceBusComponentResource: &AzureServiceBusComponentResource{}}
	resp, err := p.pt.PollUntilDone(ctx, freq, respType.AzureServiceBusComponentResource)
	if err != nil {
		return AzureServiceBusComponentResourceResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ContainerComponentResourcePoller provides polling facilities until the operation reaches a terminal state.
type ContainerComponentResourcePoller interface {
	azcore.Poller
	// FinalResponse performs a final GET to the service and returns the final response
	// for the polling operation. If there is an error performing the final GET then an error is returned.
	// If the final GET succeeded then the final ContainerComponentResourceResponse will be returned.
	FinalResponse(ctx context.Context) (ContainerComponentResourceResponse, error)
}

type containerComponentResourcePoller struct {
	pt *armcore.LROPoller
}

func (p *containerComponentResourcePoller) Done() bool {
	return p.pt.Done()
}

func (p *containerComponentResourcePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

func (p *containerComponentResourcePoller) FinalResponse(ctx context.Context) (ContainerComponentResourceResponse, error) {
	respType := ContainerComponentResourceResponse{ContainerComponentResource: &ContainerComponentResource{}}
	resp, err := p.pt.FinalResponse(ctx, respType.ContainerComponentResource)
	if err != nil {
		return ContainerComponentResourceResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

func (p *containerComponentResourcePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

func (p *containerComponentResourcePoller) pollUntilDone(ctx context.Context, freq time.Duration) (ContainerComponentResourceResponse, error) {
	respType := ContainerComponentResourceResponse{ContainerComponentResource: &ContainerComponentResource{}}
	resp, err := p.pt.PollUntilDone(ctx, freq, respType.ContainerComponentResource)
	if err != nil {
		return ContainerComponentResourceResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// DaprHTTPRouteResourcePoller provides polling facilities until the operation reaches a terminal state.
type DaprHTTPRouteResourcePoller interface {
	azcore.Poller
	// FinalResponse performs a final GET to the service and returns the final response
	// for the polling operation. If there is an error performing the final GET then an error is returned.
	// If the final GET succeeded then the final DaprHTTPRouteResourceResponse will be returned.
	FinalResponse(ctx context.Context) (DaprHTTPRouteResourceResponse, error)
}

type daprHTTPRouteResourcePoller struct {
	pt *armcore.LROPoller
}

func (p *daprHTTPRouteResourcePoller) Done() bool {
	return p.pt.Done()
}

func (p *daprHTTPRouteResourcePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

func (p *daprHTTPRouteResourcePoller) FinalResponse(ctx context.Context) (DaprHTTPRouteResourceResponse, error) {
	respType := DaprHTTPRouteResourceResponse{DaprHTTPRouteResource: &DaprHTTPRouteResource{}}
	resp, err := p.pt.FinalResponse(ctx, respType.DaprHTTPRouteResource)
	if err != nil {
		return DaprHTTPRouteResourceResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

func (p *daprHTTPRouteResourcePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

func (p *daprHTTPRouteResourcePoller) pollUntilDone(ctx context.Context, freq time.Duration) (DaprHTTPRouteResourceResponse, error) {
	respType := DaprHTTPRouteResourceResponse{DaprHTTPRouteResource: &DaprHTTPRouteResource{}}
	resp, err := p.pt.PollUntilDone(ctx, freq, respType.DaprHTTPRouteResource)
	if err != nil {
		return DaprHTTPRouteResourceResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// DaprPubSubTopicComponentResourcePoller provides polling facilities until the operation reaches a terminal state.
type DaprPubSubTopicComponentResourcePoller interface {
	azcore.Poller
	// FinalResponse performs a final GET to the service and returns the final response
	// for the polling operation. If there is an error performing the final GET then an error is returned.
	// If the final GET succeeded then the final DaprPubSubTopicComponentResourceResponse will be returned.
	FinalResponse(ctx context.Context) (DaprPubSubTopicComponentResourceResponse, error)
}

type daprPubSubTopicComponentResourcePoller struct {
	pt *armcore.LROPoller
}

func (p *daprPubSubTopicComponentResourcePoller) Done() bool {
	return p.pt.Done()
}

func (p *daprPubSubTopicComponentResourcePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

func (p *daprPubSubTopicComponentResourcePoller) FinalResponse(ctx context.Context) (DaprPubSubTopicComponentResourceResponse, error) {
	respType := DaprPubSubTopicComponentResourceResponse{DaprPubSubTopicComponentResource: &DaprPubSubTopicComponentResource{}}
	resp, err := p.pt.FinalResponse(ctx, respType.DaprPubSubTopicComponentResource)
	if err != nil {
		return DaprPubSubTopicComponentResourceResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

func (p *daprPubSubTopicComponentResourcePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

func (p *daprPubSubTopicComponentResourcePoller) pollUntilDone(ctx context.Context, freq time.Duration) (DaprPubSubTopicComponentResourceResponse, error) {
	respType := DaprPubSubTopicComponentResourceResponse{DaprPubSubTopicComponentResource: &DaprPubSubTopicComponentResource{}}
	resp, err := p.pt.PollUntilDone(ctx, freq, respType.DaprPubSubTopicComponentResource)
	if err != nil {
		return DaprPubSubTopicComponentResourceResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// DaprStateStoreComponentResourcePoller provides polling facilities until the operation reaches a terminal state.
type DaprStateStoreComponentResourcePoller interface {
	azcore.Poller
	// FinalResponse performs a final GET to the service and returns the final response
	// for the polling operation. If there is an error performing the final GET then an error is returned.
	// If the final GET succeeded then the final DaprStateStoreComponentResourceResponse will be returned.
	FinalResponse(ctx context.Context) (DaprStateStoreComponentResourceResponse, error)
}

type daprStateStoreComponentResourcePoller struct {
	pt *armcore.LROPoller
}

func (p *daprStateStoreComponentResourcePoller) Done() bool {
	return p.pt.Done()
}

func (p *daprStateStoreComponentResourcePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

func (p *daprStateStoreComponentResourcePoller) FinalResponse(ctx context.Context) (DaprStateStoreComponentResourceResponse, error) {
	respType := DaprStateStoreComponentResourceResponse{DaprStateStoreComponentResource: &DaprStateStoreComponentResource{}}
	resp, err := p.pt.FinalResponse(ctx, respType.DaprStateStoreComponentResource)
	if err != nil {
		return DaprStateStoreComponentResourceResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

func (p *daprStateStoreComponentResourcePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

func (p *daprStateStoreComponentResourcePoller) pollUntilDone(ctx context.Context, freq time.Duration) (DaprStateStoreComponentResourceResponse, error) {
	respType := DaprStateStoreComponentResourceResponse{DaprStateStoreComponentResource: &DaprStateStoreComponentResource{}}
	resp, err := p.pt.PollUntilDone(ctx, freq, respType.DaprStateStoreComponentResource)
	if err != nil {
		return DaprStateStoreComponentResourceResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// HTTPPoller provides polling facilities until the operation reaches a terminal state.
type HTTPPoller interface {
	azcore.Poller
	// FinalResponse performs a final GET to the service and returns the final response
	// for the polling operation. If there is an error performing the final GET then an error is returned.
	// If the final GET succeeded then the final *http.Response will be returned.
	FinalResponse(ctx context.Context) (*http.Response, error)
}

type httpPoller struct {
	pt *armcore.LROPoller
}

func (p *httpPoller) Done() bool {
	return p.pt.Done()
}

func (p *httpPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

func (p *httpPoller) FinalResponse(ctx context.Context) (*http.Response, error) {
	return p.pt.FinalResponse(ctx, nil)
}

func (p *httpPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

func (p *httpPoller) pollUntilDone(ctx context.Context, freq time.Duration) (*http.Response, error) {
	return p.pt.PollUntilDone(ctx, freq, nil)
}

// HTTPRouteResourcePoller provides polling facilities until the operation reaches a terminal state.
type HTTPRouteResourcePoller interface {
	azcore.Poller
	// FinalResponse performs a final GET to the service and returns the final response
	// for the polling operation. If there is an error performing the final GET then an error is returned.
	// If the final GET succeeded then the final HTTPRouteResourceResponse will be returned.
	FinalResponse(ctx context.Context) (HTTPRouteResourceResponse, error)
}

type httpRouteResourcePoller struct {
	pt *armcore.LROPoller
}

func (p *httpRouteResourcePoller) Done() bool {
	return p.pt.Done()
}

func (p *httpRouteResourcePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

func (p *httpRouteResourcePoller) FinalResponse(ctx context.Context) (HTTPRouteResourceResponse, error) {
	respType := HTTPRouteResourceResponse{HTTPRouteResource: &HTTPRouteResource{}}
	resp, err := p.pt.FinalResponse(ctx, respType.HTTPRouteResource)
	if err != nil {
		return HTTPRouteResourceResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

func (p *httpRouteResourcePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

func (p *httpRouteResourcePoller) pollUntilDone(ctx context.Context, freq time.Duration) (HTTPRouteResourceResponse, error) {
	respType := HTTPRouteResourceResponse{HTTPRouteResource: &HTTPRouteResource{}}
	resp, err := p.pt.PollUntilDone(ctx, freq, respType.HTTPRouteResource)
	if err != nil {
		return HTTPRouteResourceResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// MicrosoftSQLComponentResourcePoller provides polling facilities until the operation reaches a terminal state.
type MicrosoftSQLComponentResourcePoller interface {
	azcore.Poller
	// FinalResponse performs a final GET to the service and returns the final response
	// for the polling operation. If there is an error performing the final GET then an error is returned.
	// If the final GET succeeded then the final MicrosoftSQLComponentResourceResponse will be returned.
	FinalResponse(ctx context.Context) (MicrosoftSQLComponentResourceResponse, error)
}

type microsoftSQLComponentResourcePoller struct {
	pt *armcore.LROPoller
}

func (p *microsoftSQLComponentResourcePoller) Done() bool {
	return p.pt.Done()
}

func (p *microsoftSQLComponentResourcePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

func (p *microsoftSQLComponentResourcePoller) FinalResponse(ctx context.Context) (MicrosoftSQLComponentResourceResponse, error) {
	respType := MicrosoftSQLComponentResourceResponse{MicrosoftSQLComponentResource: &MicrosoftSQLComponentResource{}}
	resp, err := p.pt.FinalResponse(ctx, respType.MicrosoftSQLComponentResource)
	if err != nil {
		return MicrosoftSQLComponentResourceResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

func (p *microsoftSQLComponentResourcePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

func (p *microsoftSQLComponentResourcePoller) pollUntilDone(ctx context.Context, freq time.Duration) (MicrosoftSQLComponentResourceResponse, error) {
	respType := MicrosoftSQLComponentResourceResponse{MicrosoftSQLComponentResource: &MicrosoftSQLComponentResource{}}
	resp, err := p.pt.PollUntilDone(ctx, freq, respType.MicrosoftSQLComponentResource)
	if err != nil {
		return MicrosoftSQLComponentResourceResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// MongoDBComponentResourcePoller provides polling facilities until the operation reaches a terminal state.
type MongoDBComponentResourcePoller interface {
	azcore.Poller
	// FinalResponse performs a final GET to the service and returns the final response
	// for the polling operation. If there is an error performing the final GET then an error is returned.
	// If the final GET succeeded then the final MongoDBComponentResourceResponse will be returned.
	FinalResponse(ctx context.Context) (MongoDBComponentResourceResponse, error)
}

type mongoDBComponentResourcePoller struct {
	pt *armcore.LROPoller
}

func (p *mongoDBComponentResourcePoller) Done() bool {
	return p.pt.Done()
}

func (p *mongoDBComponentResourcePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

func (p *mongoDBComponentResourcePoller) FinalResponse(ctx context.Context) (MongoDBComponentResourceResponse, error) {
	respType := MongoDBComponentResourceResponse{MongoDBComponentResource: &MongoDBComponentResource{}}
	resp, err := p.pt.FinalResponse(ctx, respType.MongoDBComponentResource)
	if err != nil {
		return MongoDBComponentResourceResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

func (p *mongoDBComponentResourcePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

func (p *mongoDBComponentResourcePoller) pollUntilDone(ctx context.Context, freq time.Duration) (MongoDBComponentResourceResponse, error) {
	respType := MongoDBComponentResourceResponse{MongoDBComponentResource: &MongoDBComponentResource{}}
	resp, err := p.pt.PollUntilDone(ctx, freq, respType.MongoDBComponentResource)
	if err != nil {
		return MongoDBComponentResourceResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// RabbitMQComponentResourcePoller provides polling facilities until the operation reaches a terminal state.
type RabbitMQComponentResourcePoller interface {
	azcore.Poller
	// FinalResponse performs a final GET to the service and returns the final response
	// for the polling operation. If there is an error performing the final GET then an error is returned.
	// If the final GET succeeded then the final RabbitMQComponentResourceResponse will be returned.
	FinalResponse(ctx context.Context) (RabbitMQComponentResourceResponse, error)
}

type rabbitMQComponentResourcePoller struct {
	pt *armcore.LROPoller
}

func (p *rabbitMQComponentResourcePoller) Done() bool {
	return p.pt.Done()
}

func (p *rabbitMQComponentResourcePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

func (p *rabbitMQComponentResourcePoller) FinalResponse(ctx context.Context) (RabbitMQComponentResourceResponse, error) {
	respType := RabbitMQComponentResourceResponse{RabbitMQComponentResource: &RabbitMQComponentResource{}}
	resp, err := p.pt.FinalResponse(ctx, respType.RabbitMQComponentResource)
	if err != nil {
		return RabbitMQComponentResourceResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

func (p *rabbitMQComponentResourcePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

func (p *rabbitMQComponentResourcePoller) pollUntilDone(ctx context.Context, freq time.Duration) (RabbitMQComponentResourceResponse, error) {
	respType := RabbitMQComponentResourceResponse{RabbitMQComponentResource: &RabbitMQComponentResource{}}
	resp, err := p.pt.PollUntilDone(ctx, freq, respType.RabbitMQComponentResource)
	if err != nil {
		return RabbitMQComponentResourceResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// RedisComponentResourcePoller provides polling facilities until the operation reaches a terminal state.
type RedisComponentResourcePoller interface {
	azcore.Poller
	// FinalResponse performs a final GET to the service and returns the final response
	// for the polling operation. If there is an error performing the final GET then an error is returned.
	// If the final GET succeeded then the final RedisComponentResourceResponse will be returned.
	FinalResponse(ctx context.Context) (RedisComponentResourceResponse, error)
}

type redisComponentResourcePoller struct {
	pt *armcore.LROPoller
}

func (p *redisComponentResourcePoller) Done() bool {
	return p.pt.Done()
}

func (p *redisComponentResourcePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

func (p *redisComponentResourcePoller) FinalResponse(ctx context.Context) (RedisComponentResourceResponse, error) {
	respType := RedisComponentResourceResponse{RedisComponentResource: &RedisComponentResource{}}
	resp, err := p.pt.FinalResponse(ctx, respType.RedisComponentResource)
	if err != nil {
		return RedisComponentResourceResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

func (p *redisComponentResourcePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

func (p *redisComponentResourcePoller) pollUntilDone(ctx context.Context, freq time.Duration) (RedisComponentResourceResponse, error) {
	respType := RedisComponentResourceResponse{RedisComponentResource: &RedisComponentResource{}}
	resp, err := p.pt.PollUntilDone(ctx, freq, respType.RedisComponentResource)
	if err != nil {
		return RedisComponentResourceResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// VolumeResourcePoller provides polling facilities until the operation reaches a terminal state.
type VolumeResourcePoller interface {
	azcore.Poller
	// FinalResponse performs a final GET to the service and returns the final response
	// for the polling operation. If there is an error performing the final GET then an error is returned.
	// If the final GET succeeded then the final VolumeResourceResponse will be returned.
	FinalResponse(ctx context.Context) (VolumeResourceResponse, error)
}

type volumeResourcePoller struct {
	pt *armcore.LROPoller
}

func (p *volumeResourcePoller) Done() bool {
	return p.pt.Done()
}

func (p *volumeResourcePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

func (p *volumeResourcePoller) FinalResponse(ctx context.Context) (VolumeResourceResponse, error) {
	respType := VolumeResourceResponse{VolumeResource: &VolumeResource{}}
	resp, err := p.pt.FinalResponse(ctx, respType.VolumeResource)
	if err != nil {
		return VolumeResourceResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

func (p *volumeResourcePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

func (p *volumeResourcePoller) pollUntilDone(ctx context.Context, freq time.Duration) (VolumeResourceResponse, error) {
	respType := VolumeResourceResponse{VolumeResource: &VolumeResource{}}
	resp, err := p.pt.PollUntilDone(ctx, freq, respType.VolumeResource)
	if err != nil {
		return VolumeResourceResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}
