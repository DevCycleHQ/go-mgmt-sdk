/*
 * DevCycle Management API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package go_mgmt_sdk

type NotFoundErrorResponse struct {
	// Response status code
	StatusCode float64 `json:"statusCode"`
	// Error details
	Message *interface{} `json:"message"`
	// Error type
	Error_ string `json:"error"`
}
