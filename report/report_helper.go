package report

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"
)

func convertToProductValuesReportStruct(productStocks []productStock) []productValuesReport {
	var productsValueReport []productValuesReport

	for _, v := range productStocks {
		productValue := productValuesReport{
			Sku:             v.Product.Sku,
			Name:            v.Product.Name,
			Stock:           v.Stock,
			AverageBuyPrice: v.AverageBuyPrice,
			TotalPrice:      v.AverageBuyPrice * float64(v.Stock),
		}

		productsValueReport = append(productsValueReport, productValue)
	}

	return productsValueReport
}

func convertToSalesReportStruct(orders []order) []salesDetailReport {
	var salesReports []salesDetailReport

	for _, v := range orders {
		salesReportUnit := salesDetailReport{
			OrderID:    v.ID,
			OrderTime:  v.CreatedAt,
			Sku:        v.Product.Sku,
			Name:       v.Product.Name,
			NumberSold: v.NumberSold,
			SellPrice:  v.SellPrice,
			TotalPrice: float64(v.TotalPrice),
			BuyPrice:   v.ProductStock.AverageBuyPrice,
			Profit:     float64(v.TotalPrice) - (v.ProductStock.AverageBuyPrice * float64(v.NumberSold)),
		}

		salesReports = append(salesReports, salesReportUnit)
	}

	return salesReports
}

func extractStructHeaders(structToExtractHeaders interface{}) []string {
	var headers []string

	val := reflect.ValueOf(structToExtractHeaders).Elem()
	for i := 0; i < val.NumField(); i++ {
		headers = append(headers, val.Type().Field(i).Name)
	}

	return headers
}

func extractStructBody(structHeaders []string, structToExtractBody interface{}) []string {
	var body []string

	var maps map[string]interface{}
	jsonStruct, _ := json.Marshal(structToExtractBody)
	json.Unmarshal(jsonStruct, &maps)

	for _, v := range structHeaders {
		observedKey := v
		for k, v := range maps {
			if k == observedKey {
				body = append(body, fmt.Sprintf("%v", v))
			}
		}
	}

	return body
}

func convertToCsv(data [][]string, fileName string) {
	file, err := os.Create("./csv_exports/" + fileName + ".csv")
	checkError("Cannot create file", err)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, value := range data {
		err := writer.Write(value)
		checkError("Cannot write to file", err)
	}
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}
