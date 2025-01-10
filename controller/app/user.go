package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
)

func Register(context *gin.Context) {

	var form request.Register
	if err := context.ShouldBindJSON(&form); err != nil {
		response.ValidateFail(context, request.GetErrorMsg(form, err))
		return
	}
	if err, user := service.UserService.Register(form); err != nil {
		response.BusinessFail(context, err.Error())
	} else {
		response.Success(context, user)
		return
	}
}
