package restaurantbiz

import (
	"context"
	"errors"
	"gofirst/common"
	"gofirst/modules/restaurant/restaurantmodel"
)

type GetRestaurantBiz struct {
	store GetRestaurantBiz
}

func NewGetRestaurantBiz(store GetRestaurantBiz) *GetRestaurantBiz {
	return &GetRestaurantBiz{store: store}
}

func (biz *GetRestaurantBiz) GetRestaurant(
	ctx context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging) ([]restaurantmodel.Restaurant, error) {
	result, err := biz.store.GetRestaurant(ctx, filter, paging)

	return result, err
}
