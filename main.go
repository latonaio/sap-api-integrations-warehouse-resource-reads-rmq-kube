package main

import (
	sap_api_caller "sap-api-integrations-warehouse-resource-reads-rmq-kube/SAP_API_Caller"
	sap_api_input_reader "sap-api-integrations-warehouse-resource-reads-rmq-kube/SAP_API_Input_Reader"
	"sap-api-integrations-warehouse-resource-reads-rmq-kube/config"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	rabbitmq "github.com/latonaio/rabbitmq-golang-client"
	"golang.org/x/xerrors"
)

func main() {
	l := logger.NewLogger()
	conf := config.NewConf()
	rmq, err := rabbitmq.NewRabbitmqClient(conf.RMQ.URL(), conf.RMQ.QueueFrom(), conf.RMQ.QueueTo())
	if err != nil {
		l.Fatal(err.Error())
	}
	defer rmq.Close()

	caller := sap_api_caller.NewSAPAPICaller(
		conf.SAP.BaseURL(),
		conf.RMQ.QueueTo(),
		rmq,
		l,
	)

	iter, err := rmq.Iterator()
	if err != nil {
		l.Fatal(err.Error())
	}
	defer rmq.Stop()

	for msg := range iter {
		err = callWarehouseResource(caller, msg)
		if err != nil {
			msg.Fail()
			l.Error(err)
			continue
		}
		msg.Success()
	}
}

func callWarehouseResource(caller *sap_api_caller.SAPAPICaller, msg rabbitmq.RabbitmqMessage) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = xerrors.Errorf("error occurred: %w", e)
			return
		}
	}()
	eWMWarehouse,eWMResource := extractData(msg.Data())
	accepter := getAccepter(msg.Data())
	caller.AsyncGetWarehouseResource(eWMWarehouse,eWMResource, accepter)
	return nil
}

func extractData(data map[string]interface{}) (eWMWarehouse, eWMResource string) {
	sdc := sap_api_input_reader.ConvertToSDC(data)
	eWMWarehouse = sdc.WarehouseResource.EWMWarehouse
	eWMResource = sdc.WarehouseResource.EWMResource
	return
}

func getAccepter(data map[string]interface{}) []string {
	sdc := sap_api_input_reader.ConvertToSDC(data)
	accepter := sdc.Accepter
	if len(sdc.Accepter) == 0 {
		accepter = []string{"All"}
	}

	if accepter[0] == "All" {
		accepter = []string{
			"Header", 
		}
	}
	return accepter
}
