package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leobelini-studies/arthur-404-dev_api_moderna_golang/schemas"
)

func UpdateOpeningHandler(ctx *gin.Context) {

	req := UpdateOpeningRequest{}
	id := ctx.Query("id")

	ctx.BindJSON(&req)

	if err := req.Validate(); err != nil {
		logger.Errorf("error validating request: %s", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	opening := schemas.Opening{}
	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("opening with id: %s not found", id))
		return
	}

	if req.Role != "" {
		opening.Role = req.Role
	}

	if req.Company != "" {
		opening.Company = req.Company
	}

	if req.Location != "" {
		opening.Location = req.Location
	}

	if req.Remote != nil {
		opening.Remote = *req.Remote
	}

	if req.Link != "" {
		opening.Link = req.Link
	}

	if req.Salary > 0 {
		opening.Salary = req.Salary
	}

	if err := db.Save(&opening).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error updating opening: %s", err.Error()))
		return
	}

	sendSuccess(ctx, "update-opening", opening)
}
