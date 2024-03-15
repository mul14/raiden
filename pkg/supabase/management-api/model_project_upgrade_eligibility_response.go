/*
 * Supabase API (v1)
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package management_api

type ProjectUpgradeEligibilityResponse struct {
	Eligible                   bool             `json:"eligible"`
	CurrentAppVersion          string           `json:"current_app_version"`
	LatestAppVersion           string           `json:"latest_app_version"`
	TargetUpgradeVersions      []ProjectVersion `json:"target_upgrade_versions"`
	RequiresManualIntervention string           `json:"requires_manual_intervention"`
	PotentialBreakingChanges   []string         `json:"potential_breaking_changes"`
	DurationEstimateHours      float64          `json:"duration_estimate_hours"`
	LegacyAuthCustomRoles      []string         `json:"legacy_auth_custom_roles"`
	ExtensionDependentObjects  []string         `json:"extension_dependent_objects"`
}