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

func (userService *userService) Login(params request.Login) (err error, user models.User) {
	err = global.DB.Where("mobile=?", params.Mobile).First(&user).Error
	if err != nil || !utils.BcryptMakeCheck([]byte(params.Password), user.Password) {
		err = errors.New("用户名不存在或密码错误！")
	}
	return
}

func (useService userService) Info(id string) (err error, user models.User) {
	err = global.DB.Where("id=?", id).First(&user).Error
	if err != nil {
		err = errors.New("数据不存在")
	}
	return
}
