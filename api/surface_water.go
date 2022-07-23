package api

import (
	"database/sql"
	"net/http"

	db "github.com/Feruz666/eco-monitoring/db/sqlc"
	"github.com/gin-gonic/gin"
)

type createSurfaceWaterRequest struct {
	SubstanseName  string  `json:"substanse_name"`
	DateOfSampling string  `json:"date_of_sampling"`
	NumberOfSample string  `json:"number_of_sample"`
	Concentration  float32 `json:"concentration"`
	UnitOfMeasure  string  `json:"unit_of_measure"`
	Location       string  `json:"location"`
	Longitude      string  `json:"longitude"`
	Latitude       string  `json:"latitude"`
	Reservoir      string  `json:"reservoir"`
	LicenseArea    string  `json:"license_area"`
	NumOfLicense   string  `json:"num_of_license"`
	Company        string  `json:"company"`
	MethodOfDeterm string  `json:"method_of_determ"`
	Laboratory     string  `json:"laboratory"`
}

func (server *Server) createSurfaceWater(ctx *gin.Context) {
	var req createSurfaceWaterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateSurfaceWaterParams{
		SubstanseName:  req.SubstanseName,
		DateOfSampling: req.DateOfSampling,
		NumberOfSample: req.NumberOfSample,
		Concentration:  req.Concentration,
		UnitOfMeasure:  req.UnitOfMeasure,
		Location:       req.Location,
		Longitude:      req.Longitude,
		Latitude:       req.Latitude,
		LicenseArea:    req.LicenseArea,
		NumOfLicense:   req.NumOfLicense,
		Company:        req.Company,
		MethodOfDeterm: req.MethodOfDeterm,
		Laboratory:     req.Laboratory,
	}

	surfaceWater, err := server.store.CreateSurfaceWater(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	ctx.JSON(http.StatusOK, surfaceWater)
}

type getSurfaceWaterRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getSurfaceWater(ctx *gin.Context) {
	var req getSurfaceWaterRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	surfaceWater, err := server.store.GetSurfaceWater(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, surfaceWater)
}

type listSurfaceWaterRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listSurfaceWater(ctx *gin.Context) {
	var req listSurfaceWaterRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	arg := db.ListSurfaceWaterParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	surfaceWater, err := server.store.ListSurfaceWater(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, surfaceWater)
}

type deleteSurfaceWaterRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) deleteSurfaceWater(ctx *gin.Context) {
	var req deleteSurfaceWaterRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	err := server.store.DeleteSurfaceWater(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, err)
}

type updateSurfaceWaterRequest struct {
	ID            int64  `json:"id"`
	SubstanseName string `json:"substanse_name"`
}

func (server *Server) updateSurfaceWater(ctx *gin.Context) {
	var req updateSurfaceWaterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	arg := db.UpdateSurfaceWaterParams{
		ID:            req.ID,
		SubstanseName: req.SubstanseName,
	}

	err := server.store.UpdateSurfaceWater(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, err)
}
