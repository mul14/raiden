# Go API client for swagger

No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 1.0.0
- Package version: 1.0.0
- Build package: io.swagger.codegen.v3.generators.go.GoClientCodegen

## Installation
Put the package under your project folder and add the following in import:
```golang
import "./swagger"
```

## Documentation for API Endpoints

All URIs are relative to */*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*BackupsApi* | [**GetBackups**](docs/BackupsApi.md#getbackups) | **Get** /v1/projects/{ref}/database/backups | Lists all backups
*BackupsApi* | [**V1RestorePitr**](docs/BackupsApi.md#v1restorepitr) | **Post** /v1/projects/{ref}/database/backups/restore-pitr | Restores a PITR backup for a database
*CustomHostnameBetaApi* | [**Activate**](docs/CustomHostnameBetaApi.md#activate) | **Post** /v1/projects/{ref}/custom-hostname/activate | Activates a custom hostname for a project.
*CustomHostnameBetaApi* | [**CreateCustomHostnameConfig**](docs/CustomHostnameBetaApi.md#createcustomhostnameconfig) | **Post** /v1/projects/{ref}/custom-hostname/initialize | Updates project&#x27;s custom hostname configuration
*CustomHostnameBetaApi* | [**GetCustomHostnameConfig**](docs/CustomHostnameBetaApi.md#getcustomhostnameconfig) | **Get** /v1/projects/{ref}/custom-hostname | Gets project&#x27;s custom hostname config
*CustomHostnameBetaApi* | [**RemoveCustomHostnameConfig**](docs/CustomHostnameBetaApi.md#removecustomhostnameconfig) | **Delete** /v1/projects/{ref}/custom-hostname | Deletes a project&#x27;s custom hostname configuration
*CustomHostnameBetaApi* | [**Reverify**](docs/CustomHostnameBetaApi.md#reverify) | **Post** /v1/projects/{ref}/custom-hostname/reverify | Attempts to verify the DNS configuration for project&#x27;s custom hostname configuration
*DatabaseBranchesBetaApi* | [**CreateBranch**](docs/DatabaseBranchesBetaApi.md#createbranch) | **Post** /v1/projects/{ref}/branches | Create a database branch
*DatabaseBranchesBetaApi* | [**DeleteBranch**](docs/DatabaseBranchesBetaApi.md#deletebranch) | **Delete** /v1/branches/{branch_id} | Delete a database branch
*DatabaseBranchesBetaApi* | [**DisableBranch**](docs/DatabaseBranchesBetaApi.md#disablebranch) | **Delete** /v1/projects/{ref}/branches | Disables preview branching
*DatabaseBranchesBetaApi* | [**GetBranchDetails**](docs/DatabaseBranchesBetaApi.md#getbranchdetails) | **Get** /v1/branches/{branch_id} | Get database branch config
*DatabaseBranchesBetaApi* | [**GetBranches**](docs/DatabaseBranchesBetaApi.md#getbranches) | **Get** /v1/projects/{ref}/branches | List all database branches
*DatabaseBranchesBetaApi* | [**UpdateBranch**](docs/DatabaseBranchesBetaApi.md#updatebranch) | **Patch** /v1/branches/{branch_id} | Update database branch config
*DatabaseReadonlyModeApi* | [**GetReadOnlyModeStatus**](docs/DatabaseReadonlyModeApi.md#getreadonlymodestatus) | **Get** /v1/projects/{ref}/readonly | Returns project&#x27;s readonly mode status
*DatabaseReadonlyModeApi* | [**TemporarilyDisableReadonlyMode**](docs/DatabaseReadonlyModeApi.md#temporarilydisablereadonlymode) | **Post** /v1/projects/{ref}/readonly/temporary-disable | Disables project&#x27;s readonly mode for the next 15 minutes
*DatabaseVersionUpgradeBetaApi* | [**GetUpgradeStatus**](docs/DatabaseVersionUpgradeBetaApi.md#getupgradestatus) | **Get** /v1/projects/{ref}/upgrade/status | Gets the latest status of the project&#x27;s upgrade
*DatabaseVersionUpgradeBetaApi* | [**UpgradeEligibilityInformation**](docs/DatabaseVersionUpgradeBetaApi.md#upgradeeligibilityinformation) | **Get** /v1/projects/{ref}/upgrade/eligibility | Returns the project&#x27;s eligibility for upgrades
*DatabaseVersionUpgradeBetaApi* | [**UpgradeProject**](docs/DatabaseVersionUpgradeBetaApi.md#upgradeproject) | **Post** /v1/projects/{ref}/upgrade | Upgrades the project&#x27;s Postgres version
*FunctionsApi* | [**CreateFunction**](docs/FunctionsApi.md#createfunction) | **Post** /v1/projects/{ref}/functions | Create a function
*FunctionsApi* | [**DeleteFunction**](docs/FunctionsApi.md#deletefunction) | **Delete** /v1/projects/{ref}/functions/{function_slug} | Delete a function
*FunctionsApi* | [**GetFunction**](docs/FunctionsApi.md#getfunction) | **Get** /v1/projects/{ref}/functions/{function_slug} | Retrieve a function
*FunctionsApi* | [**GetFunctionBody**](docs/FunctionsApi.md#getfunctionbody) | **Get** /v1/projects/{ref}/functions/{function_slug}/body | Retrieve a function body
*FunctionsApi* | [**GetFunctions**](docs/FunctionsApi.md#getfunctions) | **Get** /v1/projects/{ref}/functions | List all functions
*FunctionsApi* | [**UpdateFunction**](docs/FunctionsApi.md#updatefunction) | **Patch** /v1/projects/{ref}/functions/{function_slug} | Update a function
*NetworkBansBetaApi* | [**GetNetworkBans**](docs/NetworkBansBetaApi.md#getnetworkbans) | **Post** /v1/projects/{ref}/network-bans/retrieve | Gets project&#x27;s network bans
*NetworkBansBetaApi* | [**RemoveNetworkBan**](docs/NetworkBansBetaApi.md#removenetworkban) | **Delete** /v1/projects/{ref}/network-bans | Remove network bans.
*NetworkRestrictionsBetaApi* | [**ApplyNetworkRestrictions**](docs/NetworkRestrictionsBetaApi.md#applynetworkrestrictions) | **Post** /v1/projects/{ref}/network-restrictions/apply | Updates project&#x27;s network restrictions
*NetworkRestrictionsBetaApi* | [**GetNetworkRestrictions**](docs/NetworkRestrictionsBetaApi.md#getnetworkrestrictions) | **Get** /v1/projects/{ref}/network-restrictions | Gets project&#x27;s network restrictions
*OauthBetaApi* | [**Authorize**](docs/OauthBetaApi.md#authorize) | **Get** /v1/oauth/authorize | Authorize user through oauth
*OauthBetaApi* | [**Token**](docs/OauthBetaApi.md#token) | **Post** /v1/oauth/token | Exchange auth code for user&#x27;s access and refresh token
*OrganizationsApi* | [**CreateOrganization**](docs/OrganizationsApi.md#createorganization) | **Post** /v1/organizations | Create an organization
*OrganizationsApi* | [**GetOrganization**](docs/OrganizationsApi.md#getorganization) | **Get** /v1/organizations/{slug} | Gets information about the organization
*OrganizationsApi* | [**GetOrganizations**](docs/OrganizationsApi.md#getorganizations) | **Get** /v1/organizations | List all organizations
*OrganizationsApi* | [**V1ListOrganizationMembers**](docs/OrganizationsApi.md#v1listorganizationmembers) | **Get** /v1/organizations/{slug}/members | List members of an organization
*PgsodiumBetaApi* | [**GetPgsodiumConfig**](docs/PgsodiumBetaApi.md#getpgsodiumconfig) | **Get** /v1/projects/{ref}/pgsodium | Gets project&#x27;s pgsodium config
*PgsodiumBetaApi* | [**UpdatePgsodiumConfig**](docs/PgsodiumBetaApi.md#updatepgsodiumconfig) | **Put** /v1/projects/{ref}/pgsodium | Updates project&#x27;s pgsodium config. Updating the root_key can cause all data encrypted with the older key to become inaccessible.
*ProjectsApi* | [**CreateProject**](docs/ProjectsApi.md#createproject) | **Post** /v1/projects | Create a project
*ProjectsApi* | [**DeleteProject**](docs/ProjectsApi.md#deleteproject) | **Delete** /v1/projects/{ref} | Deletes the given project
*ProjectsApi* | [**GetProjectApiKeys**](docs/ProjectsApi.md#getprojectapikeys) | **Get** /v1/projects/{ref}/api-keys | Get project api keys
*ProjectsApi* | [**GetProjects**](docs/ProjectsApi.md#getprojects) | **Get** /v1/projects | List all projects
*ProjectsApi* | [**GetTypescriptTypes**](docs/ProjectsApi.md#gettypescripttypes) | **Get** /v1/projects/{ref}/types/typescript | Generate TypeScript types
*ProjectsBetaApi* | [**V1EnableDatabaseWebhooks**](docs/ProjectsBetaApi.md#v1enabledatabasewebhooks) | **Post** /v1/projects/{ref}/database/webhooks/enable | Enables Database Webhooks on the project
*ProjectsBetaApi* | [**V1RunQuery**](docs/ProjectsBetaApi.md#v1runquery) | **Post** /v1/projects/{ref}/database/query | Run sql query
*ProjectsConfigApi* | [**GetConfig**](docs/ProjectsConfigApi.md#getconfig) | **Get** /v1/projects/{ref}/config/database/postgres | Gets project&#x27;s Postgres config
*ProjectsConfigApi* | [**GetV1AuthConfig**](docs/ProjectsConfigApi.md#getv1authconfig) | **Get** /v1/projects/{ref}/config/auth | Gets project&#x27;s auth config
*ProjectsConfigApi* | [**UpdateConfig**](docs/ProjectsConfigApi.md#updateconfig) | **Put** /v1/projects/{ref}/config/database/postgres | Updates project&#x27;s Postgres config
*ProjectsConfigApi* | [**UpdateV1AuthConfig**](docs/ProjectsConfigApi.md#updatev1authconfig) | **Patch** /v1/projects/{ref}/config/auth | Updates a project&#x27;s auth config
*ProjectsConfigApi* | [**V1GetPgbouncerConfig**](docs/ProjectsConfigApi.md#v1getpgbouncerconfig) | **Get** /v1/projects/{ref}/config/database/pgbouncer | Get project&#x27;s pgbouncer config
*ReadReplicasBetaApi* | [**RemoveReadReplica**](docs/ReadReplicasBetaApi.md#removereadreplica) | **Post** /v1/projects/{ref}/read-replicas/remove | Remove a read replica
*ReadReplicasBetaApi* | [**SetUpReadReplica**](docs/ReadReplicasBetaApi.md#setupreadreplica) | **Post** /v1/projects/{ref}/read-replicas/setup | Set up a read replica
*SecretsApi* | [**CreateSecrets**](docs/SecretsApi.md#createsecrets) | **Post** /v1/projects/{ref}/secrets | Bulk create secrets
*SecretsApi* | [**DeleteSecrets**](docs/SecretsApi.md#deletesecrets) | **Delete** /v1/projects/{ref}/secrets | Bulk delete secrets
*SecretsApi* | [**GetSecrets**](docs/SecretsApi.md#getsecrets) | **Get** /v1/projects/{ref}/secrets | List all secrets
*ServicesApi* | [**CheckServiceHealth**](docs/ServicesApi.md#checkservicehealth) | **Get** /v1/projects/{ref}/health | Gets project&#x27;s service health status
*ServicesApi* | [**GetPostgRESTConfig**](docs/ServicesApi.md#getpostgrestconfig) | **Get** /v1/projects/{ref}/postgrest | Gets project&#x27;s postgrest config
*ServicesApi* | [**UpdatePostgRESTConfig**](docs/ServicesApi.md#updatepostgrestconfig) | **Patch** /v1/projects/{ref}/postgrest | Updates project&#x27;s postgrest config
*SnippetsApi* | [**GetSnippet**](docs/SnippetsApi.md#getsnippet) | **Get** /v1/snippets/{id} | Gets a specific SQL snippet
*SnippetsApi* | [**ListSnippets**](docs/SnippetsApi.md#listsnippets) | **Get** /v1/snippets | Lists SQL snippets for the logged in user
*SslEnforcementBetaApi* | [**GetSslEnforcementConfig**](docs/SslEnforcementBetaApi.md#getsslenforcementconfig) | **Get** /v1/projects/{ref}/ssl-enforcement | Get project&#x27;s SSL enforcement configuration.
*SslEnforcementBetaApi* | [**UpdateSslEnforcementConfig**](docs/SslEnforcementBetaApi.md#updatesslenforcementconfig) | **Put** /v1/projects/{ref}/ssl-enforcement | Update project&#x27;s SSL enforcement configuration.
*SsoApi* | [**CreateProviderForProject**](docs/SsoApi.md#createproviderforproject) | **Post** /v1/projects/{ref}/config/auth/sso/providers | Creates a new SSO provider
*SsoApi* | [**GetProviderById**](docs/SsoApi.md#getproviderbyid) | **Get** /v1/projects/{ref}/config/auth/sso/providers/{provider_id} | Gets a SSO provider by its UUID
*SsoApi* | [**ListAllProviders**](docs/SsoApi.md#listallproviders) | **Get** /v1/projects/{ref}/config/auth/sso/providers | Lists all SSO providers
*SsoApi* | [**RemoveProviderById**](docs/SsoApi.md#removeproviderbyid) | **Delete** /v1/projects/{ref}/config/auth/sso/providers/{provider_id} | Removes a SSO provider by its UUID
*SsoApi* | [**UpdateProviderById**](docs/SsoApi.md#updateproviderbyid) | **Put** /v1/projects/{ref}/config/auth/sso/providers/{provider_id} | Updates a SSO provider by its UUID
*StorageApi* | [**GetBuckets**](docs/StorageApi.md#getbuckets) | **Get** /v1/projects/{ref}/storage/buckets | Lists all buckets
*VanitySubdomainBetaApi* | [**ActivateVanitySubdomainPlease**](docs/VanitySubdomainBetaApi.md#activatevanitysubdomainplease) | **Post** /v1/projects/{ref}/vanity-subdomain/activate | Activates a vanity subdomain for a project.
*VanitySubdomainBetaApi* | [**CheckVanitySubdomainAvailability**](docs/VanitySubdomainBetaApi.md#checkvanitysubdomainavailability) | **Post** /v1/projects/{ref}/vanity-subdomain/check-availability | Checks vanity subdomain availability
*VanitySubdomainBetaApi* | [**GetVanitySubdomainConfig**](docs/VanitySubdomainBetaApi.md#getvanitysubdomainconfig) | **Get** /v1/projects/{ref}/vanity-subdomain | Gets current vanity subdomain config
*VanitySubdomainBetaApi* | [**RemoveVanitySubdomainConfig**](docs/VanitySubdomainBetaApi.md#removevanitysubdomainconfig) | **Delete** /v1/projects/{ref}/vanity-subdomain | Deletes a project&#x27;s vanity subdomain configuration

## Documentation For Models

 - [ActivateVanitySubdomainResponse](docs/ActivateVanitySubdomainResponse.md)
 - [AllOfDatabaseUpgradeStatusResponseDatabaseUpgradeStatus](docs/AllOfDatabaseUpgradeStatusResponseDatabaseUpgradeStatus.md)
 - [ApiKeyResponse](docs/ApiKeyResponse.md)
 - [AttributeMapping](docs/AttributeMapping.md)
 - [AttributeValue](docs/AttributeValue.md)
 - [AuthConfigResponse](docs/AuthConfigResponse.md)
 - [AuthHealthResponse](docs/AuthHealthResponse.md)
 - [BillingPlanId](docs/BillingPlanId.md)
 - [BranchDetailResponse](docs/BranchDetailResponse.md)
 - [BranchResponse](docs/BranchResponse.md)
 - [CreateBranchBody](docs/CreateBranchBody.md)
 - [CreateFunctionBody](docs/CreateFunctionBody.md)
 - [CreateOrganizationBodyV1](docs/CreateOrganizationBodyV1.md)
 - [CreateProjectBody](docs/CreateProjectBody.md)
 - [CreateProviderBody](docs/CreateProviderBody.md)
 - [CreateProviderResponse](docs/CreateProviderResponse.md)
 - [CreateSecretBody](docs/CreateSecretBody.md)
 - [DatabaseResponse](docs/DatabaseResponse.md)
 - [DatabaseUpgradeStatus](docs/DatabaseUpgradeStatus.md)
 - [DatabaseUpgradeStatusResponse](docs/DatabaseUpgradeStatusResponse.md)
 - [DeleteProviderResponse](docs/DeleteProviderResponse.md)
 - [Domain](docs/Domain.md)
 - [FunctionResponse](docs/FunctionResponse.md)
 - [FunctionSlugResponse](docs/FunctionSlugResponse.md)
 - [GetProviderResponse](docs/GetProviderResponse.md)
 - [ListProvidersResponse](docs/ListProvidersResponse.md)
 - [NetworkBanResponse](docs/NetworkBanResponse.md)
 - [NetworkRestrictionsRequest](docs/NetworkRestrictionsRequest.md)
 - [NetworkRestrictionsResponse](docs/NetworkRestrictionsResponse.md)
 - [OAuthTokenBody](docs/OAuthTokenBody.md)
 - [OAuthTokenResponse](docs/OAuthTokenResponse.md)
 - [OneOfAttributeValueDefault_](docs/OneOfAttributeValueDefault_.md)
 - [OneOfServiceHealthResponseInfo](docs/OneOfServiceHealthResponseInfo.md)
 - [OrganizationResponseV1](docs/OrganizationResponseV1.md)
 - [PgsodiumConfigResponse](docs/PgsodiumConfigResponse.md)
 - [PostgresConfigResponse](docs/PostgresConfigResponse.md)
 - [PostgrestConfigResponse](docs/PostgrestConfigResponse.md)
 - [PostgrestConfigWithJwtSecretResponse](docs/PostgrestConfigWithJwtSecretResponse.md)
 - [ProjectRefResponse](docs/ProjectRefResponse.md)
 - [ProjectResponse](docs/ProjectResponse.md)
 - [ProjectUpgradeEligibilityResponse](docs/ProjectUpgradeEligibilityResponse.md)
 - [ProjectUpgradeInitiateResponse](docs/ProjectUpgradeInitiateResponse.md)
 - [ProjectVersion](docs/ProjectVersion.md)
 - [Provider](docs/Provider.md)
 - [ReadOnlyStatusResponse](docs/ReadOnlyStatusResponse.md)
 - [RealtimeHealthResponse](docs/RealtimeHealthResponse.md)
 - [RemoveNetworkBanRequest](docs/RemoveNetworkBanRequest.md)
 - [RemoveReadReplicaBody](docs/RemoveReadReplicaBody.md)
 - [RunQueryBody](docs/RunQueryBody.md)
 - [SamlDescriptor](docs/SamlDescriptor.md)
 - [SecretResponse](docs/SecretResponse.md)
 - [ServiceHealthResponse](docs/ServiceHealthResponse.md)
 - [SetUpReadReplicaBody](docs/SetUpReadReplicaBody.md)
 - [SnippetContent](docs/SnippetContent.md)
 - [SnippetList](docs/SnippetList.md)
 - [SnippetMeta](docs/SnippetMeta.md)
 - [SnippetProject](docs/SnippetProject.md)
 - [SnippetResponse](docs/SnippetResponse.md)
 - [SnippetUser](docs/SnippetUser.md)
 - [SslEnforcementRequest](docs/SslEnforcementRequest.md)
 - [SslEnforcementResponse](docs/SslEnforcementResponse.md)
 - [SslEnforcements](docs/SslEnforcements.md)
 - [SubdomainAvailabilityResponse](docs/SubdomainAvailabilityResponse.md)
 - [TypescriptResponse](docs/TypescriptResponse.md)
 - [UpdateAuthConfigBody](docs/UpdateAuthConfigBody.md)
 - [UpdateBranchBody](docs/UpdateBranchBody.md)
 - [UpdateCustomHostnameBody](docs/UpdateCustomHostnameBody.md)
 - [UpdateCustomHostnameResponse](docs/UpdateCustomHostnameResponse.md)
 - [UpdateFunctionBody](docs/UpdateFunctionBody.md)
 - [UpdatePgsodiumConfigBody](docs/UpdatePgsodiumConfigBody.md)
 - [UpdatePostgresConfigBody](docs/UpdatePostgresConfigBody.md)
 - [UpdatePostgrestConfigBody](docs/UpdatePostgrestConfigBody.md)
 - [UpdateProviderBody](docs/UpdateProviderBody.md)
 - [UpdateProviderResponse](docs/UpdateProviderResponse.md)
 - [UpgradeDatabaseBody](docs/UpgradeDatabaseBody.md)
 - [V1Backup](docs/V1Backup.md)
 - [V1BackupsResponse](docs/V1BackupsResponse.md)
 - [V1BackupsResponsePhysicalBackupData](docs/V1BackupsResponsePhysicalBackupData.md)
 - [V1OrganizationMemberResponse](docs/V1OrganizationMemberResponse.md)
 - [V1OrganizationSlugResponse](docs/V1OrganizationSlugResponse.md)
 - [V1PgbouncerConfigResponse](docs/V1PgbouncerConfigResponse.md)
 - [V1RestorePitrBody](docs/V1RestorePitrBody.md)
 - [V1StorageBucketResponse](docs/V1StorageBucketResponse.md)
 - [VanitySubdomainBody](docs/VanitySubdomainBody.md)
 - [VanitySubdomainConfigResponse](docs/VanitySubdomainConfigResponse.md)

## Documentation For Authorization

## bearer

## Author


