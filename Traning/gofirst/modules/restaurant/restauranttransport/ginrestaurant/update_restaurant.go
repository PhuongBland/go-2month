package ginrestaurant

import (
	"github.com/gin-gonic/gin"
	"gofirst/common"
	"gofirst/compoment"
	"gofirst/modules/restaurant/restaurantbiz"
	"gofirst/modules/restaurant/restaurantmodel"
	"gofirst/modules/restaurant/restaurantstorage"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func UpdateRestaurant(appCtx compoment.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data restaurantmodel.RestaurantUpdate
		id, err := c.ShouldBind(&data); err !=nil{
			c.JSON(401, gin.H{
				"error": err.Error(),
			})
			return

		}
		store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewUpdateRestaurantBiz(store)

		if err := biz.UpdateRestaurant(c.Request.Context(), &data); err != nil {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, common.SimpleSuccessRespone(true))
	}
}
