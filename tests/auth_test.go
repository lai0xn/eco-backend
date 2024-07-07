package tests

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/lai0xn/squid-tech/internal/handlers"
	"github.com/stretchr/testify/assert"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
    b := make([]byte, n)
    for i := range b {
        b[i] = letterBytes[rand.Intn(len(letterBytes))]
    }
    return string(b)
}

var randEmail = fmt.Sprintf("%s@gmail.com",RandStringBytes(8))

func TestRegisteration(t *testing.T){
   e:=echo.New()
   userJson := `{
    "name":"jhon doe",
    "email":"`+randEmail+`",
    "password":"test123",
    "gender":true
   }`
   fmt.Println(userJson)

   req := httptest.NewRequest(http.MethodPost,"/",strings.NewReader(userJson))
   req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
   rec := httptest.NewRecorder()
   c :=e.NewContext(req,rec)
   handler := handlers.NewAuthHandler()
   if assert.NoError(t,handler.Register(c)){
      assert.Equal(t,http.StatusCreated,rec.Code)
  }
}

func TestLogin(t *testing.T){
   e:=echo.New()
   userJson := `{
    "email":"`+randEmail+`",
    "password":"test123"
   }`
   fmt.Println(userJson)
   req := httptest.NewRequest(http.MethodPost,"/",strings.NewReader(userJson))
   req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
   rec := httptest.NewRecorder()
   c :=e.NewContext(req,rec)
   handler := handlers.NewAuthHandler()
   if assert.NoError(t,handler.Login(c)){
      assert.Equal(t,http.StatusCreated,rec.Code)
  }
}


