package restaurantbiz

import (
	"context"
	"errors"
	"gofirst/modules/restaurant/restaurantmodel"
)

type CreateRestaurantBiz struct {
	store CreteRestaurantStore
}

func NewCreateRestaurantBiz(store CreateRestaurantStore) *CreateRestaurantBiz {
	return &CreateRestaurantBiz{store: store}
}

func (biz *CreateRestaurantBiz) CreateRestaurant(ctx context.Context, data *restaurantmodel.RestaurantCreate) error {
	if err := data.Vadidate(); err != nil {
		return err
	}
	err := biz.store.Create(ctx, data)
	return err
}
