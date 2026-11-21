package worker

import (
	"context"
	"fmt"
	"log"
	"grades-management/config"	
	"grades-management/models"
	
	"google.golang.org/genai"
)

type IProgressService interface {
    CompressJson(data []models.Progress) (string, error)
    SaveAIResults(jsonStr string) error
	CountPending()(int,error)
	GetPendingAnalysis(studentId int) ([]models.Progress, error)
}

type GeminiWorker struct{
	ProgService IProgressService
	ApiKey string
}

func NewGeminiWorker(ps IProgressService) *GeminiWorker  {
	return &GeminiWorker{
		ProgService: ps,
		ApiKey: config.Apikey,
	}
}

func (w *GeminiWorker) ProcessBatchWithGemini(ctx context.Context, data []models.Progress) {
    // Gunakan ctx dari parameter, jangan buat Background() baru di dalam
    
    // API Client setup (Gunakan SDK yang kamu pakai, pastikan API Key benar)
    client, err := genai.NewClient(ctx, &genai.ClientConfig{
        APIKey: w.ApiKey,
        Backend: genai.BackendGeminiAPI,
    })
    if err != nil {
        log.Printf("failed to create client: %v", err)
        return
    }

	model := "gemini-2.5-flash-lite"

	compressedJSON, _ := w.ProgService.CompressJson(data)

	prompt := fmt.Sprintf(`Subjek: %s. Topik: %s. Data: %s. Tugas: Berikan rekomendasi belajar singkat (maksimal 8 kata) untuk setiap ID unik yang memiliki skor akhir < 75. Gunakan ID yang sama persis. Jangan ada duplikat.`,data[0].SubjectName,data[0].ObjectiveDesc,compressedJSON)
	config := &genai.GenerateContentConfig{
		Temperature: genai.Ptr(float32(0.0)),
		ResponseMIMEType: "application/json",
		ResponseSchema: &genai.Schema{
			Type: genai.TypeArray,
			Items: &genai.Schema{
				Type: genai.TypeObject,
				Properties: map[string]*genai.Schema{
					"id":{Type: genai.TypeInteger},
					"rec":{Type: genai.TypeString},
				},
			},
		},
	}
	
	resp, err := client.Models.GenerateContent(ctx, model,genai.Text(prompt),config)
	if err != nil {
		log.Printf("Gemini Error: %v", err)
		return
	}
	
	if resp == nil {
        log.Printf("Gemini returned nil response")
        return
    }
	
	if len(resp.Candidates) > 0 {
        part := resp.Candidates[0].Content.Parts[0]

   
		rawJSON := part.Text
		if rawJSON == "" {
            rawJSON = fmt.Sprint(part)
        }

        log.Printf("Hasil AI: %v", rawJSON)

        
        err := w.ProgService.SaveAIResults(rawJSON)
    	if err != nil {
        log.Printf("Gagal Save ke DB: %v", err)
    	}
    }
	
}

func (w *GeminiWorker)ProccessScheduleBatch(ctx context.Context)  {
	count, _:= w.ProgService.CountPending()

	if count < 20 {
		log.Printf("batch skipped: only %s data( minimum 20 data)", count)
		return
	}
	data, err:= w.ProgService.GetPendingAnalysis(0)
	if err != nil ||len(data)==0 {
		return
	}
	w.ProcessBatchWithGemini(ctx,data)
}