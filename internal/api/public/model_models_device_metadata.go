/*
Nexodus API

This is the Nexodus API Server.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package public

// ModelsDeviceMetadata struct for ModelsDeviceMetadata
type ModelsDeviceMetadata struct {
	DeviceId string                               `json:"device_id,omitempty"`
	Metadata map[string]ModelsDeviceMetadataValue `json:"metadata,omitempty"`
}
