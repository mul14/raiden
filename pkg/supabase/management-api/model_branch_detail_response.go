/*
 * Supabase API (v1)
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package management_api

type BranchDetailResponse struct {
	DbPort          int32  `json:"db_port"`
	Ref             string `json:"ref"`
	PostgresVersion string `json:"postgres_version"`
	Status          string `json:"status"`
	DbHost          string `json:"db_host"`
	DbUser          string `json:"db_user,omitempty"`
	DbPass          string `json:"db_pass,omitempty"`
	JwtSecret       string `json:"jwt_secret,omitempty"`
}
