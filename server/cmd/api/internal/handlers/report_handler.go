package handlers

import (
	"cmd/api/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ReportHandler struct {
	reportService *service.ReportService
}

func NewReportHandler(reportService *service.ReportService) *ReportHandler {
	return &ReportHandler{
		reportService: reportService,
	}
}

// GetIncomeStatement handles GET /reports/income-statement
func (h *ReportHandler) GetIncomeStatement(c *gin.Context) {
	fromDate := c.Query("from_date")
	toDate := c.Query("to_date")

	if fromDate == "" || toDate == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "from_date and to_date query parameters are required"})
		return
	}

	statement, err := h.reportService.GetIncomeStatement(fromDate, toDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, statement)
}
