/*
 * Supabase API (v1)
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package management_api

type CreateProjectBody struct {
	// Database password
	DbPass string `json:"db_pass"`
	// Name of your project, should not contain dots
	Name string `json:"name"`
	// Slug of your organization
	OrganizationId string `json:"organization_id"`
	// Subscription plan is now set on organization level and is ignored in this request
	Plan string `json:"plan"`
	// Region you want your server to reside in
	Region     string `json:"region"`
	KpsEnabled bool   `json:"kps_enabled,omitempty"`
}
