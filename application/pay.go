package application

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"microshop/common"
	"microshop/domain/service"
	"net/http"
)

func PayNotify(c *gin.Context) {
	err := service.AlipayService{}.ParseNotify(c)
	if err != nil {
		c.JSON(http.StatusOK, "fail")
	}else{
		c.JSON(http.StatusOK, "success")
	}
}

func Pay(c *gin.Context) {
	errMsg := ""
	data, err := service.AlipayService{}.Pay()
	if err != nil {
		errMsg = fmt.Sprintf("do pay err:%s", err.Error())
		common.HandleError(c, http.StatusInternalServerError,common.S_ALIPAY_REQ_ERR, errMsg)
	}else{
		common.HandleSucc(c, http.StatusOK, "", data)
	}
}
