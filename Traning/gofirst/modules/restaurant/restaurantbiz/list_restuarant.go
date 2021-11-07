package restaurantbiz

import (
	"context"
	"errors"
	"gofirst/common"
	"gofirst/modules/restaurant/restaurantmodel"
)

type ListRestaurantBiz struct {
	store CreteRestaurantStore
}

func NewListRestaurantBiz(store CreateRestaurantStore) *ListRestaurantBiz {
	return &ListRestaurantBiz{store: store}
}

func (biz *ListRestaurantBiz) listRestaurant(
	ctx context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging) ([]restaurantmodel.Restaurant, error) {
	result, err := biz.store.ListDataBycondition(ctx, nil, filter, paging)

	return result, err
}
