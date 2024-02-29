package generator

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/sev-2/raiden/pkg/logger"
	"github.com/sev-2/raiden/pkg/utils"
)

// ----- Define type, variable and constant -----
type (
	GenerateRegisterRoleData struct {
		Imports []string
		Package string
		Roles   []string
	}
)

const (
	RoleRegisterFilename = "roles.go"
	RoleRegisterDir      = "internal/bootstrap"
	RoleRegisterTemplate = `// Code generated by raiden-cli; DO NOT EDIT.
package {{ .Package }}
{{if gt (len .Imports) 0 }}
import (
{{- range .Imports}}
	{{.}}
{{- end}}
)
{{end }}
func RegisterRoles() {
	resource.RegisterRole(
		{{- range .Roles}}
		&roles.{{.}}{},
		{{- end}}
	)
}
`
)

func GenerateRoleRegister(basePath string, projectName string, generateFn GenerateFn) error {
	roleRegisterDir := filepath.Join(basePath, RoleRegisterDir)
	logger.Debugf("GenerateRoleRegister - create %s folder if not exist", roleRegisterDir)
	if exist := utils.IsFolderExists(roleRegisterDir); !exist {
		if err := utils.CreateFolder(roleRegisterDir); err != nil {
			return err
		}
	}

	roleDir := filepath.Join(basePath, RoleDir)
	logger.Debugf("GenerateRoleRegister - create %s folder if not exist", roleDir)
	if exist := utils.IsFolderExists(roleDir); !exist {
		if err := utils.CreateFolder(roleDir); err != nil {
			return err
		}
	}

	// scan all controller
	roleList, err := WalkScanRole(roleDir)
	if err != nil {
		return err
	}

	input, err := createRoleRegisterInput(projectName, roleRegisterDir, roleList)
	if err != nil {
		return err
	}

	logger.Debugf("GenerateRoleRegister - generate role register to %s", input.OutputPath)
	return generateFn(input, nil)
}

func createRoleRegisterInput(projectName string, roleRegisterDir string, roleList []string) (input GenerateInput, err error) {
	// set file path
	filePath := filepath.Join(roleRegisterDir, RoleRegisterFilename)

	// set imports path
	imports := []string{
		fmt.Sprintf("%q", "github.com/sev-2/raiden/pkg/resource"),
	}

	if len(roleList) > 0 {
		rolesImportPath := fmt.Sprintf("%s/internal/roles", utils.ToGoModuleName(projectName))
		imports = append(imports, fmt.Sprintf("%q", rolesImportPath))
	}

	// set passed parameter
	data := GenerateRegisterRoleData{
		Package: "bootstrap",
		Imports: imports,
		Roles:   roleList,
	}

	input = GenerateInput{
		BindData:     data,
		Template:     RoleRegisterTemplate,
		TemplateName: "roleRegisterTemplate",
		OutputPath:   filePath,
	}

	return
}

func WalkScanRole(roleDir string) ([]string, error) {
	logger.Debugf("GenerateRoleRegister - scan %s for register all roles", roleDir)

	roles := make([]string, 0)
	err := filepath.Walk(roleDir, func(path string, info fs.FileInfo, err error) error {
		if strings.HasSuffix(path, ".go") {
			logger.Debugf("GenerateRoleRegister - collect roles from %s", path)
			rs, e := getStructByBaseName(path, "RoleBase")
			if e != nil {
				return e
			}

			roles = append(roles, rs...)

		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return roles, nil
}