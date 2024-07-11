package handlers

import (
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

type PostsHandler struct {
	srv *services.PostService
}

func NewPostsHandler() *PostsHandler {
	return &PostsHandler{
		srv: services.NewPostService(),
	}
}

func (h *PostsHandler) hasPerm(c echo.Context) (*db.PostModel, error) {
	id := c.Param("id")
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*types.Claims)
	post, err := h.srv.GetPost(id)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if post.AuthorID != claims.ID {
		return nil, echo.NewHTTPError(http.StatusUnauthorized, "you dont have the perms to perform this actin")
	}
	return post, nil

}

// @Summary	Get Post endpoint
// @Tags		posts
// @Accept		json
// @Produce	json
// @Success	200
// @Router		/posts/post/get/:id [get]
func (h *PostsHandler) Get(c echo.Context) error {
	id := c.Param("id")
	org, err := h.srv.GetPost(id)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, org)
}

// @Summary	Get Post endpoint
// @Tags		posts
// @Accept		json
// @Produce	json
// @Param page query string true "1"
// @Success	200
// @Router		/posts [get]
func (h *PostsHandler) GetPage(c echo.Context) error {
	var p int
	page := c.QueryParam("page")
	if page == "" {
		p = 1
	}
	p, err := strconv.Atoi(page)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	org, err := h.srv.GetPosts(p)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, org)
}

// @Summary	Search Post endpoint
// @Tags		posts
// @Accept		json
// @Produce	json
// @Param		name	query		string	false	"jhon doe org"
// @Success	200			{object}	string
// @Router		/posts/post/search [get]
func (h *PostsHandler) Search(c echo.Context) error {
	name := c.QueryParam("query")
	var err error
	var user interface{}

	user, err = h.srv.SearchPost(name)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, types.Response{
		"user": user,
	})
}

// @Summary	Upload post Image endpoint
// @Tags		posts
// @Accept		form/multipart
// @Produce	json
// @Param		Authorization	header	string	true	"Bearer token"
// @Param		image	formData	file	true	"file.png"
// @Success	200
// @Router		/posts/post/:id/image [post]
func (h *PostsHandler) UploadImage(c echo.Context) error {
	post, err := h.hasPerm(c)
	if err != nil {
		return err
	}
	file, err := c.FormFile("image")

	path := fmt.Sprintf("public/uploads/posts/%s", filepath.Clean(file.Filename))
	f, err := os.Create(path)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	src, err := file.Open()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	io.Copy(f, src)
	_, err = h.srv.UploadImage(post.ID, path)
	return c.JSON(http.StatusOK, types.Response{
		"org_img": path,
	})
}

// @Summary	Update Post endpoint
// @Tags		posts
// @Accept		json
// @Produce	json
// @Param		Authorization	header	string	true	"Bearer token"
// @Param		body	body	types.OrgPayload	false	"jhon doe"
// @Success	200
// @Router		/posts/post/update/:id [patch]
func (h *PostsHandler) Update(c echo.Context) error {
	post, err := h.hasPerm(c)

	if err != nil {
		return err
	}

	var payload = types.PostPayload{
		Content:     post.Content,
		Description: post.Description,
	}
	err = c.Bind(&payload)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())

	}
	_, err = h.srv.UpdatePost(post.ID, payload)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())

	}
	return c.JSON(http.StatusOK, types.Response{
		"org": payload,
	})
}

// @Summary	Delete Post endpoint
// @Tags		posts
// @Accept		json
// @Produce	json
// @Param		Authorization	header	string	true	"Bearer token"
// @Success	200
// @Router		/posts/post/:id/delete [delete]
func (h *PostsHandler) Delete(c echo.Context) error {

	org, err := h.hasPerm(c)
	if err != nil {
		return err
	}
	deleted_id, err := h.srv.DeletePost(org.ID)

	return c.JSON(http.StatusOK, types.Response{
		"message": fmt.Sprintf("org deleted id : %s", deleted_id),
	})
}

// @Summary	create post endpoint
// @Tags		posts
// @Accept		json
// @Produce	json
// @Param		Authorization	header	string	true	"Bearer token"
// @Param		body	body	types.PostPayload	false "body"
// @Success	200
// @Router		/posts/create [post]
func (h *PostsHandler) Create(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*types.Claims)
	var payload types.PostPayload
	err := c.Bind(&payload)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())

	}
	_, err = h.srv.CreatePost(claims.ID, payload.Content, payload.Description)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())

	}
	return c.JSON(http.StatusOK, types.Response{
		"org": payload,
	})
}

// @Summary	comment post endpoint
// @Tags		posts
// @Accept		json
// @Produce	json
// @Param		Authorization	header	string	true	"Bearer token"
// @Param		body	body	types.CommentPayload	false "body"
// @Success	200
// @Router		/posts/comment [post]
func (h *PostsHandler) Comment(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*types.Claims)
	var payload types.CommentPayload
	err := c.Bind(&payload)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())

	}
	_, err = h.srv.CreateComment(claims.ID, payload.PostID, payload.Content)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())

	}
	return c.JSON(http.StatusOK, types.Response{
		"org": payload,
	})
}

// @Summary	delete comment endpoint
// @Tags		posts
// @Accept		json
// @Produce	json
// @Param		Authorization	header	string	true	"Bearer token"
// @Success	200
// @Router		/posts/comments/:id/delete [delete]
func (h *PostsHandler) DeleteComment(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*types.Claims)
	var payload types.CommentPayload

	err := c.Bind(&payload)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())

	}
	_, err = h.srv.CreateComment(claims.ID, payload.PostID, payload.Content)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())

	}
	return c.JSON(http.StatusOK, types.Response{
		"org": payload,
	})
}
