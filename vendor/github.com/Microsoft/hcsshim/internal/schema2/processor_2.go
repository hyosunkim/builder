/*
 * HCS API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package hcsschema

type Processor2 struct {
	Count int32 `json:"Count,omitempty"`

	Limit int32 `json:"Limit,omitempty"`

	Weight int32 `json:"Weight,omitempty"`

	ExposeVirtualizationExtensions bool `json:"ExposeVirtualizationExtensions,omitempty"`
}
