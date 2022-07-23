package api

import (
	"database/sql"
	"net/http"

	db "github.com/Feruz666/eco-monitoring/db/sqlc"
	"github.com/gin-gonic/gin"
)

type createAirRequest struct {
	SubstanseName  string  `json:"substanse_name"`
	DateOfSampling string  `json:"date_of_sampling"`
	NumberOfSample string  `json:"number_of_sample"`
	Concentration  float32 `json:"concentration"`
	UnitOfMeasure  string  `json:"unit_of_measure"`
	Location       string  `json:"location"`
	Longitude      string  `json:"longitude"`
	Latitude       string  `json:"latitude"`
	LicenseArea    string  `json:"license_area"`
	NumOfLicense   string  `json:"num_of_license"`
	Company        string  `json:"company"`
	MethodOfDeterm string  `json:"method_of_determ"`
	Laboratory     string  `json:"laboratory"`
}

func (server *Server) createAir(ctx *gin.Context) {
	var req createAirRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateAirParams{
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

	air, err := server.store.CreateAir(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	ctx.JSON(http.StatusOK, air)
}

type getAirRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getAir(ctx *gin.Context) {
	var req getAirRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	air, err := server.store.GetAir(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, air)
}

type listAirRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listAir(ctx *gin.Context) {
	var req listAirRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	arg := db.ListAirParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	airs, err := server.store.ListAir(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, airs)
}

type deleteAirRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) deleteAir(ctx *gin.Context) {
	var req deleteAirRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	err := server.store.DeleteAir(ctx, req.ID)
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

type updateAirRequest struct {
	ID            int64  `json:"id"`
	SubstanseName string `json:"substanse_name"`
}

func (server *Server) updateAir(ctx *gin.Context) {
	var req updateAirRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	arg := db.UpdateAirParams{
		ID:            req.ID,
		SubstanseName: req.SubstanseName,
	}

	err := server.store.UpdateAir(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, err)
}
