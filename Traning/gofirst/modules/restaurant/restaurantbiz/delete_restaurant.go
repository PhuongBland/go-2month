package restaurantbiz

import (
	"context"
	"errors"
	"gofirst/modules/restaurant/restaurantmodel"
)

type DeleteRestaurantStore interface {
	FindDataByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		morekeys ...string,
	) (*restaurantmodel.Restaurant, error)
	UpdateData(ctx context.Context,
		id int,
		data *restaurantmodel.RestaurantUpdate,
	) error
}

type DeleteRestaurantBiz struct {
	store UpdateRestaurantStore
}

func NewDeleteRestaurantBiz(store DeleteRestaurantStore) *DeleteRestaurantBiz {
	return &DeleteRestaurantBiz{store: store}
}
func (biz *updateRestaurantBiz)(ctx context.Context, id int, data *restaurantmodel.RestaurantUpdate) error {
	olddata, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return err
	}
	if olddata.Status == 0 {
		return errors.New("data deleted")
	}
	if err := biz.store.UpdateData(ctx, id, data); err != nil {
		return err
	}
	return nil
}
