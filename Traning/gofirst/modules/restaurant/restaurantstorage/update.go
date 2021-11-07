package restaurantstorage

import (
	"context"
	"gofirst/modules/restaurant/restaurantmodel"
)

func (s *sqlStore) UpdateData(
	ctx context.Context,
	data *restaurantmodel.RestaurantCreate,
) error {
	db := s.db

	if err := db.Where("id = ?", id).Updates(data).Error; err != nil {
		return err
	}
	return nil
}
