/*
 * postgres-meta
 *
 * A REST API to manage your Postgres database
 *
 * API version: 0.0.0-automated
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package meta_api

type CommandEnum string

// List of commandEnum
const (
	INSERT_CommandEnum  CommandEnum = "INSERT"
	UPDATE_CommandEnum  CommandEnum = "UPDATE"
	DELETE_CommandEnum  CommandEnum = "DELETE"
	SELECT__CommandEnum CommandEnum = "SELECT"
	ALL_CommandEnum     CommandEnum = "ALL"
)
