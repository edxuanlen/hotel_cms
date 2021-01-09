package server

import (
	"hotel_cms/api"
	"hotel_cms/common"
	"hotel_cms/middleware"
	"os"

	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 中间件, 顺序不能改
	r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	r.Use(middleware.Cors())

	// user 路由
	userV1 := r.Group("/api/v1")
	{
		userV1.POST("ping", api.Ping)

		// user
		// 用户登录
		userV1.POST("user/register", api.UserRegister)

		// 用户登录
		userV1.POST("user/login", api.UserLogin)

		// 退出登录
		userV1.DELETE("user/logout", middleware.AuthRequired(common.AllPermission), api.UserLogout)
	}

	// room order
	roomV1 := r.Group("/api/v1/room").Use(middleware.AuthRequired(common.RoomManagement))
	{
		// 房间列表
		roomV1.GET("infos", api.GetRooms)

		// 查询某个房间信息
		roomV1.GET("info/:id", api.GetRoom)

		// 新建房间
		roomV1.POST("add")

		// 更新房间状态(退房，清理等状态）
		roomV1.PUT("update/:id")

		// 订单
		roomV1.POST("order")

		// 查看订单
		roomV1.GET("order/:id")
		// 退单
		roomV1.DELETE("order/:id/delete")

		// 绑定卡号和房间号
		roomV1.POST("card/:cardId/:roomId")
	}

	// roomV2 版本2 可以同时存在，逐渐替代

	partnerV1 := r.Group("/api/v1/partner").Use(middleware.AuthRequired(common.DiplomaticManagement))
	{
		// 新建客户
		partnerV1.POST("create")

		// 更新客户状态
		partnerV1.PUT(":id/update")

		// 删除客户
		partnerV1.DELETE(":id/delete")
	}

	blackListV1 := r.Group("/api/v1/blacklist").Use(middleware.AuthRequired(common.AllPermission))
	{
		blackListV1.POST("add")
		blackListV1.GET(":id/info")
		blackListV1.PUT(":id/update")
	}

	commodityInventoryV1 := r.Group("/api/v1/commodity").Use(middleware.AuthRequired(common.WarehouseManagement))
	{

		commodityInventoryV1.GET("infos")

		// 获取商品信息
		commodityInventoryV1.GET("info/:id")

		// 添加商品
		commodityInventoryV1.POST("add")

		// 更新商品
		commodityInventoryV1.PUT(":id/update")

		// 删除商品
		commodityInventoryV1.DELETE(":id/delete")
	}

	accountBookV1 := r.Group("/api/v1/accountBook").Use(middleware.AuthRequired(common.FinancialManagement))
	{
		// 添加账单
		accountBookV1.POST("add")

		// 删除账单
		accountBookV1.DELETE(":id/delete")

		// 所有账单
		accountBookV1.GET("infos")

		// 查询账单
		accountBookV1.GET("info/:id")
	}
	return r
}
