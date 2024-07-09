package middlewares

import (
	"errors"

	"github.com/graphql-go/graphql"
	"github.com/lai0xn/squid-tech/pkg/types"
)

func IsAuthenticated(p graphql.ResolveParams)(string,error){
   u := p.Context.Value("user")
   if u == nil {
    return "",errors.New("Not Authorized")
   }
   user := u.(*types.Claims)
   return user.ID,nil
}
