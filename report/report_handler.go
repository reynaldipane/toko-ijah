package report

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/reynaldipane/test-ijah/appcontext"
	helper "github.com/reynaldipane/test-ijah/helpers"
	orders "github.com/reynaldipane/test-ijah/order"
	productstock "github.com/reynaldipane/test-ijah/product_stock"
)

var reportService ServiceInterface

func getService() ServiceInterface {
	if reportService == nil {
		return &Service{
			repo: initRepository(appcontext.GetDB()),
			productStockService: &productstock.Service{
				Repo: productstock.InitRepository(appcontext.GetDB()),
			},
			orderService: &orders.Service{
				Repo: orders.InitRepository(appcontext.GetDB()),
			},
		}
	}

	return reportService
}

// GenerateProductValuesReport will return product values report
func GenerateProductValuesReport(w http.ResponseWriter, r *http.Request) {
	result, err := getService().getProductValuesReport()

	if err != nil {
		helper.BuildResponseWithError(w, helper.ContentJSON, 404, "report not found")
		return
	}

	response, err := json.Marshal(result)
	helper.BuildResponse(w, helper.ContentJSON, 200, string(response))
}

//ExportProductValuesReportCsv will return link to exported product values report to download
func ExportProductValuesReportCsv(w http.ResponseWriter, r *http.Request) {
	fileName := getService().convertProductValuesReportToCsv()
	response, _ := json.Marshal(fileName)
	helper.BuildResponse(w, helper.ContentJSON, 200, string(response))
}

// GenerateProductSalesDetailReport will return product sales detail report
func GenerateProductSalesDetailReport(w http.ResponseWriter, r *http.Request) {
	result, err := getService().getSalesReport()

	if err != nil {
		helper.BuildResponseWithError(w, helper.ContentJSON, 404, "report not found")
		return
	}

	response, err := json.Marshal(result)
	helper.BuildResponse(w, helper.ContentJSON, 200, string(response))
}

//ExportSalesReportCsv will return link to exported sales report to download
func ExportSalesReportCsv(w http.ResponseWriter, r *http.Request) {
	fileName := getService().convertSalesReportDetailToCsv()
	response, _ := json.Marshal(fileName)
	helper.BuildResponse(w, helper.ContentJSON, 200, string(response))
}

// DownloadProductValuesReport will return exported csv file to client
func DownloadProductValuesReport(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fileName := vars["fileName"]

	if fileName == "" {
		helper.BuildResponseWithError(w, helper.ContentJSON, 400, "report not found")
		return
	}

	filePath := "./csv_exports/" + fileName

	Openfile, err := os.Open(filePath)
	defer Openfile.Close()
	if err != nil {
		helper.BuildResponseWithError(w, helper.ContentJSON, 400, "report not found")
		return
	}

	FileHeader := make([]byte, 512)
	Openfile.Read(FileHeader)

	FileContentType := http.DetectContentType(FileHeader)

	FileStat, _ := Openfile.Stat()
	FileSize := strconv.FormatInt(FileStat.Size(), 10)

	w.Header().Set("Content-Disposition", "attachment; filename="+fileName)
	w.Header().Set("Content-Type", FileContentType)
	w.Header().Set("Content-Length", FileSize)

	Openfile.Seek(0, 0)
	io.Copy(w, Openfile)
	return
}
