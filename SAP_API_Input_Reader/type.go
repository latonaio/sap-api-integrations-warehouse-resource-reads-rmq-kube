package sap_api_input_reader

type EC_MC struct {
	ConnectionKey string      `json:"connection_key"`
	Result        bool        `json:"result"`
	RedisKey      string      `json:"redis_key"`
	Filepath      string      `json:"filepath"`
	BillingDocument   struct {
		BillingDocument                string      `json:"document_no"`
		DeliverTo                      string      `json:"deliver_to"`
		Quantity                       string      `json:"quantity"`
		PickedQuantity                 string      `json:"picked_quantity"`
		TotalNetAmount                 string      `json:"price"`
	    Batch                          string      `json:"batch"`
	} `json:"document"`
	ProductionOrder struct {
		DocumentNo           string      `json:"document_no"`
		Status               string      `json:"status"`
		DeliverTo            string      `json:"deliver_to"`
		Quantity             string      `json:"quantity"`
		CompletedQuantity    string      `json:"completed_quantity"`
	    PlannedStartDate     string      `json:"planned_start_date"`
	    PlannedValidatedDate string      `json:"planned_validated_date"`
	    ActualStartDate      string      `json:"actual_start_date"`
	    ActualValidatedDate  string      `json:"actual_validated_date"`
	    Batch                string      `json:"batch"`
		Work              struct {
			WorkNo                   string      `json:"work_no"`
			Quantity                 string      `json:"quantity"`
			CompletedQuantity        string      `json:"completed_quantity"`
			ErroredQuantity          string      `json:"errored_quantity"`
			Component                string      `json:"component"`
			PlannedComponentQuantity string      `json:"planned_component_quantity"`
			PlannedStartDate         string      `json:"planned_start_date"`
			PlannedStartTime         string      `json:"planned_start_time"`
			PlannedValidatedDate     string      `json:"planned_validated_date"`
			PlannedValidatedTime     string      `json:"planned_validated_time"`
			ActualStartDate          string      `json:"actual_start_date"`
			ActualStartTime          string      `json:"actual_start_time"`
			ActualValidatedDate      string      `json:"actual_validated_date"`
			ActualValidatedTime      string      `json:"actual_validated_time"`
		} `json:"work"`
	} `json:"production_order"`
	APISchema               string      `json:"api_schema"`
	MaterialCode            string      `json:"material_code"`
	Plant_Supplier          string      `json:"plant/supplier"`
	Stock                   string      `json:"stock"`
	BillingDocumentType     string      `json:"document_type"`
	SDDocument              string      `json:"document_no"`
	PlannedDate             string      `json:"planned_date"`
	BillingDocumentDate     string      `json:"validated_date"`
	Deleted                 bool        `json:"deleted"`
}


type SDC struct {
	ConnectionKey     string `json:"connection_key"`
	Result            bool   `json:"result"`
	RedisKey          string `json:"redis_key"`
	Filepath          string `json:"filepath"`
	WarehouseResource struct {
		EWMWarehouse                  string      `json:"EWMWarehouse"`
		EWMResource                   string      `json:"EWMResource"`
		UserName                      string      `json:"UserName"`
		EWMResourceGroup              string      `json:"EWMResourceGroup"`
		EWMResourceType               string      `json:"EWMResourceType"`
		WarehouseOrderQueue           string      `json:"WarehouseOrderQueue"`
		EWMCurrentQueue               string      `json:"EWMCurrentQueue"`
		EWMStorTypeOfLastWhseTaskConf string      `json:"EWMStorTypeOfLastWhseTaskConf"`
		EWMResourceLogonDateTime      int         `json:"EWMResourceLogonDateTime"`
		EWMRsceIsLoggedOntoByCurUser  bool        `json:"EWMRsceIsLoggedOntoByCurUser"`
	} `json:"WarehouseResource"`
	APISchema       string   `json:"api_schema"`
	Accepter        []string `json:"accepter"`
	BillingDocument string   `json:"billing_document"`
	Deleted         bool     `json:"deleted"`
}
