package turno

import (
	"net/http"
	"strconv"

	"github.com/CamiloMartinez25/odontologia-go/core/web"
	"github.com/CamiloMartinez25/odontologia-go/internal/domain/turno"
	"github.com/gin-gonic/gin"
)

type Controlador struct {
	service turno.Service
}

func NewControladorTurno(service turno.Service) *Controlador {
	return &Controlador{
		service: service,
	}
}

func (c *Controlador) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var request turno.RequestTurno

		err := ctx.Bind(&request)

		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "bad request")
			return
		}

		turn, err := c.service.Create(ctx, request)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "%s", "internal server error")
			return
		}

		web.Success(ctx, http.StatusOK, gin.H{
			"data": turn,
		})

	}
}

func (c *Controlador) CreateByPaciente() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var request turno.RequestTurnoByPaciente

		err := ctx.Bind(&request)

		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "bad request")
			return
		}

		turn, err := c.service.CreateByPaciente(ctx, request)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "%s", "internal server error")
			return
		}

		web.Success(ctx, http.StatusOK, gin.H{
			"data": turn,
		})

	}
}

func (c *Controlador) GetByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "id invalido")
			return
		}

		turno, err := c.service.GetByID(ctx, id)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "%s", "internal server error")
			return
		}

		web.Success(ctx, http.StatusOK, gin.H{
			"data": turno,
		})
	}
}
func (c *Controlador) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var request turno.RequestTurno

		errBind := ctx.Bind(&request)

		if errBind != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "bad request binding")
			return
		}

		id := ctx.Param("id")

		idInt, err := strconv.Atoi(id)

		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "bad request param")
			return
		}

		turn, err := c.service.Update(ctx, request, idInt)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "%s", "internal server error")
			return
		}

		web.Success(ctx, http.StatusOK, gin.H{
			"data": turn,
		})

	}
}
func (c *Controlador) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "id invalido")
			return
		}

		err = c.service.Delete(ctx, id)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "%s", "internal server error")
			return
		}

		web.Success(ctx, http.StatusOK, gin.H{
			"mensaje": "turno eliminado",
		})
	}
}

func (c *Controlador) GetByPacienteID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("dni"))
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "dni invalido")
			return
		}

		turnos, err := c.service.GetByPacienteID(ctx, id)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "internal server error")
			return
		}

		web.Success(ctx, http.StatusOK, gin.H{
			"data": turnos,
		})
	}
}

// turno godoc
// @Summary turno example
// @Description Update Any Subject on turno
// @Tags turno
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /turnos/:id [patch]
func (c *Controlador) UpdateSubject() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var request turno.RequestUpdateTurnoSubject
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

		turn, err := c.service.UpdateSubject(ctx, idInt, request)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "%s", "internal server error")
			return
		}

		web.Success(ctx, http.StatusOK, gin.H{
			"data": turn,
		})

	}
}
