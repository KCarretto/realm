// Code generated by entc, DO NOT EDIT.

package implantcallbackconfig

const (
	// Label holds the string label denoting the implantcallbackconfig type in the database.
	Label = "implant_callback_config"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldURI holds the string denoting the uri field in the database.
	FieldURI = "uri"
	// FieldProxyURI holds the string denoting the proxyuri field in the database.
	FieldProxyURI = "proxy_uri"
	// FieldPriority holds the string denoting the priority field in the database.
	FieldPriority = "priority"
	// FieldTimeout holds the string denoting the timeout field in the database.
	FieldTimeout = "timeout"
	// FieldInterval holds the string denoting the interval field in the database.
	FieldInterval = "interval"
	// FieldJitter holds the string denoting the jitter field in the database.
	FieldJitter = "jitter"
	// EdgeImplantConfigs holds the string denoting the implantconfigs edge name in mutations.
	EdgeImplantConfigs = "implantConfigs"
	// Table holds the table name of the implantcallbackconfig in the database.
	Table = "implant_callback_configs"
	// ImplantConfigsTable is the table that holds the implantConfigs relation/edge. The primary key declared below.
	ImplantConfigsTable = "implant_config_callbackConfigs"
	// ImplantConfigsInverseTable is the table name for the ImplantConfig entity.
	// It exists in this package in order to avoid circular dependency with the "implantconfig" package.
	ImplantConfigsInverseTable = "implant_configs"
)

// Columns holds all SQL columns for implantcallbackconfig fields.
var Columns = []string{
	FieldID,
	FieldURI,
	FieldProxyURI,
	FieldPriority,
	FieldTimeout,
	FieldInterval,
	FieldJitter,
}

var (
	// ImplantConfigsPrimaryKey and ImplantConfigsColumn2 are the table columns denoting the
	// primary key for the implantConfigs relation (M2M).
	ImplantConfigsPrimaryKey = []string{"implant_config_id", "implant_callback_config_id"}
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
	// URIValidator is a validator for the "uri" field. It is called by the builders before save.
	URIValidator func(string) error
	// DefaultPriority holds the default value on creation for the "priority" field.
	DefaultPriority int
	// PriorityValidator is a validator for the "priority" field. It is called by the builders before save.
	PriorityValidator func(int) error
	// DefaultTimeout holds the default value on creation for the "timeout" field.
	DefaultTimeout int
	// TimeoutValidator is a validator for the "timeout" field. It is called by the builders before save.
	TimeoutValidator func(int) error
	// DefaultInterval holds the default value on creation for the "interval" field.
	DefaultInterval int
	// IntervalValidator is a validator for the "interval" field. It is called by the builders before save.
	IntervalValidator func(int) error
	// DefaultJitter holds the default value on creation for the "jitter" field.
	DefaultJitter int
	// JitterValidator is a validator for the "jitter" field. It is called by the builders before save.
	JitterValidator func(int) error
)
