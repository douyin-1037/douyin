package impl

import (
	"douyin/types/bizdto"
	"github.com/gin-gonic/gin"
)

type UserAppServiceImpl struct {
}

func NewUserAppService() UserAppServiceImpl {
	return *new(UserAppServiceImpl)
}
func (UserAppServiceImpl) GetUser(c *gin.Context, appUserID int64, userID int64) (user *bizdto.User, err error) {
	return nil, nil
}
