package api

import (
	"database/sql"
	"net/http"

	db "github.com/Feruz666/eco-monitoring/db/sqlc"
	"github.com/gin-gonic/gin"
)

type createSnowRequest struct {
	SubstanseName    string  `json:"substanse_name"`
	DateOfSampling   string  `json:"date_of_sampling"`
	NumberOfSample   string  `json:"number_of_sample"`
	Concentration    float32 `json:"concentration"`
	UnitOfMeasure    string  `json:"unit_of_measure"`
	Location         string  `json:"location"`
	Longitude        string  `json:"longitude"`
	Latitude         string  `json:"latitude"`
	SourceOfEmission string  `json:"source_of_emission"`
	LicenseArea      string  `json:"license_area"`
	NumOfLicense     string  `json:"num_of_license"`
	Company          string  `json:"company"`
	MethodOfDeterm   string  `json:"method_of_determ"`
	Laboratory       string  `json:"laboratory"`
}

func (server *Server) createSnow(ctx *gin.Context) {
	var req createSnowRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateSnowParams{
		SubstanseName:    req.SubstanseName,
		DateOfSampling:   req.DateOfSampling,
		NumberOfSample:   req.NumberOfSample,
		Concentration:    req.Concentration,
		UnitOfMeasure:    req.UnitOfMeasure,
		Location:         req.Location,
		Longitude:        req.Longitude,
		Latitude:         req.Latitude,
		SourceOfEmission: req.SourceOfEmission,
		LicenseArea:      req.LicenseArea,
		NumOfLicense:     req.NumOfLicense,
		Company:          req.Company,
		MethodOfDeterm:   req.MethodOfDeterm,
		Laboratory:       req.Laboratory,
	}

	snow, err := server.store.CreateSnow(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	ctx.JSON(http.StatusOK, snow)
}

type getSnowRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getSnow(ctx *gin.Context) {
	var req getSnowRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	snow, err := server.store.GetSnow(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, snow)
}

type listSnowRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listSnow(ctx *gin.Context) {
	var req listSnowRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	arg := db.ListSnowParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	snows, err := server.store.ListSnow(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, snows)
}

type deleteSnowRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) deleteSnow(ctx *gin.Context) {
	var req deleteSnowRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	err := server.store.DeleteSnow(ctx, req.ID)
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

type updateSnowRequest struct {
	ID            int64  `json:"id"`
	SubstanseName string `json:"substanse_name"`
}

func (server *Server) updateSnow(ctx *gin.Context) {
	var req updateSnowRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	arg := db.UpdateSnowParams{
		ID:            req.ID,
		SubstanseName: req.SubstanseName,
	}

	err := server.store.UpdateSnow(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, err)
}
