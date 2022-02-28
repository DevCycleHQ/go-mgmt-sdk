/*
 * DevCycle Management API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package go_mgmt_sdk

type CreateProjectDto struct {
	// Project name
	Name string `json:"name"`
	// A unique key to identify the Project
	Key string `json:"key"`
	// A description of the Project
	Description string `json:"description,omitempty"`
}
