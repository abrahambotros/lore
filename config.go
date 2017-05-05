package lore

/*
_config holds the current Config instance. Note that this should NEVER be accessed directly, and
should instead be retrieved via GetConfig to provide null safety.
*/
var _config *Config

/*
Config provides a struct for configuring all LORE queries.
*/
type Config struct {
	/*
		DB driver's placeholder format for injecting query parameters via SQL.
	*/
	SQLPlaceholderFormat SQLPlaceholderFormat
}

/*
GetConfig returns the current config object. If no config already exists, a default is given.
*/
func GetConfig() *Config {
	if _config == nil {
		_config = GetConfigDefault()
	}
	return _config
}

/*
GetConfigDefault returns the default config object.
*/
func GetConfigDefault() *Config {
	return &Config{
		SQLPlaceholderFormat: SQLPlaceholderFormatDollar,
	}
}

/*
SetConfig sets the current config for all future LORE queries.
*/
func SetConfig(c *Config) {
	_config = c
}

/*
SetConfigDefault sets a default config for all future LORE queries.
*/
func SetConfigDefault() {
	SetConfig(GetConfigDefault())
}
