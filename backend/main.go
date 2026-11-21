package main

import (
	"context"
	"grades-management/config"
	"grades-management/handlers"
	"grades-management/middleware"
	"log"

	//"grades-management/middleware"
	"grades-management/repository"
	"grades-management/services"
	"grades-management/worker"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
)

func main()  {


	db:=config.ConnectDb()
	config.AiConnect()
	config.SheetsConnect()
	
	//repo
	assignRepo:= repository.NewAssignmentRepository(db)
	objectiveRepo:= repository.NewObjectiveRepository(db)
	progressRepo:= repository.NewProgressRepository(db)
	scoreRepo:= repository.NewScoreRepository(db)
	studentRepo:= repository.NewStudentRepository(db)
	subjectRepo:= repository.NewSubjectRepository(db)
	userRepo:= repository.NewUserRepository(db)
	rTokenRepo:= repository.NewRTokenRepository(db)
	

	//service
	assignService:= services.NewAssignService(assignRepo)
	objectiveService:= services.NewObjectiveService(objectiveRepo)
	progressService:= services.NewProgressService(progressRepo)
	scoreService:= services.NewScoresService(scoreRepo)
	studentService:= services.NewStudentService(studentRepo)
	subjectService:= services.NewSubjectService(subjectRepo)
	userService:= services.NewUserService(userRepo)
	authService:= services.NewAuthService(userRepo,rTokenRepo)
	//service-g-sheets, sycfromsheets
	sheetsServices:= services.NewSheetsService(config.SheetId,config.CredentialPath)
	syncWorker := worker.NewSyncWorker(sheetsServices,progressService)
	//service-ai-recom
	geminiWorker := worker.NewGeminiWorker(progressService)

	//handler
	progressHandler := handlers.NewProgressHandler(progressService,geminiWorker)
	assignHandler := handlers.NewAssignmentHandler(assignService)
	objectiveHandler := handlers.NewObjectiveHandler(objectiveService)
	scoreHandler := handlers.NewScoreHandler(scoreService)
	studentHandler := handlers.NewStudentHandler(studentService)
	subjectHandler := handlers.NewSubjectsHandler(subjectService)
	userHandler := handlers.NewUserHandler(userService)
	authHandler := handlers.NewAuthHandler(authService)


	c := cron.New()

	// Run setiap jam 2 pagi
	c.AddFunc("0 * * * *", func() {
		log.Println("Running weekly batch job...")
		geminiWorker.ProccessScheduleBatch(context.Background())
		syncWorker.SyncSheetsToDb()
		
	})

	c.Start()


	r:= gin.Default()




	api:= r.Group("/api")
	api.Use(middleware.AuthMiddleware())
	{
		assign := api.Group("/assign")
		{
			// Teacher & Admin bisa kelola, Parent & Student cuma bisa lihat
			assign.GET("", middleware.RoleMiddleware("admin", "teacher", "parent", "student"), assignHandler.GetAssignments)
			assign.POST("", middleware.RoleMiddleware("admin", "teacher"), assignHandler.CreateAssignment)
			assign.PUT("/:id", middleware.RoleMiddleware("admin", "teacher"), assignHandler.UpdateAssignment)
			assign.DELETE("/:id", middleware.RoleMiddleware("admin", "teacher"), assignHandler.DeleteAssignment)
		}

		// --- OBJECTIVE (Tujuan Pembelajaran) ---
		objective := api.Group("/objective")
		{
			objective.GET("", middleware.RoleMiddleware("admin", "teacher", "parent", "student"), objectiveHandler.GetObjectives)
			objective.POST("", middleware.RoleMiddleware("admin", "teacher"), objectiveHandler.CreateObjective)
			objective.PUT("/:id", middleware.RoleMiddleware("admin", "teacher"), objectiveHandler.UpdateObjective)
			objective.DELETE("/:id", middleware.RoleMiddleware("admin", "teacher"), objectiveHandler.DeleteObjective)
		}

		// --- PROGRESS & SCORE (Nilai & AI Recommendation) ---
		progress := api.Group("/progress")
		{
			progress.GET("", middleware.RoleMiddleware("admin", "teacher", "parent", "student"), progressHandler.GetProgresss)
			// Progress biasanya dibuat otomatis/oleh teacher
			progress.POST("", middleware.RoleMiddleware("admin", "teacher"), progressHandler.CreateProgress)
			progress.PUT("/:id", middleware.RoleMiddleware("admin", "teacher"), progressHandler.UpdateProgress)
			progress.DELETE("/:id", middleware.RoleMiddleware("admin", "teacher"), progressHandler.DeleteProgress)
		}

		score := api.Group("/score")
		{
			score.GET("", middleware.RoleMiddleware("admin", "teacher", "parent", "student"), scoreHandler.GetScores)
			score.POST("", middleware.RoleMiddleware("admin", "teacher"), scoreHandler.CreateScore)
			score.PUT("/:id", middleware.RoleMiddleware("admin", "teacher"), scoreHandler.UpdateScore)
			score.DELETE("/:id", middleware.RoleMiddleware("admin", "teacher"), scoreHandler.DeleteScore)
		}

		// --- STUDENT ---
		student := api.Group("/student")
		{
			// Admin & Teacher bisa lihat semua siswa, Parent/Student mungkin cuma profil sendiri (diatur di logic handler)
			student.GET("", middleware.RoleMiddleware("admin", "teacher", "parent", "student"), studentHandler.GetStudents)
			student.POST("", middleware.RoleMiddleware("admin"), studentHandler.CreateStudent) // Biasanya cuma Admin
			student.PUT("/:id", middleware.RoleMiddleware("admin"), studentHandler.UpdateStudent)
			student.DELETE("/:id", middleware.RoleMiddleware("admin"), studentHandler.DeleteStudent)
		}

		// --- SUBJECT ---
		subject := api.Group("/subject")
		{
			subject.GET("", middleware.RoleMiddleware("admin", "teacher", "parent", "student"), subjectHandler.GetSubjectss)
			subject.POST("", middleware.RoleMiddleware("admin"), subjectHandler.CreateSubjects)
			subject.PUT("/:id", middleware.RoleMiddleware("admin"), subjectHandler.UpdateSubjects)
			subject.DELETE("/:id", middleware.RoleMiddleware("admin"), subjectHandler.DeleteSubjects)
		}

		// --- USERS (Khusus Admin) ---
		users := api.Group("/users")
		{
			users.Use(middleware.RoleMiddleware("admin")) // Terapkan ke semua route di group ini
			users.GET("", userHandler.GetUsers)
			users.POST("", userHandler.CreateUser)
			users.PUT("/:id", userHandler.UpdateUser)
			users.DELETE("/:id", userHandler.DeleteUser)
		}
	}
	auth:= r.Group("/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
		auth.POST("/refresh", authHandler.Refresh)
	}
	r.Run(":8080")

}