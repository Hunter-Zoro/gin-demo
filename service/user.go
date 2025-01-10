package service

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/models"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type userService struct {
}

var UserService = new(userService)

func (userService *userService) Register(params request.Register) (err error, user models.User) {
	var result = global.DB.Where("mobile=?", params.Mobile).Select("id").First(&models.User{})
	if result.RowsAffected != 0 {
		err = errors.New("手机号已存在")
		return
	}
	user = models.User{Name: params.Name, Mobile: params.Mobile, Password: utils.BcryptMake([]byte(params.Password))}
	err = global.DB.Create(&user).Error
	return
}
