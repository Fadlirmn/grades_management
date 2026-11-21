package worker

import (
	"fmt"
	"grades-management/services"
	
	"context"
	"log"
	"strconv"
)

type SyncWorker struct{
	SheetsScv *services.SheetsService
	progRepo *services.ProgressService
}

func NewSyncWorker(s *services.SheetsService, r *services.ProgressService) *SyncWorker {
	return &SyncWorker{
		SheetsScv: s,
		progRepo: r,
	}
}


func (w *SyncWorker)SyncSheetsToDb()  {
	ctx := context.Background()

	readRange := "Sheets1!A2:E"
	resp, err:= w.SheetsScv.FetchSheetsData(ctx,readRange)
	if err != nil {
		log.Printf("failed parsing credential json, %v",err)
		return 
	}

	for _,row:= range resp {
		if len(row) <  5{continue}
		
		sId,_:=strconv.Atoi(fmt.Sprintf("%v",row[0]))
		oId,_:=strconv.Atoi(fmt.Sprintf("%v",row[1]))
		week,_:=strconv.Atoi(fmt.Sprintf("%v",row[2]))
		score,_:=strconv.Atoi(fmt.Sprintf("%v",row[3]))
		status:= fmt.Sprintf("%v", row[4])

		err := w.progRepo.UpsertFromSheets(sId,oId,week,score,status)
		if err != nil {
			log.Printf("failed save data from sheets %s",err)
		}
	}
	log.Printf("sync sheets to db done")
}
