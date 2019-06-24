package report

import (
	"fmt"
	"time"

	orders "github.com/reynaldipane/test-ijah/order"
	productstock "github.com/reynaldipane/test-ijah/product_stock"
)

/*ServiceInterface wraps the report service ability
so it can be easier to initialize the reportHandler with different kind of service
*/
type ServiceInterface interface {
	getProductValuesReport() ([]productStock, error)
	getSalesReport() ([]order, error)
	convertProductValuesReportToCsv() map[string]string
	convertSalesReportDetailToCsv() map[string]string
}

/*Service is report service object that has access to the
needed services to generate report */
type Service struct {
	repo                RepositoryInterface
	productStockService productstock.ServiceInterface
	orderService        orders.ServiceInterface
}

func (service *Service) getProductValuesReport() ([]productStock, error) {
	return service.repo.getProductValuesReport()
}

func (service *Service) getSalesReport() ([]order, error) {
	return service.repo.getSalesReport()
}

func (service *Service) convertProductValuesReportToCsv() map[string]string {
	productStocks, err := service.getProductValuesReport()

	if err != nil {
		fmt.Println(err)
	}

	productValueReport := convertToProductValuesReportStruct(productStocks)
	structHeaders := extractStructHeaders(&productValuesReport{})

	convertedStruct := [][]string{structHeaders}

	for _, v := range productValueReport {
		convertedStruct = append(convertedStruct, extractStructBody(structHeaders, v))
	}

	fileName := fmt.Sprintf("product_values%v", time.Now().UnixNano())
	convertToCsv(convertedStruct, fileName)

	return map[string]string{
		"download_link": fmt.Sprintf("http://localhost:9000/report/downloads/%v%v", fileName, ".csv"),
	}
}

func (service *Service) convertSalesReportDetailToCsv() map[string]string {
	orders, err := service.getSalesReport()

	if err != nil {
		fmt.Println(err)
	}

	salesReports := convertToSalesReportStruct(orders)
	structHeaders := extractStructHeaders(&salesDetailReport{})
	convertedStruct := [][]string{structHeaders}

	for _, v := range salesReports {
		convertedStruct = append(convertedStruct, extractStructBody(structHeaders, v))
	}

	fileName := fmt.Sprintf("sales_report%v", time.Now().UnixNano())
	convertToCsv(convertedStruct, fileName)

	return map[string]string{
		"download_link": fmt.Sprintf("http://localhost:9000/report/downloads/%v%v", fileName, ".csv"),
	}
}
