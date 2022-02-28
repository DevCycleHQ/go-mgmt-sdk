/*
 * DevCycle Management API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package go_mgmt_sdk

type TargetAudience struct {
	// Audience display name, must be set for project-level audiences.
	Name string `json:"name,omitempty"`
	// Audience filters, describing logic for segmenting users
	Filters *AllOfTargetAudienceFilters `json:"filters"`
}
