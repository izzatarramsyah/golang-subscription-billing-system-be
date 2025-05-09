package routes

import (
	"subscription-billing-system/api/v1/handlers"
	"subscription-billing-system/repository"
	"subscription-billing-system/usecase"
	"subscription-billing-system/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
	"github.com/gin-contrib/cors"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // ganti sesuai domain frontend
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Inisialisasi layer
	userRepo := repository.NewUserRepository(db)
	planRepo := repository.NewPlanRepository(db)
	productRepo := repository.NewProductRepository(db)
	subscriptionRepo := repository.NewSubscriptionRepository(db)
	paymentRepo := repository.NewPaymentRepository(db)
	libraryItemRepo := repository.NewLibraryItemRepository(db)
	subscriberRepo := repository.NewSubscriberRepository(db)
	reportRepo := repository.NewRepository(db)
	reminderRepo := repository.NewReminderRepository(db)
	ebookRepo := repository.NewEbookRepository(db)
	dashboardRepo := repository.NewDashboardRepository(db)

	userUseCase := usecase.NewUserUseCase(userRepo)
	planUseCase := usecase.NewPlanUseCase(planRepo)
	productUseCase := usecase.NewProductUseCase(productRepo)
	subscriptionUseCase := usecase.NewSubscriptionUseCase(subscriptionRepo, planRepo, productRepo, ebookRepo)
	paymentUseCase := usecase.NewPaymentUseCase(paymentRepo)
	libraryItemUseCase := usecase.NewLibraryItemUseCase(libraryItemRepo)
	subscriberUseCase := usecase.NewSubscriberUseCase(subscriberRepo)
	reportUseCase := usecase.NewReportUseCase(reportRepo)
	reminderUseCase := usecase.NewReminderUseCase(reminderRepo)
	ebookUseCase := usecase.NewEbookUseCase(ebookRepo)
	dashboardUseCase := usecase.NewDashboardUseCase(dashboardRepo)

	authAPI := handlers.NewAuthAPI(userUseCase)
	userAPI := handlers.NewUserAPI(userUseCase)
	planAPI := handlers.NewPlanAPI(planUseCase)
	productAPI := handlers.NewProductAPI(productUseCase, ebookUseCase)
	subscriptionAPI := handlers.NewSubscriptionAPI(subscriptionUseCase)
	paymentAPI := handlers.NewPaymentAPI(paymentUseCase)
	libraryItemAPI := handlers.NewLibraryItemAPI(libraryItemUseCase)
	subscriberAPI := handlers.NewSubscriberAPI(subscriberUseCase)
	reportAPI := handlers.NewReportAPI(reportUseCase)
	reminderAPI := handlers.NewReminderAPI(reminderUseCase)
	ebookAPI := handlers.NewEbookAPI(ebookUseCase, subscriptionUseCase)
	dashboardAPI := handlers.NewDashboardAPI(dashboardUseCase)

	v1 := router.Group("/api/v1")
	{
		v1.POST("/register", authAPI.Register)
		v1.POST("/login", authAPI.Login) 

		userGroup := v1.Group("/user")
		userGroup.Use(middleware.AuthRequired())
		{
			userGroup.GET("/getListUsers",middleware.RoleAuthorization("admin"), userAPI.GetListUsers)
			userGroup.POST("/updateUser", userAPI.UpdateUser)
			userGroup.POST("/updateUserStatus", userAPI.UpdateUserStatus)
			userGroup.GET("/getByID/:id", userAPI.GetByID)
			userGroup.GET("/GetByRole/:role", userAPI.GetByRole)
			userGroup.GET("/getUser", userAPI.GetUser)
		}

		planGroup := v1.Group("/plan")
		planGroup.Use(middleware.AuthRequired())
		{
			planGroup.POST("/createPlan", planAPI.Create)
			planGroup.GET("/getListPlans", planAPI.List)
			planGroup.GET("/getByID/:id", planAPI.GetByID)
			planGroup.GET("/getByProductID/:id", planAPI.GetByProductID)
			planGroup.POST("/updatePlan", planAPI.Update)
			planGroup.DELETE("/deletePlan/:id", planAPI.Delete)
		}

		productGroup := v1.Group("/product")
		productGroup.GET("/getListProduct", productAPI.List)
		productGroup.Use(middleware.AuthRequired())
		{
			productGroup.POST("/createProduct", productAPI.Create)
			productGroup.GET("/getByID/:id", productAPI.GetByID)
			productGroup.POST("/updateProduct", productAPI.Update)
			productGroup.POST("/deleteProduct/:id", productAPI.Delete)
			productGroup.GET("/getByOwnerID", productAPI.GetByOwnerID)
			productGroup.POST("/UpdateStatusProduct/:id", productAPI.UpdateStatusProduct)
		}

		subscriptionGroup := v1.Group("/subscription")
		subscriptionGroup.Use(middleware.AuthRequired())
		{
			subscriptionGroup.POST("/Subscribe", subscriptionAPI.Subscribe)
			subscriptionGroup.POST("/MySubscription", subscriptionAPI.MySubscription)
			subscriptionGroup.GET("/Unsubscribe/:id", subscriptionAPI.Unsubscribe)
		}

		paymentGroup := v1.Group("/payment")
		paymentGroup.Use(middleware.AuthRequired())
		{
			paymentGroup.POST("/CreatePayment", paymentAPI.CreatePayment)
			paymentGroup.GET("/GetPayments", paymentAPI.GetPayments)
			paymentGroup.GET("/GetAllPaymentDetails", middleware.RoleAuthorization("admin"), paymentAPI.GetAllPaymentDetails)
			paymentGroup.POST("/updatePaymentStatus", paymentAPI.UpdatePaymentStatus)
		}

		libraryItemGroup := v1.Group("/library")
		libraryItemGroup.Use(middleware.AuthRequired())
		{
			libraryItemGroup.GET("/getByID", libraryItemAPI.GetByUser)
		}

		subscriberGroup := v1.Group("/subscriber")
		subscriberGroup.Use(middleware.AuthRequired())
		{
			subscriberGroup.GET("/getByOwnerID", subscriberAPI.GetSubscribersByOwner)
		}

		reportGroup := v1.Group("/report")
		reportGroup.Use(middleware.AuthRequired())
		{
			reportGroup.POST("/getRevenueReport", reportAPI.GetRevenueReport)
		}

		reminderGroup := v1.Group("/reminder")
		reminderGroup.Use(middleware.AuthRequired())
		{
			reminderGroup.GET("/getAll", reminderAPI.GetAll)
			reminderGroup.GET("/getByID/:id", reminderAPI.GetByID)
			reminderGroup.GET("/delete/:id", reminderAPI.Delete)
			reminderGroup.POST("/update", reminderAPI.Update)
			reminderGroup.POST("/create", reminderAPI.Create)
		}

		ebookGroup := v1.Group("/ebook")
		ebookGroup.Use(middleware.AuthRequired())
		{
			ebookGroup.POST("/upload", ebookAPI.UploadEbook)
			ebookGroup.GET("/getAll", ebookAPI.ListEbooks)
			ebookGroup.GET("/download/:id", ebookAPI.DownloadEbook)
			ebookGroup.GET("/getAccess/:id", ebookAPI.GetEbookAccess)
			ebookGroup.POST("/serve", ebookAPI.ServeEbook)
		}

		dashboardGroup := v1.Group("/dashboard")
		dashboardGroup.GET("/", dashboardAPI.Dashboard)

		// dashboardGroup.Use(middleware.AuthRequired())
		// {
		// 	dashboardGroup.GET("/", middleware.RoleAuthorization("admin", "owner"), dashboardAPI.Dashboard)
		// }

	}

	return router
}
