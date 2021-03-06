/*
 * DevCycle Management API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package go_mgmt_sdk

type UpdateFeatureConfigDto struct {
	// The targets to evaluate what variation a user should be delivered
	Targets []UpdateTargetDto `json:"targets"`
	// Status of the Feature Configuration
	Status string `json:"status,omitempty"`
}
