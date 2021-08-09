package application

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"microshop/common"
	"microshop/domain/entity"
	"microshop/domain/service"
	"net/http"
)

func CatalogList(c *gin.Context){
	ctx := c.Request.Context()
	errMsg := ""
	list, err := service.NewCatalogService().List(ctx, 0, 4)
	if err != nil {
		errMsg = fmt.Sprintf("list catalog err:%s",err.Error())
		common.HandleError(c, http.StatusInternalServerError, common.S_MYSQL_ERR, errMsg)
		return
	}
	common.HandleSucc(c, http.StatusOK, "", entity.CataListResponse{CategoryList: list})
}
