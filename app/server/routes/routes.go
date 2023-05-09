package routes

import (
	"context"
	controller "roller-tempo/controller/api"

	"github.com/labstack/echo/v4"
)

func RegisterAttractionRoutes(app *echo.Echo, ctx context.Context, controller *controller.AttractionController) {
	app.GET("/api/attractions", controller.Attractions)
	app.GET("/api/attractions/:id", controller.Attractions)
}

func RegisterRewardRoutes(app *echo.Echo, ctx context.Context, controller *controller.RewardController) {
	app.GET("/api/rewards", controller.Rewards)
	app.GET("/api/rewards/:id", controller.GetRewardByID)

	app.POST("/api/rewards/register", controller.CreateNewReward)
}

func RegisterUserRoutes(app *echo.Echo, ctx context.Context, controller *controller.UserController) {
	app.GET("/api/users", controller.Users)
}

/*

	MISSING ROUTES:

	app.GET("/api/userinfo/:id", controllers.GetUserInfo(ctx, db))
	app.GET("/api/attractioninfo/:name", controllers.GetAttractionInfo(ctx, db))
	app.GET("/api/rewardinfo/:name", controllers.GetRewardInfo(ctx, db))

	app.POST("/api/userregister/", controllers.PostUserRegister(ctx, db))
	app.POST("/api/usernextturn/", controllers.PostUserNextTurn(ctx, db))
	app.POST("/api/attractionregister/", controllers.PostAttractionRegister(ctx, db))
	app.POST("/api/rewardregister/", controllers.PostRewardRegister(ctx, db))
*/
