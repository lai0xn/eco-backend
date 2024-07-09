package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/lai0xn/squid-tech/internal/services"
	"github.com/lai0xn/squid-tech/pkg/types"
	"github.com/lai0xn/squid-tech/prisma/db"
)

func NewEventHandler() *eventHandler {
	return &eventHandler{
		srv: services.NewEventsService(),
    osrv : services.NewOrgService(),
	}
}

type eventHandler struct {
  srv *services.EventsService
  osrv *services.OrgService
}
func (h *eventHandler)hasPerm(c echo.Context) (*db.EventModel, error) {
	id := c.Param("id")
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*types.Claims)
	event, err := h.srv.GetEvent(id)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
  org,err := h.osrv.GetOrg(event.OrganizerID)
  if err != nil {
    return nil,echo.NewHTTPError(http.StatusBadGateway,err)
  }
	if org.OwnerID != claims.ID {
		return nil, echo.NewHTTPError(http.StatusUnauthorized, "you dont have the perms to perform this actin")
	}
	return event, nil

}


// @Summary	Add Event Image endpoint
// @Tags		events
// @Accept		form/multipart
// @Produce	json
// @Param		Authorization	header	string	true	"Bearer token"
// @Param		image	formData	file	true	"file.png"
// @Success	200
// @Router		/events/event/:id/upload [post]
func (h *eventHandler) AddImage(c echo.Context) error {
	org, err := h.hasPerm(c)
	if err != nil {
		return err
	}
	file, err := c.FormFile("file")

  if err != nil {
    return echo.NewHTTPError(http.StatusBadRequest,err)
  }

	path := fmt.Sprintf("public/uploads/events/%s", filepath.Clean(file.Filename))
	f, err := os.Create(path)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	src, err := file.Open()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	io.Copy(f, src)
	_, err = h.srv.AddImage(org.ID, path)
	return c.JSON(http.StatusOK, types.Response{
		"org_img": path,
	})
}

