package responses

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
