package middleware

import (
	config "AuthInGoService/config/env"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func JwtAuth(next http.Handler) http.Handler{
	return http.HandlerFunc(func (w http.ResponseWriter,r *http.Request){

		authHeader:=r.Header.Get("Authorization")

		if authHeader=="" {
			http.Error(w,"Authorization Handler is required",http.StatusUnauthorized)
			return
		}
		if !strings.HasPrefix(authHeader,"Bearer"){
			http.Error(w,"Bearer must be present",http.StatusUnauthorized)
			return
		}
		token:=strings.TrimSpace(strings.TrimPrefix(authHeader,"Bearer"))
		if token==""{
			http.Error(w,"Token is required",http.StatusUnauthorized)
			return
		}
		claims:=jwt.MapClaims{}

		_,err:=jwt.ParseWithClaims(token,&claims,func (token *jwt.Token) (interface {},error){
			return ([] byte(config.GetString("JWT_SECRET_KEY","TOKEN"))),nil;
		})
		if err!=nil{
			http.Error(w,"Invalid Token: "+err.Error(),http.StatusUnauthorized)
			return
		}
		UserId,ok:=claims["user_id"].(float64)
		email,okemail:=claims["email"].(string)
		if !ok || !okemail{
			http.Error(w,"Invalid Token: user_id or email not found",http.StatusUnauthorized)
			return
		}
		fmt.Println("Authenticated User ID:", int64(UserId)," Email:", email)
		next.ServeHTTP(w,r)
	})

}