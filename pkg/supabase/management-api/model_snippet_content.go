/*
 * Supabase API (v1)
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package management_api

type SnippetContent struct {
	Favorite      bool   `json:"favorite"`
	SchemaVersion string `json:"schema_version"`
	Sql           string `json:"sql"`
}
