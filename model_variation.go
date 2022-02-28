/*
 * DevCycle Management API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package go_mgmt_sdk

import "github.com/antihax/optional"

type Variation struct {
	// Unique key by Feature, can be used in the SDK / API to reference by 'key' rather than _id. Must only contain lower-case characters and `_` or `-`.
	Key string `json:"key"`
	// Variation display name.
	Name string `json:"name"`
	// A key-value map of variables to their value for this variation
	Variables map[string]optional.Interface `json:"variables,omitempty"`
	// A unique Variation ID
	Id string `json:"_id"`
}
