package main

import (
	"grades-management/config"
	"grades-management/handlers"
	"grades-management/middleware"
	"grades-management/repository"
	"grades-management/services"

	"github.com/gin-gonic/gin"
)

func main()  {
	r:= gin.Default()

	db:=config.ConnectDb()
	Aikey:=config.Apikey
	

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

	assignHandler := handlers.NewAssignmentHandler(assignService)
	objectiveHandler := handlers.NewObjectiveHandler(objectiveService)
	progressHandler := handlers.NewProgressHandler(progressService)
	scoreHandler := handlers.NewScoreHandler(scoreService)
	studentHandler := handlers.NewStudentHandler(studentService)
	subjectHandler := handlers.NewSubjectsHandler(subjectService)


	api:= r.Group("/api")
	{
		
	}

}