package api

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

	swaggerFiles "github.com/swaggo/files"

	ginSwagger "github.com/swaggo/gin-swagger"

	_ "auth-service/api/docs"
	"auth-service/api/handlers"
	"auth-service/api/middleware"
)

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func NewRouter(PollConn *grpc.ClientConn) *gin.Engine {
	router := gin.Default()
	h := handlers.NewHandler(PollConn)

	router.Use(CORSMiddleware())
	router.GET("/api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// #################### AUTH SERVICE ######################### //
	router.POST("/register", h.Register)
	router.POST("/confirm-registration", h.ConfirmRegistration)
	router.POST("/login", h.Login)
	router.POST("/forgot-password", h.ForgotPassword)
	router.POST("/recover-password", h.RecoverPassword)

	protected := router.Group("/", middleware.JWTMiddleware())
	protected.GET("/profile", h.Profile)
	router.GET("/user/:id", h.GetByID)

	// #################### USER SERVICE ######################### //
	for_user := protected.Group("/", middleware.IsUserMiddleware())
	for_user.POST("/send-answers", h.SendAnswers)

	// #################### POLLING SERVICE ######################### //
	for_admin := protected.Group("/", middleware.IsAdminMiddleware())

	for_admin.GET(("/results"), h.GetUserResultsInExcel)

	poll := for_admin.Group("/poll")
	{
		poll.POST("/", h.CreatePoll)
		poll.PUT("/", h.UpdatePoll)
		poll.DELETE("/:id", h.DeletePoll)
		poll.GET("/:id", h.GetPollByID)
	}

	for_admin.GET("/polls", h.GetAllPolls)

	question := for_admin.Group("/question")
	{
		question.POST("/", h.CreateQuestion)
		question.PUT("/", h.UpdateQuestion)
		question.DELETE("/:id", h.DeleteQuestion)
		question.GET("/:id", h.GetQuestionByID)
	}
	for_admin.GET("/questions/:poll_id", h.GetAllQuestions)

	return router
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		} else {
			c.Next()
		}
	}
}
