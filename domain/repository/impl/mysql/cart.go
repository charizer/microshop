package mysql

import (
	"context"
	"github.com/jinzhu/gorm"
	"microshop/domain/entity"
	"microshop/domain/repository/impl"
	"microshop/infrastructure/db/mysql"
)

type CartRepo struct {
}

func (o CartRepo) List(ctx context.Context, page, size, userId int) ([]entity.Cart, error){
	list := []entity.Cart{}
	filter := map[string]interface{}{}
	filter["user_id"] = userId
	where, values, _ := mysql.SQLBuilder(filter)
	err := mysql.DB.Table(impl.CART_T_NAME).Where(where, values...).Order("create_time DESC").Offset(page*size).Limit(size).Find(&list).Error
	if err == gorm.ErrRecordNotFound {
		err = nil
	}
	return list,err
}

func (o CartRepo) All(ctx context.Context, filter map[string]interface{}) ([]entity.Cart, error){
	list := []entity.Cart{}
	where, values, _ := mysql.SQLBuilder(filter)
	err := mysql.DB.Table(impl.CART_T_NAME).Where(where, values...).Order("create_time DESC").Find(&list).Error
	if err == gorm.ErrRecordNotFound {
		err = nil
	}
	return list,err
}

func (o CartRepo) One(ctx context.Context, filter map[string]interface{}) (*entity.Cart, error) {
	where, values, _ := mysql.SQLBuilder(filter)
	var cart entity.Cart
	err := mysql.DB.Table(impl.CART_T_NAME).Where(where, values...).First(&cart).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &cart, err
}

func (o CartRepo) Insert(ctx context.Context, cart entity.Cart) (int, error) {
	err := mysql.DB.Table(impl.CART_T_NAME).Create(&cart).Error
	return cart.Id, err
}

func (o CartRepo) Update(ctx context.Context, filter, update map[string]interface{}) error {
	where, values, _ := mysql.SQLBuilder(filter)
	return mysql.DB.Table(impl.CART_T_NAME).Where(where, values...).UpdateColumn(update).Error
}

func (o CartRepo) Delete(ctx context.Context, filter map[string]interface{}) error {
	where, values, _ := mysql.SQLBuilder(filter)
	var result interface{}
	return mysql.DB.Table(impl.CART_T_NAME).Where(where, values...).Delete(result).Error
}

func (o CartRepo) Count(ctx context.Context, filter map[string]interface{}) (int64, error) {
	where, values, _ := mysql.SQLBuilder(filter)
	var result interface{}
	err := mysql.DB.Table(impl.CART_T_NAME).Where(where, values...).Count(result).Error
	if err != nil {
		return 0, err
	}
	return result.(int64), err
}