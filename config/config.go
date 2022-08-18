package config

import (
	"fmt"

	"github.com/spf13/viper"
)

const (
	ShowAllConstraintsKey          = "showAllConstraints"
	UseAllTablesKey                = "useAllTables"
	SelectedTablesKey              = "selectedTables"
	SchemaKey                      = "schema"
	ConnectionStringKey            = "connectionString"
	DB_USER												 = "db_user"
	DB_PASSWORD										 = "db_password"
	DB_PORT												 = "db_port"
	DB_NAME												 = "db_name"
	environmentVariables           = "environmentVariables"
	ConnectionStringSuggestionsKey = "connectionStringSuggestions"
	OutputFileNameKey              = "outputFileName"
	EncloseWithMermaidBackticksKey = "encloseWithMermaidBackticks"
	DebugKey                       = "debug"
	OmitConstraintLabelsKey        = "omitConstraintLabels"
)

type config struct{}

type MermerdConfig interface {
	ShowAllConstraints() bool
	UseAllTables() bool
	Schema() string
	ConnectionString() string
	OutputFileName() string
	ConnectionStringSuggestions() []string
	SelectedTables() []string
	EncloseWithMermaidBackticks() bool
	Debug() bool
	OmitConstraintLabels() bool
}

func NewConfig() MermerdConfig {
	return config{}
}

func (c config) ShowAllConstraints() bool {
	return viper.GetBool(ShowAllConstraintsKey)
}

func (c config) UseAllTables() bool {
	return viper.GetBool(UseAllTablesKey)
}

func (c config) Schema() string {
	return viper.GetString(SchemaKey)
}

func (c config) ConnectionString() string {
	var environment = viper.GetBool(environmentVariables)
	var connectionString = ""

	if environment {

		var port = viper.Get(DB_PORT)
		var portString = fmt.Sprintf("%v", port)

		connectionString = "postgresql://" + viper.Get(DB_USER).(string) + ":" + viper.Get(DB_PASSWORD).(string) + "@host.docker.internal" + ":" + portString + "/" + viper.Get(DB_NAME).(string)
	} else {
		return viper.GetString(ConnectionStringKey)
	}
	return connectionString
}

func (c config) OutputFileName() string {
	return viper.GetString(OutputFileNameKey)
}

func (c config) ConnectionStringSuggestions() []string {
	return viper.GetStringSlice(ConnectionStringSuggestionsKey)
}

func (c config) SelectedTables() []string {
	return viper.GetStringSlice(SelectedTablesKey)
}

func (c config) EncloseWithMermaidBackticks() bool {
	return viper.GetBool(EncloseWithMermaidBackticksKey)
}

func (c config) Debug() bool {
	return viper.GetBool(DebugKey)
}

func (c config) OmitConstraintLabels() bool {
	return viper.GetBool(OmitConstraintLabelsKey)
}
