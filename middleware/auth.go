package middleware

import (
	"blog/config"
	"blog/responses"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type Claim struct {
	UserId   uint   `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

const (
	CTXUserIdKey   = "userId"
	CTXUsernameKey = "username"
)

func NewToken(cfg *config.Config, uid uint, username string) (string, error) {
	claims := &Claim{
		UserId:   uid,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(cfg.JWTTTLHours) * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   strconv.FormatUint(uint64(uid), 10),
		},
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(cfg.JWTSecret))
}

func AuthRequred(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		au := c.GetHeader("Authorization")
		if len(au) < 8 || au[:7] != "Bearer " {
			responses.JOSNError(c, http.StatusUnauthorized, "unauthorized", "missing or invalid Authorization header")
			return
		}

		tokenStr := au[7:]
		token, err := jwt.ParseWithClaims(tokenStr, &Claim{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(cfg.JWTSecret), nil
		})
		if err != nil || !token.Valid {
			responses.JOSNError(c, http.StatusUnauthorized, "unauthorized", "invalid token")
			return
		}

		claim, ok := token.Claims.(*Claim)
		if !ok {
			responses.JOSNError(c, http.StatusUnauthorized, "unauthorized", "invalid claims")
			return
		}
		c.Set(CTXUserIdKey, claim.UserId)
		c.Set(CTXUsernameKey, claim.Username)
		c.Next()
	}
}

func MustGetUserID(c *gin.Context) uint {
	if v, exists := c.Get(CTXUserIdKey); exists {
		if id, ok := v.(uint); ok {
			return id
		}
	}

	return 0
}
