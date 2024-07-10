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
)

type profileHandler struct {
	srv *services.ProfileService
}

func NewProfileHandler() *profileHandler {
	return &profileHandler{
		srv: services.NewProfileService(),
	}
}

// @Summary	Get Profile endpoint
// @Tags		profiles
// @Accept		json
// @Produce	json
// @Success	200
// @Router		/profiles/get/:id [get]
func (h *profileHandler) Get(c echo.Context) error {
	id := c.Param("id")
	user, err := h.srv.GetUser(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	user.Password = ""
	return c.JSON(http.StatusOK, types.Response{
		"user": user,
	})
}

// @Summary	Search Profile endpoint
// @Tags		profiles
// @Accept		json
// @Produce	json
// @Param		email		query		string	false	"example@gmail.com"
// @Param		full_name	query		string	false	"aymen charfaoui"
// @Success	200			{object}	string
// @Router		/profiles/search [get]
func (h *profileHandler) Search(c echo.Context) error {
	email := c.QueryParam("email")
	name := c.QueryParam("name")
	var err error
	var user interface{}
	if email == "" && name == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "at least provide one query param")
	}
	if email == "" {
		user, err = h.srv.SearchByName(name)
	}
	user, err = h.srv.GetUserByEmail(email)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, types.Response{
		"user": user,
	})
}

// @Summary	Change Profile Image endpoint
// @Tags		profiles
// @Accept		form/multipart
// @Produce	json
// @Param		Authorization	header	string	true	"Bearer token"
// @Param		image	formData	file	true	"file.png"
// @Success	200
// @Router		/profiles/profile/pfp [patch]
func (h *profileHandler) ChangePfp(c echo.Context) error {
	file, err := c.FormFile("image")
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*types.Claims)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	path := fmt.Sprintf("public/uploads/profiles/%s", filepath.Clean(file.Filename))
	f, err := os.Create(path)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	src, err := file.Open()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	io.Copy(f, src)
	_, err = h.srv.UpdateUserImage(claims.ID, path)
	return c.JSON(http.StatusOK, types.Response{
		"user": path,
	})
}

// @Summary	Change Profile Bg Image endpoint
// @Tags		profiles
// @Accept		form/multipart
// @Produce	json
// @Param		Authorization	header	string	true	"Bearer token"
// @Param		image	formData	file	true	"file.png"
// @Success	200
// @Router		/profiles/profile/bg [patch]
func (h *profileHandler) ChangeBg(c echo.Context) error {
	file, err := c.FormFile("image")
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*types.Claims)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
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
	_, err = h.srv.UpdateUserBg(claims.ID, path)
	return c.JSON(http.StatusOK, types.Response{
		"user": path,
	})
}

// @Summary	Update Profile endpoint
// @Tags		profiles
// @Accept		json
// @Produce	json
// @Param		Authorization	header	string	true	"Bearer token"
// @Param		body	body	types.ProfileUpdate	false	"jhon doe"
// @Success	200
// @Router		/profiles/profile/update [patch]
func (h *profileHandler) Update(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*types.Claims)
	u, err := h.srv.GetUser(claims.ID)
	fmt.Println(claims.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	adress, ok := u.Adress()
	if !ok {
		adress = ""
	}
	var payload = types.ProfileUpdate{
		Email:  claims.Email,
		Name:   claims.Name,
		Bio:    u.Bio,
		Phone:  u.Password,
		Adress: adress,
	}
	err = c.Bind(&payload)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())

	}
	fmt.Println(payload)
	_, err = h.srv.UpdateUser(claims.ID, payload)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())

	}
	return c.JSON(http.StatusOK, types.Response{
		"user": payload,
	})
}

// @Summary	Get Current Profile endpoint
// @Tags		profiles
// @Accept		json
// @Produce	json
// @Param		Authorization	header	string	true	"Bearer token"
// @Success	200
// @Router		/profiles/profile [get]
func (h *profileHandler) CurrentUser(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*types.Claims)
	u, err := h.srv.GetUser(claims.ID)
  if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
  p,err := h.srv.GetUser(u.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, p)
}

// @Summary	Delete Profile endpoint
// @Tags		profiles
// @Accept		json
// @Produce	json
// @Param		Authorization	header	string	true	"Bearer token"
// @Success	200
// @Router		/profiles/profile/delete [delete]
func (h *profileHandler) Delete(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*types.Claims)
	u, err := h.srv.DeleteUser(claims.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, types.Response{
		"message": fmt.Sprintf("user deleted id : %s", u),
	})
}
