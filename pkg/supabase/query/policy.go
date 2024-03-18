package query

import (
	"fmt"
	"strings"

	"github.com/lib/pq"
	"github.com/sev-2/raiden/pkg/supabase/objects"
)

func BuildCreatePolicyQuery(policy objects.Policy) string {
	name := fmt.Sprintf("%q", strings.ToLower(policy.Name))

	definitionClause := ""
	if policy.Definition != "" {
		definitionClause = "USING (" + policy.Definition + ")"
	}
	checkClause := ""
	if policy.Check != nil && *policy.Check != "" {
		checkClause = "WITH CHECK (" + *policy.Check + ")"
	}

	roleList := ""
	for i, role := range policy.Roles {
		if i > 0 {
			roleList += ", "
		}
		roleList += pq.QuoteIdentifier(role)
	}

	return fmt.Sprintf(`
	CREATE POLICY %s ON %s.%s
	AS %s
	FOR %s
	TO %s
	%s %s;`, name, policy.Schema, policy.Table, policy.Action, policy.Command, roleList, definitionClause, checkClause)
}

func BuildUpdatePolicyQuery(policy objects.Policy, updatePolicyParams objects.UpdatePolicyParam) string {
	alter := fmt.Sprintf("ALTER POLICY %q ON %s.%s", updatePolicyParams.Name, policy.Schema, policy.Table)

	var nameSql, definitionSql, checkSql, rolesSql string
	for _, ut := range updatePolicyParams.ChangeItems {
		switch ut {
		case objects.UpdatePolicyName:
			if policy.Name != "" {
				nameSql = fmt.Sprintf("%s RENAME TO %s;", alter, policy.Name)
			}
		case objects.UpdatePolicyCheck:
			if policy.Check == nil || (policy.Check != nil && *policy.Check != "") {
				defaultCheck := "true"
				policy.Check = &defaultCheck
			}
			checkSql = fmt.Sprintf("%s WITH CHECK (%s);", alter, *policy.Check)
		case objects.UpdatePolicyDefinition:
			if policy.Definition != "" {
				definitionSql = fmt.Sprintf("%s USING (%s);", alter, policy.Definition)
			}
		case objects.UpdatePolicyRoles:
			if len(policy.Roles) > 0 {
				rolesSql = fmt.Sprintf("%s TO %s;", alter, strings.Join(policy.Roles, ","))
			}
		}
	}

	return fmt.Sprintf("BEGIN; %s %s %s %s COMMIT;", definitionSql, checkSql, rolesSql, nameSql)
}

func BuildDeletePolicyQuery(policy objects.Policy) string {
	return fmt.Sprintf("DROP POLICY %q ON %s.%s;", policy.Name, policy.Schema, policy.Table)
}