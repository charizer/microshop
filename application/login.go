package application

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"microshop/common"
	"microshop/domain/entity"
	"microshop/domain/service"
	"microshop/infrastructure/cache"
	"microshop/utils"
	"net/http"
)

func SendSms(c *gin.Context){
	ctx := c.Request.Context()
	errMsg := ""
	var req entity.SmsReq
	c.BindJSON(&req)
	if !utils.VerifyMobile(req.Mobile) {
		errMsg = fmt.Sprintf("verfiy mobile:%s err",req.Mobile)
		common.HandleError(c, http.StatusBadRequest, common.C_MOBILE_ERR, errMsg)
		return
	}
	lastSend := req.Mobile + "_last"
	_, ok := cache.MemCachePool.Get(lastSend)
	if ok  {
		errMsg = fmt.Sprintf("mobile:%s send sms to much",req.Mobile)
		common.HandleError(c, http.StatusBadRequest, common.C_SEND_MUCH_ERR, errMsg)
		return
	}
	if !req.Valid {
		common.HandleSucc(c, http.StatusOK, "", entity.SmsResult{Code: http.StatusOK, Message: "0"})
		return
	}
	result, err := service.NewLoginService().SendSms(ctx, req.Mobile)
	if err != nil {
		common.HandleError(c, http.StatusBadRequest, common.C_MOBILE_ERR, err.Error())
		return
	}
	common.HandleSucc(c, http.StatusOK, "", result)
}

func Login(c *gin.Context){
	ctx := c.Request.Context()
	errMsg := ""
	var req entity.LoginReq
	c.BindJSON(&req)
	if req.UserInfo.Mobile == "" || len(req.Code) != 6 {
		errMsg = fmt.Sprintf("login req:%+v err",req)
		common.HandleError(c, http.StatusBadRequest, common.C_REQ_PARAM_ERR, errMsg)
		return
	}
	if req.UserInfo.Valid {
		code, ok := cache.MemCachePool.Get(req.UserInfo.Mobile)
		if !ok || code.(string) != req.Code {
			errMsg = "验证码不对"
			common.HandleError(c, http.StatusBadRequest, common.C_VERIFY_CODE_ERR, errMsg)
			return
		}
	}
	token, userInfo, err := service.NewLoginService().Login(ctx, c, req)
	if err != nil {
		errMsg = fmt.Sprintf("do login err:%s",err.Error())
		common.HandleError(c, http.StatusInternalServerError, common.S_MYSQL_ERR, errMsg)
	}
	common.HandleSucc(c, http.StatusOK, "", gin.H{"token": token, "userInfo": userInfo})
}

