package mysql

import (
	"context"
	"github.com/jinzhu/gorm"
	"microshop/domain/entity"
	"microshop/domain/repository/impl"
	"microshop/infrastructure/db/mysql"
)

type BannerRepo struct {
}

func (o BannerRepo) List(ctx context.Context, page, size int) ([]entity.Banner, error){
	list := []entity.Banner{}
	filter := map[string]interface{}{}
	filter["enabled"] = 1
	where, values, _ := mysql.SQLBuilder(filter)
	err := mysql.DB.Table(impl.BANNER_T_NAME).Where(where, values...).Order("sort_order ASC").Offset(page * size).Limit(size).Find(&list).Error
	if err == gorm.ErrRecordNotFound {
		err = nil
	}
	return list,err
}
