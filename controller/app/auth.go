package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var form request.Login
	if err := c.ShouldBindJSON(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}
	if err, user := service.UserService.Login(form); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		tokenData, err, _ := service.JwtService.CreateToken(service.AppGuardName, user)
		if err != nil {
			response.BusinessFail(c, err.Error())
		} else {
			response.Success(c, tokenData)
		}
	}
}
