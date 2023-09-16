package handlers

import (
	"net/http"

	"eos2git.cec.lab.emc.com/Gargi-Bandyopadhyay/vyasa-cveinformation-registrar/src/utilities"
	"github.com/gin-gonic/gin"
)

func GetApidoc(c *gin.Context) {
	resp := utilities.APIDoc{}
	resp.Link = utilities.DocURL
	c.JSON(http.StatusOK, resp)
}
