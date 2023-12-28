package middleware

import (
	"errors"
	"fmt"
	"idstar-idp/rest-api/app/config"
	"idstar-idp/rest-api/app/util"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	authTokenKey string
}

func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{
		authTokenKey: config.GetConfigValue("authtoken.secret")}
}

func (a *AuthMiddleware) Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Printf("checking authorization for: %s %s", c.Request.Method, c.Request.URL)

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			util.SetErrorResponse(c, errors.New("header 'Authorization' not found"), http.StatusUnauthorized)
			return
		}

		tokenString := authHeader[len("Bearer "):]
		tokenPayload, err := util.ValidateToken(tokenString, a.authTokenKey)
		if err != nil {
			util.SetErrorResponse(c, err, http.StatusUnauthorized)
			return
		}

		path := c.Request.URL.EscapedPath()
		isAllowed := util.CheckPermission(tokenPayload, path, c.Request.Method)
		if !isAllowed {
			if strings.EqualFold(tokenPayload.RoleType, "admin") {
				util.SetErrorResponse(c, errors.New(fmt.Sprint("access forbidden for ", path, "; contact superadmin for assistance")), http.StatusForbidden)
				return
			} else {
				util.SetErrorResponse(c, errors.New(fmt.Sprint("access forbidden for ", path, "; contact admin for assistance")), http.StatusForbidden)
				return
			}
		}

		c.Next()
	}
}
