package Transcribe

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
)

func Save() {
	client := &http.Client{}
	u := uuid.New()
	StartTime := time.Now()
	for {
		req, _ := http.NewRequest("GET", POLLING_URL, nil)
		req.Header.Set("content-type", "application/json")
		req.Header.Set("authorization", API_KEY)
		res, err := client.Do(req)

		if err != nil {
			log.Fatalln(err)
		}
		defer res.Body.Close()

		var result map[string]interface{}
		json.NewDecoder(res.Body).Decode(&result)

		if result["status"] == "completed" {
			text := result["text"].(string)
			CompletePath := fmt.Sprintf("%s%s", SubTitlePath, u)
			transcriptFile, err := os.Create(CompletePath + ".txt")
			if err != nil {
				panic(err)
			}
			err = ioutil.WriteFile(CompletePath+".txt", []byte(text), 0644)
			if err != nil {
				panic(err)
			}
			defer transcriptFile.Close()
			duration := time.Since(StartTime)
			fmt.Println("\n\n[ SAVE ]"+"\nTranscript URL: "+POLLING_URL+"\nSave Location: "+CompletePath+"\nTime: ", duration)
			break
		} else {
			time.Sleep(1 * time.Second)
			continue
		}
	}
}
