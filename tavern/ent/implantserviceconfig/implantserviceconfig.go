// Code generated by entc, DO NOT EDIT.

package implantserviceconfig

const (
	// Label holds the string label denoting the implantserviceconfig type in the database.
	Label = "implant_service_config"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldExecutablePath holds the string denoting the executablepath field in the database.
	FieldExecutablePath = "executable_path"
	// EdgeImplantConfigs holds the string denoting the implantconfigs edge name in mutations.
	EdgeImplantConfigs = "implantConfigs"
	// Table holds the table name of the implantserviceconfig in the database.
	Table = "implant_service_configs"
	// ImplantConfigsTable is the table that holds the implantConfigs relation/edge. The primary key declared below.
	ImplantConfigsTable = "implant_config_serviceConfigs"
	// ImplantConfigsInverseTable is the table name for the ImplantConfig entity.
	// It exists in this package in order to avoid circular dependency with the "implantconfig" package.
	ImplantConfigsInverseTable = "implant_configs"
)

// Columns holds all SQL columns for implantserviceconfig fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldDescription,
	FieldExecutablePath,
}

var (
	// ImplantConfigsPrimaryKey and ImplantConfigsColumn2 are the table columns denoting the
	// primary key for the implantConfigs relation (M2M).
	ImplantConfigsPrimaryKey = []string{"implant_config_id", "implant_service_config_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DefaultDescription holds the default value on creation for the "description" field.
	DefaultDescription string
	// ExecutablePathValidator is a validator for the "executablePath" field. It is called by the builders before save.
	ExecutablePathValidator func(string) error
)
