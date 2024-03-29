/*
 * postgres-meta
 *
 * A REST API to manage your Postgres database
 *
 * API version: 0.0.0-automated
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package meta_api

type TablesRelationships struct {
	Id                int32  `json:"id"`
	ConstraintName    string `json:"constraint_name"`
	SourceSchema      string `json:"source_schema"`
	SourceTableName   string `json:"source_table_name"`
	SourceColumnName  string `json:"source_column_name"`
	TargetTableSchema string `json:"target_table_schema"`
	TargetTableName   string `json:"target_table_name"`
	TargetColumnName  string `json:"target_column_name"`
}
