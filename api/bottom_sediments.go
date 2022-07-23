package api

import (
	"database/sql"
	"net/http"

	db "github.com/Feruz666/eco-monitoring/db/sqlc"
	"github.com/gin-gonic/gin"
)

type createBottomSedimentsRequest struct {
	SubstanseName   string  `json:"substanse_name"`
	DateOfSampling  string  `json:"date_of_sampling"`
	NumberOfSample  string  `json:"number_of_sample"`
	TypeOfSediments string  `json:"type_of_sediments"`
	Concentration   float32 `json:"concentration"`
	UnitOfMeasure   string  `json:"unit_of_measure"`
	Location        string  `json:"location"`
	Longitude       string  `json:"longitude"`
	Latitude        string  `json:"latitude"`
	Reservoir       string  `json:"reservoir"`
	LicenseArea     string  `json:"license_area"`
	NumOfLicense    string  `json:"num_of_license"`
	Company         string  `json:"company"`
	MethodOfDeterm  string  `json:"method_of_determ"`
	Laboratory      string  `json:"laboratory"`
}

func (server *Server) createBottomSediments(ctx *gin.Context) {
	var req createBottomSedimentsRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateBottomSedimentsParams{
		SubstanseName:   req.SubstanseName,
		DateOfSampling:  req.DateOfSampling,
		NumberOfSample:  req.NumberOfSample,
		TypeOfSediments: req.TypeOfSediments,
		Concentration:   req.Concentration,
		UnitOfMeasure:   req.UnitOfMeasure,
		Location:        req.Location,
		Longitude:       req.Longitude,
		Latitude:        req.Latitude,
		Reservoir:       req.Reservoir,
		LicenseArea:     req.LicenseArea,
		NumOfLicense:    req.NumOfLicense,
		Company:         req.Company,
		MethodOfDeterm:  req.MethodOfDeterm,
		Laboratory:      req.Laboratory,
	}

	bottomSediments, err := server.store.CreateBottomSediments(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	ctx.JSON(http.StatusOK, bottomSediments)
}

type getBottomSedimentsRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getBottomSediments(ctx *gin.Context) {
	var req getBottomSedimentsRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	bottomSediments, err := server.store.GetBottomSediments(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, bottomSediments)
}

type listBottomSedimentsRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listBottomSediments(ctx *gin.Context) {
	var req listBottomSedimentsRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	arg := db.ListBottomSedimentsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	bottomSedimentss, err := server.store.ListBottomSediments(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, bottomSedimentss)
}

type deleteBottomSedimentsRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) deleteBottomSediments(ctx *gin.Context) {
	var req deleteBottomSedimentsRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	err := server.store.DeleteBottomSediments(ctx, req.ID)
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

type updateBottomSedimentsRequest struct {
	ID            int64  `json:"id"`
	SubstanseName string `json:"substanse_name"`
}

func (server *Server) updateBottomSediments(ctx *gin.Context) {
	var req updateBottomSedimentsRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	arg := db.UpdateBottomSedimentsParams{
		ID:            req.ID,
		SubstanseName: req.SubstanseName,
	}

	err := server.store.UpdateBottomSediments(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, err)
}
