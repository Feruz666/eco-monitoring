package api

import (
	"database/sql"
	"net/http"

	db "github.com/Feruz666/eco-monitoring/db/sqlc"
	"github.com/gin-gonic/gin"
)

type createSoilRequest struct {
	SubstanseName  string  `json:"substanse_name"`
	DateOfSampling string  `json:"date_of_sampling"`
	NumberOfSample string  `json:"number_of_sample"`
	SoilSubtype    string  `json:"soil_subtype"`
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

func (server *Server) createSoil(ctx *gin.Context) {
	var req createSoilRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateSoilParams{
		SubstanseName:  req.SubstanseName,
		DateOfSampling: req.DateOfSampling,
		NumberOfSample: req.NumberOfSample,
		SoilSubtype:    req.SoilSubtype,
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

	soil, err := server.store.CreateSoil(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	ctx.JSON(http.StatusOK, soil)
}

type getSoilRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getSoil(ctx *gin.Context) {
	var req getSoilRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	soil, err := server.store.GetSoil(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, soil)
}

type listSoilRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listSoil(ctx *gin.Context) {
	var req listSoilRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	arg := db.ListSoilParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	soils, err := server.store.ListSoil(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, soils)
}

type deleteSoilRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) deleteSoil(ctx *gin.Context) {
	var req deleteSoilRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	err := server.store.DeleteSoil(ctx, req.ID)
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

type updateSoilRequest struct {
	ID            int64  `json:"id"`
	SubstanseName string `json:"substanse_name"`
}

func (server *Server) updateSoil(ctx *gin.Context) {
	var req updateSoilRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	arg := db.UpdateSoilParams{
		ID:            req.ID,
		SubstanseName: req.SubstanseName,
	}

	err := server.store.UpdateSoil(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, err)
}
