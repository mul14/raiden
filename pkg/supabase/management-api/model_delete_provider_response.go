/*
 * Supabase API (v1)
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package management_api

type DeleteProviderResponse struct {
	Id        string          `json:"id"`
	Saml      *SamlDescriptor `json:"saml,omitempty"`
	Domains   []Domain        `json:"domains,omitempty"`
	CreatedAt string          `json:"created_at,omitempty"`
	UpdatedAt string          `json:"updated_at,omitempty"`
}
