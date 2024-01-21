
package roles

import (
	"github.com/sev-2/raiden/pkg/postgres"
)

var PgsodiumKeyholder = &postgres.Role{
	ActiveConnections : 0,
	CanBypassRLS : false,
	CanCreateDB : false,
	CanCreateRole : false,
	CanLogin : false,
	Config : map[string]any{},
	ConnectionLimit : 100,
	ID : 16733,
	InheritRole : true,
	IsReplicationRole : false,
	IsSuperuser : false,
	Name : "pgsodium_keyholder",
	ValidUntil : nil,
}
