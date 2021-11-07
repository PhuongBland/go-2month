package restaurantstorage

import (
	"context"
	"gofirst/common"
	"gofirst/modules/restaurant/restaurantmodel"
)

func (s *sqlStore) ListDataByCondition(ctx context.Context,
	conditions map[string]interface{},
	filter *restaurantmodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) error {

	var result []restaurantmodel.RestaurantCreate
	db := s.db

	for i := range moreKeys {
		db = db.preload(moreKeys[i])
	}
	db = db.where(conditions)
	if v := filter; v != nil {
		if v.CityId > 0 {
			db = db.where("city_id= ?", v.CityId)
		}
	}
	if err := db.count(&paging.Total).Error; err != nil {
		return err
	}
	if err := db.Create(data).Error; err != nil {
		return err
	}
	return nil
}
