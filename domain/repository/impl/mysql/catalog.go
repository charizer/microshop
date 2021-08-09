package mysql

import (
	"context"
	"github.com/jinzhu/gorm"
	"microshop/domain/entity"
	"microshop/domain/repository/impl"
	"microshop/infrastructure/db/mysql"
)

type CatalogRepo struct {
}

func (o CatalogRepo) List(ctx context.Context, page, size int) ([]entity.Catalog, error){
	list := []entity.Catalog{}
	filter := map[string]interface{}{}
	filter["parent_id"] = 0
	filter["is_show"] = 1
	where, values, _ := mysql.SQLBuilder(filter)
	err := mysql.DB.Table(impl.CATALOG_T_NAME).Where(where, values...).Order("sort_order ASC").Offset(page*size).Limit(size).Find(&list).Error
	if err == gorm.ErrRecordNotFound {
		err = nil
	}
	return list,err
}

func (o CatalogRepo) One(ctx context.Context, id int) (*entity.Catalog, error){
	filter := map[string]interface{}{}
	filter["id"] = id
	where, values, _ := mysql.SQLBuilder(filter)
	var catalog *entity.Catalog
	err := mysql.DB.Table(impl.CATALOG_T_NAME).Where(where, values...).First(&catalog).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return catalog, err
}


