package users

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type handler struct {
	service *service
}

func newHandler(service *service) *handler {
	return &handler{service: service}
}

func (h *handler) createUser(c echo.Context) error {
	var req createUserDTO
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request"})
	}
	if err := req.validate(); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	jwt, err := h.service.createUser(
		c.Request().Context(),
		req.Email,
		req.Username,
		req.Pass,
	)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, jwt)
}

func (h *handler) login(c echo.Context) error {
	var req loginDTO
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request"})
	}

	jwt, err := h.service.generateJWT(
		c.Request().Context(),
		req.Email,
		req.Pass,
	)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "invalid request"})
	}

	return c.JSON(http.StatusOK, jwt)
}

// func (h *handler) getUsers(c echo.Context) error {
// 	auth := c.Request().Header.Get("auth")
//
// 	return c.JSON(http.StatusOK, users)
// }
