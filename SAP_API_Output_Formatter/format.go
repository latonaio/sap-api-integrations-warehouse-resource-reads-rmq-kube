package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-warehouse-resource-reads-rmq-kube/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToHeader(raw []byte, l *logger.Logger) (*Header, error) {
	pm := &responses.Header{}
	err := json.Unmarshal(raw, &pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Header. unmarshal error: %w", err)
	}

	return &Header{
		EWMWarehouse:                  pm.EWMWarehouse,
		EWMResource:                   pm.EWMResource,
		UserName:                      pm.UserName,
		EWMResourceGroup:              pm.EWMResourceGroup,
		EWMResourceType:               pm.EWMResourceType,
		WarehouseOrderQueue:           pm.WarehouseOrderQueue,
		EWMCurrentQueue:               pm.EWMCurrentQueue,
		EWMStorTypeOfLastWhseTaskConf: pm.EWMStorTypeOfLastWhseTaskConf,
		EWMResourceLogonDateTime:      pm.EWMResourceLogonDateTime,
		EWMRsceIsLoggedOntoByCurUser:  pm.EWMRsceIsLoggedOntoByCurUser,
	}, nil
}
