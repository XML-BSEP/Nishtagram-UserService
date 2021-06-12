package middleware

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	logger "github.com/jelena-vlajkov/logger/logger"
	"log"
	"net/http"
	"os"
	"strings"
)

func AuthMiddleware(logger *logger.Logger) gin.HandlerFunc {
	return func (c *gin.Context) {
		role, err := ExtractRole(c.Request, logger)
		if err != nil {
			logger.Logger.Errorf("unauthorized request from IP address: %v", c.Request.Host)
			c.JSON(401, gin.H{"message" : "Unauthorized"})
			c.Abort()
			return
		}

		if role == "" {
			logger.Logger.Errorf("unauthorized request from IP address: %v", c.Request.Host)
			c.JSON(401, gin.H{"message" : "Unauthorized"})
			c.Abort()
			return
		}

		ok, err := enforce(role, c.Request.URL.Path, c.Request.Method, logger)

		if err != nil {
			log.Println(err)
			c.JSON(500, gin.H{"message" : "error occurred when authorizing user"})
			c.Abort()
			return
		}

		if !ok {
			logger.Logger.Errorf("forbidden request from IP address: %v", c.Request.Host)
			c.JSON(403, gin.H{"message" : "forbidden"})
			c.Abort()
			return
		}
		c.Next()
	}
}


func enforce(role string, obj string, act string, logger *logger.Logger) (bool, error) {
	enforcer, err := casbin.NewEnforcer("http/middleware/rbac_model.conf", "http/middleware/rbac_policy.csv")
	if err != nil {
		logger.Logger.Errorf("failed to load policy from file: %v", err)
		return false, fmt.Errorf("failed to load policy from DB: %w", err)
	}
	err = enforcer.LoadPolicy()
	if err != nil {
		logger.Logger.Errorf("failed to load policy from file: %v", err)
		return false, fmt.Errorf("failed to load policy from DB: %w", err)
	}
	ok, _ := enforcer.Enforce(role, obj, act)
	return ok, nil
}

func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2{
		return strArr[1]
	} else {
		if len(strArr) == 1 {
			if strArr[0] != "" {
				strArr2 := strings.Split(strArr[0], "\"")

				return strArr2[1]
			}
		}
	}
	return ""
}

func ExtractRole(r *http.Request, logger *logger.Logger) (string, error) {
	tokenString := ExtractToken(r)
	if tokenString == "" {
		return "ANONYMOUS", nil
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})/*
		if err != nil {
			return "", err
		}*/

	claims, ok := token.Claims.(jwt.MapClaims)

	if ok  {
		role, ok := claims["role"].(string)
		if !ok {
			logger.Logger.Error("error while reading role from token")
			return "ANONYMOUS", err
		}

		return strings.ToUpper(role), nil
	}
	return "ANONYMOUS", err
}

func ExtractUserId(r *http.Request, logger *logger.Logger) (string, error) {
	tokenString := ExtractToken(r)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})

	claims, ok := token.Claims.(jwt.MapClaims)

	if ok  {
		userId, ok := claims["user_id"].(string)
		if !ok {
			logger.Logger.Error("error while reading user id from token")
			return "", err
		}

		return userId, nil
	}
	return "", err
}