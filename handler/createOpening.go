package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leobelini-studies/arthur-404-dev_api_moderna_golang/schemas"
)

// @BasePath /api/v1
// @Summary Create Opening
// @Description Create a new job opening
// @Accept json
// @Param request body CreateOpeningRequest true "Request body"
// @Produce json
// @Success 200 {object} CreateOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening [post]
func CreateOpeningHandler(ctx *gin.Context) {

	req := CreateOpeningRequest{}

	ctx.BindJSON(&req)

	if err := req.Validate(); err != nil {
		logger.Errorf("error validating request: %s", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	opening := schemas.Opening{
		Role:     req.Role,
		Company:  req.Company,
		Location: req.Location,
		Remote:   *req.Remote,
		Link:     req.Link,
		Salary:   req.Salary,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.Errorf("error creating opening: %s", err.Error())
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccess(ctx, "create-opening", opening)
}
