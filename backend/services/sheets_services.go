package services

import(
	"context"
	"os"
	"fmt"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

type SheetsService struct{
	SpreadsheetID string
	Credentials string
}

func NewSheetsService(id, creds string) *SheetsService  {
	return &SheetsService{
		SpreadsheetID: id,
		Credentials: creds,
	}
}

func (s *SheetsService)FetchSheetsData(ctx context.Context, readRange string)([][]interface{}, error)  {
	data, err:= os.ReadFile(s.Credentials)
	if err != nil {
		return nil, fmt.Errorf("failed create sheets service %s ", err)
	}

	srv, err := sheets.NewService(ctx,option.WithAuthCredentialsJSON(option.ServiceAccount,data))
	if err != nil {
		return nil, fmt.Errorf("failed create sheets service: %v", err)
	}

	resp, err := srv.Spreadsheets.Values.Get(s.SpreadsheetID,readRange).Do()
	if err != nil {
		return nil, fmt.Errorf("failed parsing data %s ", err)
	}
	return resp.Values,nil
}