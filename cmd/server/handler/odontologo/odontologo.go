package odontologo

import (
	"net/http"
	"strconv"

	"github.com/CamiloMartinez25/odontologia-go/core/web"
	"github.com/CamiloMartinez25/odontologia-go/internal/domain/odontologo"
	"github.com/gin-gonic/gin"
)

type Controlador struct {
	service odontologo.Service
}

func NewControladorOdontologo(service odontologo.Service) *Controlador {
	return &Controlador{
		service: service,
	}
}

// Odontologo godoc
// @Summary odontologo example
// @Description UpdateName odontologo
// @Tags odontologo
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /odontologos/:id [patch]
func (c *Controlador) UpdateName() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var request odontologo.RequestUpdateOdntologoName
		errBind := ctx.Bind(&request)
		if errBind != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "bad request")
			return
		}

		id := ctx.Param("id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "bad request param")
			return
		}

		odontologo, err := c.service.UpdateName(ctx, idInt, request.Nombre)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "%s", "internal server error")
			return
		}

		web.Succses(ctx, http.StatusOK, gin.H{
			"data": odontologo,
		})

	}
}
