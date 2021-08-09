package application

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"microshop/common"
	"microshop/domain/entity"
	"microshop/domain/service"
	"microshop/utils"
	"net/http"
	"time"
)

func ListAddress(c *gin.Context){
	ctx := c.Request.Context()
	errMsg := ""
	userIdStr := c.GetString("userId")
	userId := utils.String2Int(userIdStr)
	list, err := service.NewAddressService().List(ctx, 0, 20, userId)
	if err != nil {
		errMsg = fmt.Sprintf("list address err:%s",err.Error())
		common.HandleError(c, http.StatusInternalServerError, common.S_MYSQL_ERR, errMsg)
		return
	}
	common.HandleSucc(c, http.StatusOK, "", list)
}

func DetailAddress(c *gin.Context){
	ctx := c.Request.Context()
	id := utils.GetIntParam(c, "id")
	errMsg := ""
	userIdStr := c.GetString("userId")
	userId := utils.String2Int(userIdStr)
	address := entity.Address{}
	if id == -1 {
		list, err := service.NewAddressService().List(ctx, 0, 20, userId)
		if err != nil {
			errMsg = fmt.Sprintf("detail address list address err:%s",err.Error())
			common.HandleError(c, http.StatusInternalServerError, common.S_MYSQL_ERR, errMsg)
			return
		}
		if len(list) > 0 {
			address = list[0]
		}
	}else{
		addr, err := service.NewAddressService().GetAddress(ctx, id, userId)
		if err != nil {
			errMsg = fmt.Sprintf("detail address get address:%d err:%s",id, err.Error())
			common.HandleError(c, http.StatusInternalServerError, common.S_MYSQL_ERR, errMsg)
			return
		}
		if addr != nil {
			address = *addr
		}
	}
	common.HandleSucc(c, http.StatusOK, "", address)
}

func SaveAddress(c *gin.Context){
	var req entity.Address
	c.BindJSON(&req)
	if req.Id == 0 {
		insertAddress(c, req)
	}else{
		updateAddress(c,req)
	}
}

func insertAddress(c *gin.Context, address entity.Address){
	now := time.Now().UnixNano() / 1e6
	userIdStr := c.GetString("userId")
	userId := utils.String2Int(userIdStr)
	ctx := c.Request.Context()
	errMsg := ""
	address.UserId = userId
	address.CreateTime = now
	address.UpdateTime = now
	id, err := service.NewAddressService().AddAddress(ctx, address)
	if err != nil {
		errMsg = fmt.Sprintf("add address:%d err:%s",err.Error())
		common.HandleError(c, http.StatusInternalServerError, common.S_MYSQL_ERR, errMsg)
		return
	}
	address.Id = id
	if *address.IsDefault{
		err = service.NewAddressService().UpdateAll(ctx, map[string]interface{}{
			"user_id": userId,
			"id <>": address.Id,
		}, map[string]interface{}{
			"is_default": false,
			"update_time": now,
		})
		if err != nil {
			errMsg = fmt.Sprintf("update address is_default err:%s",err.Error())
			common.HandleError(c, http.StatusInternalServerError, common.S_MYSQL_ERR, errMsg)
			return
		}
	}
	common.HandleSucc(c, http.StatusOK, "", address)
}

func updateAddress(c *gin.Context, address entity.Address) {
	now := time.Now().UnixNano() / 1e6
	userIdStr := c.GetString("userId")
	userId := utils.String2Int(userIdStr)
	ctx := c.Request.Context()
	errMsg := ""
	addr, err := service.NewAddressService().GetAddress(ctx, address.Id, userId)
	if err != nil {
		errMsg = fmt.Sprintf("get address:%d err:%s",address.Id, err.Error())
		common.HandleError(c, http.StatusInternalServerError, common.S_MYSQL_ERR, errMsg)
		return
	}
	if addr == nil {
		errMsg = fmt.Sprintf("address:%d not exsit",address.Id)
		common.HandleError(c, http.StatusBadRequest, common.C_ADDR_NOT_EXSIT_ERR, errMsg)
		return
	}
	address.UserId = userId
	address.UpdateTime = now
	address.CreateTime = addr.CreateTime
	err = service.NewAddressService().UpdateAddress(ctx, address.Id, address)
	if err != nil {
		errMsg = fmt.Sprintf("update address:%d err:%s",address.Id, err.Error())
		common.HandleError(c, http.StatusInternalServerError, common.S_MYSQL_ERR, errMsg)
		return
	}
	if *address.IsDefault{
		err = service.NewAddressService().UpdateAll(ctx, map[string]interface{}{
			"user_id": userId,
			"id <>": address.Id,
		}, map[string]interface{}{
			"is_default": false,
			"update_time": now,
		})
		if err != nil {
			errMsg = fmt.Sprintf("update address is_default err:%s",err.Error())
			common.HandleError(c, http.StatusInternalServerError, common.S_MYSQL_ERR, errMsg)
			return
		}
	}
	common.HandleSucc(c, http.StatusOK, "", address)
}

func DeleteAddress(c *gin.Context){
	errMsg := ""
	ctx := c.Request.Context()
	idStr := c.DefaultQuery("id", "")
	id := utils.String2Int(idStr)
	if id == -1 {
		errMsg = fmt.Sprintf("delete address id:%s err", idStr)
		common.HandleError(c, http.StatusBadRequest, common.C_REQ_PARAM_ERR, errMsg)
		return
	}
	userIdStr := c.GetString("userId")
	userId := utils.String2Int(userIdStr)
	err := service.NewAddressService().DeleteAddress(ctx, map[string]interface{}{
		"user_id": userId,
		"id": id,
	})
	if err != nil {
		errMsg = fmt.Sprintf("delete address:%d err:%s",id, err.Error())
		common.HandleError(c, http.StatusInternalServerError, common.S_MYSQL_ERR, errMsg)
		return
	}
	common.HandleSucc(c, http.StatusOK, "", gin.H{})
}