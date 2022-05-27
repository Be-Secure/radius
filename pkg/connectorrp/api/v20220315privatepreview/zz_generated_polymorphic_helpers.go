//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package v20220315privatepreview

import "encoding/json"

func unmarshalDaprStateStorePropertiesClassification(rawMsg json.RawMessage) (DaprStateStorePropertiesClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b DaprStateStorePropertiesClassification
	switch m["kind"] {
	case "generic":
		b = &DaprStateStoreGenericResourceProperties{}
	case "state.azure.tablestorage":
		b = &DaprStateStoreAzureTableStorageResourceProperties{}
	case "state.sqlserver":
		b = &DaprStateStoreSQLServerResourceProperties{}
	default:
		b = &DaprStateStoreProperties{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalDaprStateStorePropertiesClassificationArray(rawMsg json.RawMessage) ([]DaprStateStorePropertiesClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]DaprStateStorePropertiesClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalDaprStateStorePropertiesClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalDaprStateStorePropertiesClassificationMap(rawMsg json.RawMessage) (map[string]DaprStateStorePropertiesClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages map[string]json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fMap := make(map[string]DaprStateStorePropertiesClassification, len(rawMessages))
	for key, rawMessage := range rawMessages {
		f, err := unmarshalDaprStateStorePropertiesClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fMap[key] = f
	}
	return fMap, nil
}

