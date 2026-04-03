package services

import (
	"encoding/json"
	"fmt"
	"grades-management/models"
	"grades-management/repository"

	"log"
	"strings"
)

type ProgressService struct{
	progressRepo repository.ProgressRepository
}

type CompressData struct{
	Subject string          `json:"subject"`
    Topic   string          `json:"topic"`
    Week    int             `json:"week"`
    Data    [][]interface{} `json:"data"`
}

func NewProgressService(progress repository.ProgressRepository) *ProgressService  {
	return & ProgressService{
		progressRepo: progress,
	}
}

func (s *ProgressService) GetProgress()([]models.Progress)  {
	return s.progressRepo.FindAllProgress()
}

func (s *ProgressService) CreateProgress(Progress models.Progress)  {
		s.progressRepo.SaveProgress(Progress)
}

func (s *ProgressService) UpdateProgress(id int,Progress models.Progress)  error{
	return s.progressRepo.UpdateProgress(id, Progress)
}
func (s *ProgressService) DeleteProgress(id int) error {
	return s.progressRepo.DeleteProgress(id)
}

func (s *ProgressService) CountPending()(int,error) {
	count,err := s.progressRepo.CountPending()
	if err != nil {
	 log.Printf("Not Counted : %s",err)
	}
	return count,err
}

func (s *ProgressService)FindAnalysisByStudentId(studentId int) ([] models.Progress,error) {
	result, err:= s.progressRepo.FindAnalysisByStudentId(studentId)
	if err != nil {
		log.Printf("Failed fetch data for id %d: %v", studentId, err)
		return nil, err
	}
	
	if len(result)==0{
		log.Printf("recommendation not found, for ID %d",studentId)
		return nil,nil
	}
	return result,nil
}

func (s *ProgressService) GetPendingAnalysis(studentId int) ([]models.Progress, error) {
    
	return s.progressRepo.FindPendingAnalysis(studentId)    
}

func (s *ProgressService) CompressJson(prog []models.Progress) (string, error) {
    if len(prog) == 0 {
        return "", nil
    }

    compress := CompressData{
        Subject: prog[0].SubjectName,
        Topic:     prog[0].ObjectiveDesc,
        Week:     prog[0].Week,
        Data:    make([][]interface{}, 0),
    }
    for _, g := range prog {
        // Ambil data yang dibutuhkan AI saja
        row := []interface{}{g.ProgressId, g.FinalScore} // Pakai g.Id (ID Row) bukan StudentId untuk update nanti
        compress.Data = append(compress.Data, row)
	}
    
    jsonData, err := json.Marshal(compress)
    if err != nil {
        return "", err
    }
    return string(jsonData), nil
}

func (s *ProgressService)SaveAIResults(JsonStr string)error  {
	start:= strings.Index(JsonStr,"[")
	end:= strings.Index(JsonStr,"]")

	if start==-1||end==-1 ||end < start{
		log.Printf("format json invalid %s",JsonStr)
		return fmt.Errorf("invalid json format from ai")
	}
	
	cleanJson:= JsonStr[start : end+1]
	
	var result []models.AIRecommendation

	if err:= json.Unmarshal([]byte(cleanJson),&result); err != nil {
		return  err
	}

	
	err:= s.progressRepo.UpdateBatchRecommendation(result)
	if err != nil {
		log.Printf("failed update id %s ",err)
	}
	
	log.Printf("has been updated: %d ",len(result))
	return nil
}

func (s *ProgressService) UpsertFromSheets(studentId, objId, week, score int, status string) error {
    return s.progressRepo.UpsertFromSheets(studentId, objId, week, score, status)
}