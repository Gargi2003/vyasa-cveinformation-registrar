package handlers

import (
	"net/http"

	"eos2git.cec.lab.emc.com/Gargi-Bandyopadhyay/vyasa-cveinformation-registrar/src/utilities"
	"github.com/gin-gonic/gin"
)

func GetAppInfo(c *gin.Context) {
	info := &utilities.CommonAppInformation{
		Id:          "12973",
		Name:        "CVEInformation",
		Identifier:  "cve-info",
		Description: "An app to retrieve CVE information",
		Version:     "2.0",
		BuildNumber: "54321",
	}

	c.JSON(http.StatusOK, info)
}
