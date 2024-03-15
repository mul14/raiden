/*
 * Supabase API (v1)
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package management_api

type PostgrestConfigResponse struct {
	MaxRows           int32  `json:"max_rows"`
	DbSchema          string `json:"db_schema"`
	DbExtraSearchPath string `json:"db_extra_search_path"`
}