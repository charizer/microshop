package mysql

import (
	"context"
	"github.com/jinzhu/gorm"
	"microshop/domain/entity"
	"microshop/domain/repository/impl"
	"microshop/infrastructure/db/mysql"
)

type LoginRepo struct {
}

func (o LoginRepo) One(ctx context.Context, mobile string) (*entity.UserInfo, error){
	filter := map[string]interface{}{}
	filter["mobile"] = mobile
	where, values, _ := mysql.SQLBuilder(filter)
	var userInfo entity.UserInfo
	err := mysql.DB.Table(impl.USER_T_NAME).Where(where, values...).First(&userInfo).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &userInfo, err
}

func (o LoginRepo) Insert(ctx context.Context, userInfo entity.UserInfo) error {
	return mysql.DB.Table(impl.USER_T_NAME).Create(&userInfo).Error
}

func (o LoginRepo) Update(mobile string, update map[string]interface{}) error {
	return mysql.DB.Table(impl.USER_T_NAME).Where("mobile = ?", mobile).UpdateColumn(update).Error
}
