/*
 * Supabase API (v1)
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package management_api

type V1Backup struct {
	Status           string `json:"status"`
	IsPhysicalBackup bool   `json:"is_physical_backup"`
	InsertedAt       string `json:"inserted_at"`
}
