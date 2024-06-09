package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"payment-system-four/internal/api"
	"payment-system-four/internal/middleware"
	"payment-system-four/internal/ports"
	"time"
)

// SetupRouter is where router endpoints are called
func SetupRouter(handler *api.HTTPHandler, repository ports.Repository) *gin.Engine {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "GET", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r := router.Group("/")
	{
		r.GET("/", handler.Readiness)
		r.POST("/create", handler.RegisterUser)
		r.POST("/login", handler.LoginUser)
		r.POST("/createAdmin", handler.CreateAdminAccount)
		r.POST("/adminlogin", handler.AdminLogin)

	}
	authorizeUser := r.Group("/user")
	{

	}
	authorizeUser.Use(middleware.AuthorizeAdmin(repository.FindUserByEmail, repository.TokenInBlacklist))
	{
		authorizeUser.POST("/transfer", handler.UserTransferHandler)
		authorizeUser.POST("/addfunds", handler.DepositHandler)
		authorizeUser.GET("/balance", handler.BalanceCheck)
		authorizeUser.GET("/transaction", handler.UserTransactionHistory)
	}

	// authorizeAdmin authorizes all authorized users handlers
	authorizeAdmin := r.Group("/admin")
	authorizeAdmin.Use(middleware.AuthorizeAdmin(repository.FindUserByEmail, repository.TokenInBlacklist))
	{
		authorizeAdmin.GET("/getuser", handler.GetUserByEmail)
		authorizeAdmin.GET("/alltransactions", handler.AdminTransactionsHandler)

	}

	return router
}
