package api

import (
	db "github.com/Feruz666/eco-monitoring/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	// air
	router.POST("/air", server.createAir)
	router.GET("/air/:id", server.getAir)
	router.GET("/air", server.listAir)
	router.DELETE("/air/:id", server.deleteAir)
	router.PUT("/air", server.updateAir)

	// snow
	router.POST("/snow", server.createSnow)
	router.GET("/snow/:id", server.getSnow)
	router.GET("/snow", server.listSnow)
	router.DELETE("/snow/:id", server.deleteSnow)
	router.PUT("/snow", server.updateSnow)

	// bottom_sediments
	router.POST("/bot-sed", server.createBottomSediments)
	router.GET("/bot-sed/:id", server.getBottomSediments)
	router.GET("/bot-sed", server.listBottomSediments)
	router.DELETE("/bot-sed/:id", server.deleteBottomSediments)
	router.PUT("/bot-sed", server.updateBottomSediments)

	// soil
	router.POST("/soil", server.createSoil)
	router.GET("/soil/:id", server.getSoil)
	router.GET("/soil", server.listSoil)
	router.DELETE("/soil/:id", server.deleteSoil)
	router.PUT("/soil", server.updateSoil)

	// surface_water
	router.POST("/sur-wat", server.createSurfaceWater)
	router.GET("/sur-wat/:id", server.getSurfaceWater)
	router.GET("/sur-wat", server.listSurfaceWater)
	router.DELETE("/sur-wat/:id", server.deleteSurfaceWater)
	router.PUT("/sur-wat", server.updateSurfaceWater)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
