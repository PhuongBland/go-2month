package ginrestaurant

import (
	"github.com/gin-gonic/gin"
	"gofirst/compoment"
	"gofirst/modules/restaurant/restaurantmodel"
	"gorm.io/gorm"
)

func CreateRestaurant(appCtx compoment.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data restaurantmodel.RestaurantCreate

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(401, gin.H{
				"erro": err.Error(),
			})
			return
		}

		store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewCreateRestaurantBiz(store)

		if err := biz.CreateRestaurant(c.Request.Context(), &data); err != nil {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})
			return
		}
	}
}
