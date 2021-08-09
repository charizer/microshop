package mysql

import (
	"context"
	"github.com/jinzhu/gorm"
	"microshop/domain/entity"
	"microshop/domain/repository/impl"
	"microshop/infrastructure/db/mysql"
)

type GoodsRepo struct {
}

func (o GoodsRepo) List(ctx context.Context, page, size int, filter map[string]interface{}) ([]entity.SampleGoods, error) {
	list := []entity.SampleGoods{}
	where, values, _ := mysql.SQLBuilder(filter)
	err := mysql.DB.Table(impl.GOODS_T_NAME).Where(where, values...).Order("sort_order ASC").Offset(page * size).Limit(size).Find(&list).Error
	if err == gorm.ErrRecordNotFound {
		err = nil
	}
	return list, err
}

func (o GoodsRepo) One(ctx context.Context, id int) (*entity.Goods, error) {
	filter := map[string]interface{}{}
	filter["id"] = id
	where, values, _ := mysql.SQLBuilder(filter)
	var goods entity.Goods
	err := mysql.DB.Table(impl.GOODS_T_NAME).Where(where, values...).First(&goods).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &goods, err
}

func (o GoodsRepo) GetGallerys(ctx context.Context, page, size, id int) ([]entity.GoodsGallery, error) {
	list := []entity.GoodsGallery{}
	filter := make(map[string]interface{})
	filter["goods_id"] = id
	where, values, _ := mysql.SQLBuilder(filter)
	err := mysql.DB.Table(impl.GOODS_GALLERY_T_NAME).Where(where, values...).Order("sort_order ASC").Offset(page * size).Limit(size).Find(&list).Error
	if err == gorm.ErrRecordNotFound {
		err = nil
	}
	return list, err
}

func (o GoodsRepo) GetAttrs(ctx context.Context, page, size, id int) ([]entity.GoodsAttr, error) {
	list := []entity.GoodsAttr{}
	filter := make(map[string]interface{})
	filter["goods_id"] = id
	where, values, _ := mysql.SQLBuilder(filter)
	err := mysql.DB.Table(impl.GOODS_ATTRIBUTE_T_NAME).Where(where, values...).Offset(page * size).Limit(size).Find(&list).Error
	if err == gorm.ErrRecordNotFound {
		err = nil
	}
	return list, err
}

func (o GoodsRepo) GetProducts(ctx context.Context, page, size, id int) ([]entity.GoodsProduct, error) {
	list := []entity.GoodsProduct{}
	filter := make(map[string]interface{})
	filter["goods_id"] = id
	where, values, _ := mysql.SQLBuilder(filter)
	err := mysql.DB.Table(impl.GOODS_PRODUCT_T_NAME).Where(where, values...).Offset(page * size).Limit(size).Find(&list).Error
	if err == gorm.ErrRecordNotFound {
		err = nil
	}
	return list, err
}

func (o GoodsRepo) GetProduct(ctx context.Context,filter map[string]interface{})(*entity.GoodsProduct, error){
	where, values, _ := mysql.SQLBuilder(filter)
	var product entity.GoodsProduct
	err := mysql.DB.Table(impl.GOODS_PRODUCT_T_NAME).Where(where, values...).First(&product).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &product, err
}
