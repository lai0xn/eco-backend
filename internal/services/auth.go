package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/lai0xn/squid-tech/pkg/utils"
	"github.com/lai0xn/squid-tech/prisma"
	"github.com/lai0xn/squid-tech/prisma/db"
)

type AuthService struct{}

func NewAuthService()*AuthService{
  return &AuthService{}
}

func(s *AuthService)CreateUser(name string,email string,password string)error{
  ctx := context.Background()
  client := db.NewClient()
  if err := client.Prisma.Connect();err!= nil {
    return err
  }

  encrypted_password,err := utils.Encrypt(password)
  if err!= nil {
    return err
  }
  _,err = client.User.CreateOne(
    db.User.Email.Set(email),
    db.User.Name.Set(name),
    db.User.Bio.Set(""),
    db.User.Image.Set("uploads/default.jpg"),
    db.User.Password.Set(encrypted_password),
    ).Exec(ctx)
  if err != nil {
    return err
  }
  return nil

}


func(s *AuthService)CheckUser(email string,password string)(*db.UserModel,error){
  ctx := context.Background()
  client := db.NewClient()
  if err := client.Prisma.Connect();err!= nil {
    panic(err)
  }

  user,err := prisma.Client.User.FindUnique(
    db.User.Email.Equals(email),
    ).Exec(ctx)
  if err != nil {
    return nil,errors.New("Wrong Credentials")
  }
  fmt.Println(user.Email)
  enc_pass := user.Password
  err = utils.CheckPassword(enc_pass,password)
  if err != nil {
    return nil,errors.New("Wrong Credetials")
  }
  return user,nil

}
