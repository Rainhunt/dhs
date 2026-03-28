package users

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/rainhunt/dhs/internal/config"
)

type handler struct {
	service *service
}

func newHandler(service *service) *handler {
	return &handler{service: service}
}

func (h *handler) getUser(c *echo.Context) error {
	claims, err := config.GetCustomClaims(c)
	if err != nil {
		return echo.ErrForbidden
	}

	id := c.Param("id")

	if claims.IsAdmin || claims.Subject == id {
		pgUuid, err := config.StringToPgUUID(id)
		if err != nil {
			return echo.ErrBadRequest
		}
		user, err := h.service.repo.getUserByID(c.Request().Context(), pgUuid)
		if err != nil {
			return echo.ErrNotFound
		}
		c.JSON(http.StatusOK, user)
	}

	return echo.ErrForbidden
}

func (h *handler) listUsers(c *echo.Context) error {
	claims, err := config.GetCustomClaims(c)
	if err != nil {
		return echo.ErrForbidden
	}

	if claims.IsAdmin {
		users, err := h.service.repo.listUsers(c.Request().Context(), 50, 0)
		if err != nil {
			return echo.ErrInternalServerError
		}
		c.JSON(http.StatusOK, users)
	}

	return echo.ErrForbidden
}

func (h *handler) createUser(c *echo.Context) error {
	var req createUserDTO
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}
	if err := req.validate(); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	jwt, err := h.service.createUser(
		c.Request().Context(),
		req.Email,
		req.Username,
		req.Pass,
	)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, jwt)
}

func (h *handler) login(c *echo.Context) error {
	var req loginDTO
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	jwt, err := h.service.generateJWT(
		c.Request().Context(),
		req.Email,
		req.Pass,
	)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "invalid request"})
	}

	return c.JSON(http.StatusOK, jwt)
}

func (h *handler) editUser(c *echo.Context) error {
	var req editUserDTO
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}
	if err := req.validate(); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	claims, err := config.GetCustomClaims(c)
	if err != nil {
		return echo.ErrForbidden
	}

	id := c.Param("id")

	if claims.IsAdmin || claims.Subject == id {
		pgUuid, err := config.StringToPgUUID(id)
		if err != nil {
			return echo.ErrBadRequest
		}

		user, err := h.service.repo.editUser(
			c.Request().Context(),
			req.Email,
			req.Username,
			pgUuid,
		)
		if err != nil {
			return echo.ErrInternalServerError.Wrap(err)
		}
		return c.JSON(http.StatusOK, user)
	}
	return echo.ErrForbidden
}

func (h *handler) editUserPass(c *echo.Context) error {
	var req editUserPassDTO
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}
	if err := req.validate(); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	claims, err := config.GetCustomClaims(c)
	if err != nil {
		return echo.ErrForbidden
	}

	id := c.Param("id")

	if claims.Subject == id {
		pgUuid, err := config.StringToPgUUID(id)
		if err != nil {
			return echo.ErrBadRequest
		}

		user, err := h.service.editUserPass(
			c.Request().Context(),
			req.Pass,
			pgUuid,
		)
		if err != nil {
			return echo.ErrInternalServerError.Wrap(err)
		}
		return c.JSON(http.StatusOK, user)
	}
	return echo.ErrForbidden
}

func (h *handler) deleteUser(c *echo.Context) error {
	claims, err := config.GetCustomClaims(c)
	if err != nil {
		return echo.ErrForbidden
	}

	id := c.Param("id")

	if claims.Subject == id {
		pgUuid, err := config.StringToPgUUID(id)
		if err != nil {
			return echo.ErrBadRequest
		}

		if err := h.service.repo.deleteUser(c.Request().Context(), pgUuid); err != nil {
			return echo.ErrInternalServerError.Wrap(err)
		}
		return c.JSON(http.StatusNoContent, "deleted user")
	}
	return echo.ErrForbidden
}
