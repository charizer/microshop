package mysql

import (
	"context"
	"github.com/jinzhu/gorm"
	"microshop/domain/entity"
	"microshop/domain/repository/impl"
	"microshop/infrastructure/db/mysql"
)

type OrderRepo struct {
}

func (o OrderRepo) List(ctx context.Context, page, size int, filter map[string]interface{}) ([]entity.Order, error){
	list := []entity.Order{}
	where, values, _ := mysql.SQLBuilder(filter)
	err := mysql.DB.Table(impl.ORDER_T_NAME).Where(where, values...).Order("create_time DESC").Offset(page*size).Limit(size).Find(&list).Error
	if err == gorm.ErrRecordNotFound {
		err = nil
	}
	return list,err
}

func (o OrderRepo) One(ctx context.Context, filter map[string]interface{}) (*entity.Order, error) {
	where, values, _ := mysql.SQLBuilder(filter)
	var order entity.Order
	err := mysql.DB.Table(impl.ORDER_T_NAME).Where(where, values...).First(&order).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &order, err
}

func (o OrderRepo) Insert(ctx context.Context, order entity.Order) (int, error) {
	err := mysql.DB.Table(impl.ORDER_T_NAME).Create(&order).Error
	return order.Id, err
}

func (o OrderRepo) Update(ctx context.Context, filter, update map[string]interface{}) error {
	where, values, _ := mysql.SQLBuilder(filter)
	return mysql.DB.Table(impl.ORDER_T_NAME).Where(where, values...).UpdateColumn(update).Error
}

func (o OrderRepo) Delete(ctx context.Context, filter map[string]interface{}) error {
	where, values, _ := mysql.SQLBuilder(filter)
	var result interface{}
	return mysql.DB.Table(impl.ORDER_T_NAME).Where(where, values...).Delete(result).Error
}

func (o OrderRepo) InsertOrderGoods(ctx context.Context, orderGoods entity.OrderGoods) (int, error) {
	err := mysql.DB.Table(impl.ORDER_GOODS_T_NAME).Create(&orderGoods).Error
	return orderGoods.Id, err
}

func (o OrderRepo) ListOrderGoods(ctx context.Context, page, size int, filter map[string]interface{}) ([]entity.OrderGoods, error){
	list := []entity.OrderGoods{}
	where, values, _ := mysql.SQLBuilder(filter)
	err := mysql.DB.Table(impl.ORDER_GOODS_T_NAME).Where(where, values...).Order("create_time DESC").Offset(page*size).Limit(size).Find(&list).Error
	if err == gorm.ErrRecordNotFound {
		err = nil
	}
	return list,err
}