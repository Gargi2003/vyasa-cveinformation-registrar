package handlers

import (
	"net/http"

	"eos2git.cec.lab.emc.com/Gargi-Bandyopadhyay/vyasa-cveinformation-registrar/src/utilities"
	"github.com/gin-gonic/gin"
)

func GetHealthStatus(c *gin.Context) {
	status := utilities.CommonHealthCheckResponse{
		ServiceName: "CVEInformation",
		Status:      "Up",
	}

	c.JSON(http.StatusOK, status)
}
