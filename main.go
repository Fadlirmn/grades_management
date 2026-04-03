package main

import (
	"context"
	"grades-management/config"
	"grades-management/handlers"
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
	

	assignRepo:= repository.NewAssignmentRepository(db)
	objectiveRepo:= repository.NewObjectiveRepository(db)
	progressRepo:= repository.NewProgressRepository(db)
	scoreRepo:= repository.NewScoreRepository(db)
	studentRepo:= repository.NewStudentRepository(db)
	subjectRepo:= repository.NewSubjectRepository(db)
	
	assignService:= services.NewAssignService(assignRepo)
	objectiveService:= services.NewObjectiveService(objectiveRepo)
	progressService:= services.NewProgressService(progressRepo)
	scoreService:= services.NewScoresService(scoreRepo)
	studentService:= services.NewStudentService(studentRepo)
	subjectService:= services.NewSubjectService(subjectRepo)

	sheetsServices:= services.NewSheetsService(config.SheetId,config.CredentialPath)
	
	geminiWorker := worker.NewGeminiWorker(progressService)
	syncWorker := worker.NewSyncWorker(sheetsServices,progressService)


	progressHandler := handlers.NewProgressHandler(progressService,geminiWorker)
	assignHandler := handlers.NewAssignmentHandler(assignService)
	objectiveHandler := handlers.NewObjectiveHandler(objectiveService)
	scoreHandler := handlers.NewScoreHandler(scoreService)
	studentHandler := handlers.NewStudentHandler(studentService)
	subjectHandler := handlers.NewSubjectsHandler(subjectService)


	c := cron.New()

	// Run setiap jam 2 pagi
	c.AddFunc("*/5 * * * *", func() {
		log.Println("Running weekly batch job...")
		geminiWorker.ProccessScheduleBatch(context.Background())
		syncWorker.SyncSheetsToDb()
		
	})

	c.Start()


	r:= gin.Default()




	api:= r.Group("/api")
	{
		api.POST("/test",progressHandler.TriggerAnalyis)
		assign:=api.Group("/assign")
		{
			assign.GET("", assignHandler.GetAssignments)
			assign.POST("", assignHandler.CreateAssignment)
			assign.PUT("/:id", assignHandler.UpdateAssignment)
			assign.DELETE("/:id", assignHandler.DeleteAssignment)
		}
		objective:=api.Group("/objective")
		{
			objective.GET("", objectiveHandler.GetObjectives)
			objective.POST("", objectiveHandler.CreateObjective)
			objective.PUT("/:id", objectiveHandler.UpdateObjective)
			objective.DELETE("/:id", objectiveHandler.DeleteObjective)
		}
		progress:=api.Group("/progress")
		{
			progress.GET("", progressHandler.GetProgresss)
			progress.POST("", progressHandler.CreateProgress)
			progress.PUT("/:id", progressHandler.UpdateProgress)
			progress.DELETE("/:id", progressHandler.DeleteProgress)
		}
		score:=api.Group("/score")
		{
			score.GET("", scoreHandler.GetScores)
			score.POST("", scoreHandler.CreateScore)
			score.PUT("/:id", scoreHandler.UpdateScore)
			score.DELETE("/:id", scoreHandler.DeleteScore)
		}
		student:=api.Group("/student")
		{
			student.GET("", studentHandler.GetStudents)
			student.POST("", studentHandler.CreateStudent)
			student.PUT("/:id", studentHandler.UpdateStudent)
			student.DELETE("/:id", studentHandler.DeleteStudent)
		}
		subject:=api.Group("/subject")
		{
			subject.GET("", subjectHandler.GetSubjectss)
			subject.POST("", subjectHandler.CreateSubjects)
			subject.PUT("/:id", subjectHandler.UpdateSubjects)
			subject.DELETE("/:id", subjectHandler.DeleteSubjects)
		}
	}
	r.Run(":8080")

}