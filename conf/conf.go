package conf

/*
router重复了以最后一个出现的为准只看key
*/

// global config
var Gconf map[string]interface{}

// store satic file path
var Static map[string]interface{}

// store proxy router
var Router map[string]interface{}

// first read the config file
func init() {
	FlagInit()
	Static = Gconf["static"].(map[string]interface{})
	Router = Gconf["router"].(map[string]interface{})
}
