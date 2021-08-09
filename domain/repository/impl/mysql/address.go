package mysql

import (
	"context"
	"github.com/jinzhu/gorm"
	"microshop/domain/entity"
	"microshop/domain/repository/impl"
	"microshop/infrastructure/db/mysql"
)

type AddressRepo struct {
}

func (o AddressRepo) List(ctx context.Context, page, size int, userId int) ([]entity.Address, error){
	list := []entity.Address{}
	filter := map[string]interface{}{}
	filter["user_id"] = userId
	where, values, _ := mysql.SQLBuilder(filter)
	err := mysql.DB.Table(impl.ADDRESS_T_NAME).Where(where, values...).Order("is_default DESC").Offset(page*size).Limit(size).Find(&list).Error
	if err == gorm.ErrRecordNotFound {
		err = nil
	}
	return list,err
}

func (o AddressRepo) One(ctx context.Context, id, userId int) (*entity.Address, error) {
	filter := map[string]interface{}{}
	filter["id"] = id
	filter["user_id"] = userId
	where, values, _ := mysql.SQLBuilder(filter)
	var address entity.Address
	err := mysql.DB.Table(impl.ADDRESS_T_NAME).Where(where, values...).First(&address).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &address, err
}

func (o AddressRepo) Insert(ctx context.Context, address entity.Address) (int, error) {
	err := mysql.DB.Table(impl.ADDRESS_T_NAME).Create(&address).Error
	return address.Id, err
}

func (o AddressRepo) UpdateOne(ctx context.Context, id int, address entity.Address) error {
	return mysql.DB.Table(impl.ADDRESS_T_NAME).Where("id = ?", id).Update(address).Error
}

func (o AddressRepo) Update(ctx context.Context, filter, update map[string]interface{}) error {
	where, values, _ := mysql.SQLBuilder(filter)
	return mysql.DB.Table(impl.ADDRESS_T_NAME).Where(where, values...).UpdateColumn(update).Error
}

func (o AddressRepo) Delete(ctx context.Context, filter map[string]interface{}) error {
	where, values, _ := mysql.SQLBuilder(filter)
	var result interface{}
	return mysql.DB.Table(impl.ADDRESS_T_NAME).Where(where, values...).Delete(result).Error
}


