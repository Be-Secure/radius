//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package v20220315privatepreview

import "net/http"

// DaprSecretStoresCreateOrUpdateResponse contains the response from method DaprSecretStores.CreateOrUpdate.
type DaprSecretStoresCreateOrUpdateResponse struct {
	DaprSecretStoresCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DaprSecretStoresCreateOrUpdateResult contains the result from method DaprSecretStores.CreateOrUpdate.
type DaprSecretStoresCreateOrUpdateResult struct {
	DaprSecretStoreResource
}

// DaprSecretStoresDeleteResponse contains the response from method DaprSecretStores.Delete.
type DaprSecretStoresDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DaprSecretStoresGetResponse contains the response from method DaprSecretStores.Get.
type DaprSecretStoresGetResponse struct {
	DaprSecretStoresGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DaprSecretStoresGetResult contains the result from method DaprSecretStores.Get.
type DaprSecretStoresGetResult struct {
	DaprSecretStoreResource
}

// DaprSecretStoresListBySubscriptionResponse contains the response from method DaprSecretStores.ListBySubscription.
type DaprSecretStoresListBySubscriptionResponse struct {
	DaprSecretStoresListBySubscriptionResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DaprSecretStoresListBySubscriptionResult contains the result from method DaprSecretStores.ListBySubscription.
type DaprSecretStoresListBySubscriptionResult struct {
	DaprSecretStoreList
}

// DaprSecretStoresListResponse contains the response from method DaprSecretStores.List.
type DaprSecretStoresListResponse struct {
	DaprSecretStoresListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DaprSecretStoresListResult contains the result from method DaprSecretStores.List.
type DaprSecretStoresListResult struct {
	DaprSecretStoreList
}

// DaprStateStoresCreateOrUpdateResponse contains the response from method DaprStateStores.CreateOrUpdate.
type DaprStateStoresCreateOrUpdateResponse struct {
	DaprStateStoresCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DaprStateStoresCreateOrUpdateResult contains the result from method DaprStateStores.CreateOrUpdate.
type DaprStateStoresCreateOrUpdateResult struct {
	DaprStateStoreResource
}

// DaprStateStoresDeleteResponse contains the response from method DaprStateStores.Delete.
type DaprStateStoresDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DaprStateStoresGetResponse contains the response from method DaprStateStores.Get.
type DaprStateStoresGetResponse struct {
	DaprStateStoresGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DaprStateStoresGetResult contains the result from method DaprStateStores.Get.
type DaprStateStoresGetResult struct {
	DaprStateStoreResource
}

// DaprStateStoresListBySubscriptionResponse contains the response from method DaprStateStores.ListBySubscription.
type DaprStateStoresListBySubscriptionResponse struct {
	DaprStateStoresListBySubscriptionResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DaprStateStoresListBySubscriptionResult contains the result from method DaprStateStores.ListBySubscription.
type DaprStateStoresListBySubscriptionResult struct {
	DaprStateStoreList
}

// DaprStateStoresListResponse contains the response from method DaprStateStores.List.
type DaprStateStoresListResponse struct {
	DaprStateStoresListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DaprStateStoresListResult contains the result from method DaprStateStores.List.
type DaprStateStoresListResult struct {
	DaprStateStoreList
}

// MongoDatabasesCreateOrUpdateResponse contains the response from method MongoDatabases.CreateOrUpdate.
type MongoDatabasesCreateOrUpdateResponse struct {
	MongoDatabasesCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// MongoDatabasesCreateOrUpdateResult contains the result from method MongoDatabases.CreateOrUpdate.
type MongoDatabasesCreateOrUpdateResult struct {
	MongoDatabaseResource
}

// MongoDatabasesDeleteResponse contains the response from method MongoDatabases.Delete.
type MongoDatabasesDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// MongoDatabasesGetResponse contains the response from method MongoDatabases.Get.
type MongoDatabasesGetResponse struct {
	MongoDatabasesGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// MongoDatabasesGetResult contains the result from method MongoDatabases.Get.
type MongoDatabasesGetResult struct {
	MongoDatabaseResource
}

// MongoDatabasesListBySubscriptionResponse contains the response from method MongoDatabases.ListBySubscription.
type MongoDatabasesListBySubscriptionResponse struct {
	MongoDatabasesListBySubscriptionResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// MongoDatabasesListBySubscriptionResult contains the result from method MongoDatabases.ListBySubscription.
type MongoDatabasesListBySubscriptionResult struct {
	MongoDatabaseList
}

// MongoDatabasesListResponse contains the response from method MongoDatabases.List.
type MongoDatabasesListResponse struct {
	MongoDatabasesListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// MongoDatabasesListResult contains the result from method MongoDatabases.List.
type MongoDatabasesListResult struct {
	MongoDatabaseList
}

// MongoDatabasesListSecretsResponse contains the response from method MongoDatabases.ListSecrets.
type MongoDatabasesListSecretsResponse struct {
	MongoDatabasesListSecretsResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// MongoDatabasesListSecretsResult contains the result from method MongoDatabases.ListSecrets.
type MongoDatabasesListSecretsResult struct {
	MongoDatabaseSecrets
}

// RedisCachesCreateOrUpdateResponse contains the response from method RedisCaches.CreateOrUpdate.
type RedisCachesCreateOrUpdateResponse struct {
	RedisCachesCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// RedisCachesCreateOrUpdateResult contains the result from method RedisCaches.CreateOrUpdate.
type RedisCachesCreateOrUpdateResult struct {
	RedisCacheResource
}

// RedisCachesDeleteResponse contains the response from method RedisCaches.Delete.
type RedisCachesDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// RedisCachesGetResponse contains the response from method RedisCaches.Get.
type RedisCachesGetResponse struct {
	RedisCachesGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// RedisCachesGetResult contains the result from method RedisCaches.Get.
type RedisCachesGetResult struct {
	RedisCacheResource
}

// RedisCachesListBySubscriptionResponse contains the response from method RedisCaches.ListBySubscription.
type RedisCachesListBySubscriptionResponse struct {
	RedisCachesListBySubscriptionResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// RedisCachesListBySubscriptionResult contains the result from method RedisCaches.ListBySubscription.
type RedisCachesListBySubscriptionResult struct {
	RedisCacheList
}

// RedisCachesListResponse contains the response from method RedisCaches.List.
type RedisCachesListResponse struct {
	RedisCachesListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// RedisCachesListResult contains the result from method RedisCaches.List.
type RedisCachesListResult struct {
	RedisCacheList
}

// SQLDatabasesCreateOrUpdateResponse contains the response from method SQLDatabases.CreateOrUpdate.
type SQLDatabasesCreateOrUpdateResponse struct {
	SQLDatabasesCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SQLDatabasesCreateOrUpdateResult contains the result from method SQLDatabases.CreateOrUpdate.
type SQLDatabasesCreateOrUpdateResult struct {
	SQLDatabaseResource
}

// SQLDatabasesDeleteResponse contains the response from method SQLDatabases.Delete.
type SQLDatabasesDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SQLDatabasesGetResponse contains the response from method SQLDatabases.Get.
type SQLDatabasesGetResponse struct {
	SQLDatabasesGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SQLDatabasesGetResult contains the result from method SQLDatabases.Get.
type SQLDatabasesGetResult struct {
	SQLDatabaseResource
}

// SQLDatabasesListBySubscriptionResponse contains the response from method SQLDatabases.ListBySubscription.
type SQLDatabasesListBySubscriptionResponse struct {
	SQLDatabasesListBySubscriptionResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SQLDatabasesListBySubscriptionResult contains the result from method SQLDatabases.ListBySubscription.
type SQLDatabasesListBySubscriptionResult struct {
	SQLDatabaseList
}

// SQLDatabasesListResponse contains the response from method SQLDatabases.List.
type SQLDatabasesListResponse struct {
	SQLDatabasesListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SQLDatabasesListResult contains the result from method SQLDatabases.List.
type SQLDatabasesListResult struct {
	SQLDatabaseList
}

