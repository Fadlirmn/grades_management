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
	api:=config.Apikey
	

	assignRepo:= repository.NewAssignmentRepository(db)
	objectiveRepo:= repository.NewObjectiveRepository(db)
	progressRepo:= repository.NewProgressRepository(db)
	scoreRepo:= repository.NewScoreRepository(db)
	studentRepo:= repository.NewStudentRepository(db)
	subjectRepo:= repository.NewSubjectRepository(db)
	
	
}