package sap_api_output_formatter

type WarehouseResource struct {
	ConnectionKey     string `json:"connection_key"`
	Result            bool   `json:"result"`
	RedisKey          string `json:"redis_key"`
	Filepath          string `json:"filepath"`
	APISchema         string `json:"api_schema"`
	WarehouseResource string `json:"warehouse_resource"`
	Deleted           bool   `json:"deleted"`
}

type Header struct {
		EWMWarehouse                  string `json:"EWMWarehouse"`
		EWMResource                   string `json:"EWMResource"`
		UserName                      string `json:"UserName"`
		EWMResourceGroup              string `json:"EWMResourceGroup"`
		EWMResourceType               string `json:"EWMResourceType"`
		WarehouseOrderQueue           string `json:"WarehouseOrderQueue"`
		EWMCurrentQueue               string `json:"EWMCurrentQueue"`
		EWMStorTypeOfLastWhseTaskConf string `json:"EWMStorTypeOfLastWhseTaskConf"`
		EWMResourceLogonDateTime      int    `json:"EWMResourceLogonDateTime"`
		EWMRsceIsLoggedOntoByCurUser  bool   `json:"EWMRsceIsLoggedOntoByCurUser"`
	}
