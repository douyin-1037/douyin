package auth

import (
	errno "douyin/common/code"
	config "douyin/common/conf"
	"douyin/common/constant"
	respond "douyin/types/coredto"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/gin-gonic/gin"
	"time"
)

var mw *jwt.GinJWTMiddleware

func Init() {
	var err error
	mw, err = jwt.New(&jwt.GinJWTMiddleware{
		Key:        []byte(config.JWT.Secret),
		Timeout:    config.JWT.Expires,
		MaxRefresh: time.Hour,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(int64); ok {
				return jwt.MapClaims{
					constant.IdentityKey: v,
				}
			}
			return jwt.MapClaims{}
		},
		TokenLookup:   "query: token, form: token, header: Authorization, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})
	if err != nil {
		klog.Fatal(err)
	}
}

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, err := GetUserID(c)
		if err != nil {
			respond.Error(c, err)
			c.Abort()
			return
		}
		c.Set(constant.IdentityKey, userID)
		c.Next()
	}
}

func GetUserID(c *gin.Context) (int64, error) {
	claims, err := mw.GetClaimsFromJWT(c)
	if err != nil {
		return 0, errno.UnauthorizedErr.WithMessage(err.Error())
	}
	tempUserID, ok := claims[constant.IdentityKey].(float64)
	userID := int64(tempUserID)
	if !ok || userID <= 0 {
		return 0, errno.UnauthorizedErr.WithMessage("jwt user_id error")
	}
	return userID, nil
}

func GenerateToken(userID int64) (string, error) {
	token, _, err := mw.TokenGenerator(userID)
	return token, err
}
