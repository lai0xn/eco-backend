package services

import (
	"context"
	"fmt"

	"github.com/lai0xn/squid-tech/pkg/types"
	"github.com/lai0xn/squid-tech/prisma"
	"github.com/lai0xn/squid-tech/prisma/db"
)


type ProfileService struct {}

func NewProfileService()*ProfileService{
  return &ProfileService{}
}

func (s *ProfileService)GetUser(id string) (*db.UserModel,error){
  ctx := context.Background()
  user,err := prisma.Client.User.FindUnique(
    db.User.ID.Equals(id),
  ).Exec(ctx)
  if err != nil {
    return nil,err 
  }
  return user,nil
}

func (s *ProfileService)GetUserByEmail(email string) (*db.UserModel,error){
  ctx := context.Background()
  user,err := prisma.Client.User.FindUnique(
    db.User.Email.Equals(email),
  ).Exec(ctx)
  if err != nil {
    return nil,err 
  }
  return user,nil
}

func (s *ProfileService)SearchByName(name string) ([]db.UserModel,error){
  ctx := context.Background()
  users,err := prisma.Client.User.FindMany(
    db.User.Name.Contains(name),
  ).Exec(ctx)
  if err != nil {
    return nil,err 
  }
  return users,nil
}

func (s *ProfileService)UpdateUser(id string,payload types.ProfileUpdate) (*db.UserModel,error){
  ctx := context.Background()
  users,err := prisma.Client.User.FindUnique(
    db.User.ID.Equals(id),
  ).Update(
    db.User.Email.Set(payload.Email),
    db.User.Name.Set(payload.Name),
    db.User.Bio.Set(payload.Bio),
  ).Exec(ctx)
  if err != nil {
    return nil,err 
  }
  return users,nil
}


func (s *ProfileService)UpdateUserImage(id string,path string) (string,error){
  fmt.Println(id)
  ctx := context.Background()
  user,err := prisma.Client.User.FindUnique(
    db.User.ID.Equals(id),
  ).Update(
    db.User.Image.Set(path),
  ).Exec(ctx)
  if err != nil {
    return "",err 
  }
  return user.Image,nil
}



func (s *ProfileService)DeleteUser(id string) (string,error){
  fmt.Println(id)
  ctx := context.Background()
  deleted,err := prisma.Client.User.FindUnique(
    db.User.ID.Equals(id),
  ).Delete().Exec(ctx)
  if err != nil {
    return "",nil 
  }
  fmt.Println(deleted.ID)
  return deleted.ID,nil
}
