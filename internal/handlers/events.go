package handlers

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

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


// @Summary	Get event endpoint
// @Tags		events
// @Accept		json
// @Produce	json
// @Success	200
// @Router		/events/event/get/:id [get]
func (h *eventHandler) Get(c echo.Context) error {
	id := c.Param("id")
	org, err := h.srv.GetEvent(id)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK,org)
}

// @Summary	Get acheivment endpoint
// @Tags		events
// @Accept		json
// @Produce	json
// @Success	200
// @Router		/events/acheivment/get/:id [get]
func (h *eventHandler) GetAcheivment(c echo.Context) error {
	id := c.Param("id")
	org, err := h.srv.GetAcheivment(id)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK,org)
}

// @Summary	Get Post endpoint
// @Tags		events
// @Accept		json
// @Produce	json
// @Param page query string true "1"
// @Success	200
// @Router		/events [get]
func (h *eventHandler) GetPage(c echo.Context) error {
  var p int
  page := c.QueryParam("page")
  if page == "" {
    p = 1
  }
  p,err := strconv.Atoi(page)
  if err != nil {
     	return echo.NewHTTPError(http.StatusBadRequest, err)
  }
	org, err := h.srv.GetEvents(p)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK,org)
}

// @Summary	Search event endpoint
// @Tags		events
// @Accept		json
// @Produce	json
// @Param		name	query		string	false	"jhon doe org"
// @Success	200			{object}	string
// @Router		/events/event/search [get]
func (h *eventHandler) Search(c echo.Context) error {
	name := c.QueryParam("query")
	var err error
	var user interface{}

	user, err = h.srv.SearchEvent(name)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, user)
}



// @Summary	create event endpoint
// @Tags		events
// @Accept		json
// @Produce	json
// @Param		Authorization	header	string	true	"Bearer token"
// @Param		body	body	types.EventPayload	false "body"	
// @Success	200
// @Router		/events/create [post]
func (h *eventHandler) Create(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*types.Claims)
	var payload types.EventPayload
	err := c.Bind(&payload)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())

	}
  org,err := h.osrv.GetOrg(payload.OrgID)
  if err != nil {
    	return echo.NewHTTPError(http.StatusBadRequest, err.Error())

  }
  if org.OwnerID != claims.ID {
    	return echo.NewHTTPError(http.StatusBadRequest, errors.New("no perms to perform this action"))
  }
  result, err := h.srv.CreateEvent(payload)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())

	}
	return c.JSON(http.StatusOK, result)
}

// @Summary	create acheivment endpoint
// @Tags		events
// @Accept		json
// @Produce	json
// @Param		Authorization	header	string	true	"Bearer token"
// @Param		body	body	types.AcheivmentPayload	false "body"	
// @Success	200
// @Router		/events/acheivment/create [post]
func (h *eventHandler) CreateAcheivment(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*types.Claims)
	var payload types.AcheivmentPayload
	err := c.Bind(&payload)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())

	}
  org,err := h.osrv.GetOrg(payload.OrgID)
  if err != nil {
    	return echo.NewHTTPError(http.StatusBadRequest, err.Error())

  }
  if org.OwnerID != claims.ID {
    	return echo.NewHTTPError(http.StatusBadRequest, errors.New("no perms to perform this action"))
  }
  result, err := h.srv.CreateAcheivment(payload,org.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())

	}
	return c.JSON(http.StatusOK, result)
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
		"img": path,
	})
}


// @Summary	Delete acheivment endpoint
// @Tags		organizations
// @Accept		json
// @Produce	json
// @Param		Authorization	header	string	true	"Bearer token"
// @Success	200
// @Router		/events/acheivment/delete/:id [delete]
func (h *eventHandler) DeleteAcheivment(c echo.Context) error {
  id := c.Param("id")
  u := c.Get("user").(*jwt.Token)
  claims := u.Claims.(*types.Claims)
  acheivment,err := h.srv.GetAcheivment(id)
  if err != nil {
     return echo.NewHTTPError(http.StatusBadRequest,err)
  }
  if acheivment.OrgID != claims.ID {
    return echo.NewHTTPError(http.StatusUnauthorized,errors.New("no authorized to do this action"))
  }

	deleted_id, err := h.srv.DeleteAcheivment(id)
  if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, deleted_id)
}
