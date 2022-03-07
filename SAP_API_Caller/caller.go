package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	sap_api_output_formatter "sap-api-integrations-warehouse-resource-reads-rmq-kube/SAP_API_Output_Formatter"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

type RMQOutputter interface {
	Send(sendQueue string, payload map[string]interface{}) error
}

type SAPAPICaller struct {
	baseURL      string
	apiKey       string
	outputQueues []string
	outputter    RMQOutputter
	log          *logger.Logger
}

func NewSAPAPICaller(baseUrl string, outputQueueTo []string, outputter RMQOutputter, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL:      baseUrl,
		apiKey:       GetApiKey(),
		outputQueues: outputQueueTo,
		outputter:    outputter,
		log:          l,
	}
}

func (c *SAPAPICaller) AsyncGetWarehouseResource(eWMWarehouse, eWMResource string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "Header":
			func() {
				c.Header(eWMWarehouse, eWMResource)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}
	wg.Wait()
}

func (c *SAPAPICaller) Header(eWMWarehouse, eWMResource string) {
	data, err := c.callWarehouseSrvAPIRequirementHeader("WarehouseResource", eWMWarehouse, eWMResource)
	if err != nil {
		c.log.Error(err)
		return
	}
	err = c.outputter.Send(c.outputQueues[0], map[string]interface{}{"message": data, "function": "WarehouseResourceHeader"})
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)

}

func (c *SAPAPICaller) callWarehouseSrvAPIRequirementHeader(api, eWMWarehouse, eWMResource string) (*sap_api_output_formatter.Header, error) {
	url := strings.Join([]string{c.baseURL, "api_warehouse_resource_2/srvd_a2x/sap/warehouseresource/0001", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithHeader(req, eWMWarehouse, eWMResource)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToHeader(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) setHeaderAPIKeyAccept(req *http.Request) {
	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")
}

func (c *SAPAPICaller) getQueryWithHeader(req *http.Request, eWMWarehouse, eWMResource string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("eWMWarehouse eq '%s' and eWMResource eq '%s'", eWMWarehouse, eWMResource))
	req.URL.RawQuery = params.Encode()
}
