/*
 * postgres-meta
 *
 * A REST API to manage your Postgres database
 *
 * API version: 0.0.0-automated
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package meta_api

type Policy struct {
	Id         int32        `json:"id,omitempty"`
	Schema     string       `json:"schema,omitempty"`
	Table      string       `json:"table"`
	TableId    int32        `json:"table_id,omitempty"`
	Name       string       `json:"name"`
	Action     string       `json:"action"`
	Roles      []string     `json:"roles"`
	Command    *CommandEnum `json:"command"`
	Definition string       `json:"definition,omitempty"`
	Check      string       `json:"check,omitempty"`
}
