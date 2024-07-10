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

type orgHandler struct {
	srv *services.OrgService
}

func NewOrgHandler() *orgHandler {
	return &orgHandler{
		srv: services.NewOrgService(),
	}
}

func (h *orgHandler) hasPerm(c echo.Context) (*db.OrganizationModel, error) {
	id := c.Param("id")
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*types.Claims)
	org, err := h.srv.GetOrg(id)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if org.OwnerID != claims.ID {
		return nil, echo.NewHTTPError(http.StatusUnauthorized, "you dont have the perms to perform this actin")
	}
	return org, nil

}

// @Summary	Get Organization endpoint
// @Tags		organizations
// @Accept		json
// @Produce	json
// @Success	200
// @Router		/organizations/org/get/:id [get]
func (h *orgHandler) Get(c echo.Context) error {
	id := c.Param("id")
	org, err := h.srv.GetOrg(id)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, org)
}

// @Summary	Search Organization endpoint
// @Tags		organizations
// @Accept		json
// @Produce	json
// @Param		name	query		string	false	"jhon doe org"
// @Success	200			{object}	string
// @Router		/organizations/org/search [get]
func (h *orgHandler) Search(c echo.Context) error {
	name := c.QueryParam("name")
	var err error
	var user interface{}

	user, err = h.srv.SearchOrg(name)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, user)
}

// @Summary	Change Organization Image endpoint
// @Tags		organizations
// @Accept		form/multipart
// @Produce	json
// @Param		Authorization	header	string	true	"Bearer token"
// @Param		image	formData	file	true	"file.png"
// @Success	200
// @Router		/organizations/org/:id/pfp [patch]
func (h *orgHandler) ChangePfp(c echo.Context) error {
	org, err := h.hasPerm(c)
	if err != nil {
		return err
	}
	file, err := c.FormFile("image")

	path := fmt.Sprintf("public/uploads/organizations/%s", filepath.Clean(file.Filename))
	f, err := os.Create(path)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	src, err := file.Open()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	io.Copy(f, src)
	_, err = h.srv.UpdateOrgImage(org.ID, path)
	return c.JSON(http.StatusOK, types.Response{
		"org_img": path,
	})
}

// @Summary	Change Organization Bg Image endpoint
// @Tags		organizations
// @Accept		form/multipart
// @Produce	json
// @Param		Authorization	header	string	true	"Bearer token"
// @Param		image	formData	file	true	"file.png"
// @Success	200
// @Router		/organizations/org/:id/bg [patch]
func (h *orgHandler) ChangeBg(c echo.Context) error {
	org, err := h.hasPerm(c)
	if err != nil {
		return err
	}
	file, err := c.FormFile("image")
	path := fmt.Sprintf("public/uploads/bgs/%s", filepath.Clean(file.Filename))
	f, err := os.Create(path)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	src, err := file.Open()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	io.Copy(f, src)
	_, err = h.srv.UpdateOrgBg(org.ID, path)
	return c.JSON(http.StatusOK, types.Response{
		"img": path,
	})
}

// @Summary	Update Organization endpoint
// @Tags		organizations
// @Accept		json
// @Produce	json
// @Param		Authorization	header	string	true	"Bearer token"
// @Param		body	body	types.OrgPayload	false	"jhon doe"
// @Success	200
// @Router		/organizations/org/update/:id [patch]
func (h *orgHandler) Update(c echo.Context) error {
	org, err := h.hasPerm(c)

	if err != nil {
		return err
	}

	var payload = types.OrgPayload{
		Name:        org.Name,
		Description: org.Description,
	}
	err = c.Bind(&payload)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())

	}
	_, err = h.srv.UpdateOrg(org.ID, payload)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())

	}
	return c.JSON(http.StatusOK, types.Response{
		"org": payload,
	})
}

// @Summary	Get Current Organization endpoint
// @Tags		organizations
// @Accept		json
// @Produce	json
// @Param		Authorization	header	string	true	"Bearer token"
// @Success	200
// @Router		/organizations/me [get]
func (h *orgHandler) MyOrgs(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*types.Claims)
	orgs, err := h.srv.GetUserOrgs(claims.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, types.Response{
		"orgs": orgs,
	})
}

// @Summary	Delete Organization endpoint
// @Tags		organizations
// @Accept		json
// @Produce	json
// @Param		Authorization	header	string	true	"Bearer token"
// @Success	200
// @Router		/organizations/org/delete/:id [delete]
func (h *orgHandler) Delete(c echo.Context) error {

	org, err := h.hasPerm(c)
	if err != nil {
		return err
	}
	deleted_id, err := h.srv.DeleteOrg(org.ID)

	return c.JSON(http.StatusOK, types.Response{
		"message": fmt.Sprintf("org deleted id : %s", deleted_id),
	})
}

// @Summary	Create Organization endpoint
// @Tags		organizations
// @Accept		json
// @Produce	json
// @Param		Authorization	header	string	true	"Bearer token"
// @Param		body	body	types.OrgPayload	false	"jhon doe"
// @Success	200
// @Router		/organizations/create [post]
func (h *orgHandler) Create(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*types.Claims)
	var payload types.OrgPayload
	err := c.Bind(&payload)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())

	}
	_, err = h.srv.CreateOrg(claims.ID, payload)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())

	}
	return c.JSON(http.StatusOK, types.Response{
		"org": payload,
	})
}
