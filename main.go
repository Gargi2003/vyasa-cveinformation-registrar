package main

import (
	"eos2git.cec.lab.emc.com/Gargi-Bandyopadhyay/vyasa-cveinformation-registrar/src/handlers"
	"eos2git.cec.lab.emc.com/Gargi-Bandyopadhyay/vyasa-cveinformation-registrar/src/utilities"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func main() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	router := gin.Default()
	router.GET(utilities.CVEEndpointById, handlers.GetCVEInfo)
	router.GET(utilities.CVEEndpoint, handlers.ListCVEInfo)
	router.POST(utilities.CVEEndpoint, handlers.CreateCVEInfo)
	router.PUT(utilities.CVEEndpointById, handlers.UpdateCVEInfo)
	router.DELETE(utilities.CVEEndpointById, handlers.DeleteCVEInfo)

	router.GET(utilities.HealthEndpoint, handlers.GetHealthStatus)
	router.GET(utilities.InfoEndpoint, handlers.GetAppInfo)

	router.GET(utilities.ApiDocEndpoint, handlers.GetApidoc)
	router.Run(":8080")
}
