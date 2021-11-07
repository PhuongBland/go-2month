package ginrestaurant

import (
	"github.com/gin-gonic/gin"
	"gofirst/common"
	"gofirst/compoment"
	"gofirst/modules/restaurant/restaurantbiz"
	"gofirst/modules/restaurant/restaurantmodel"
	"gofirst/modules/restaurant/restaurantstorage"
	"net/http"
)

func ListRestaurant(appCtx compoment.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter restaurantmodel.Filter

		if err := c.ShouldBind(&filter); err != nil {
			c.JSON(401, gin.H{
				"erro": err.Error(),
			})
			return
		}
		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})
			return
		}

		store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewListRestaurantBiz(store)

		result, err := biz.ListRestaurantBiz(c.Request.Context(), &filter, &paging)
		if err != nil {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})
			return
		}
		store := restaurantstorage.NewSqlStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewListRestaurantBiz(store)
	}
}
