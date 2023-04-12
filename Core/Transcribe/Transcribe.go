package Transcribe

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func Transcribe() {
	StartTime := time.Now()
	ExitCheck()
	values := map[string]string{"audio_url": AUDIO_URL}
	jsonData, _ := json.Marshal(values)

	// setup HTTP client and set header
	client := &http.Client{}
	req, err := http.NewRequest("POST", TRANSCRIPT_URL, bytes.NewBuffer(jsonData))
	req.Header.Set("content-type", "application/json")
	req.Header.Set("authorization", API_KEY)
	if err != nil {
		log.Println("Error creating request: ", err)
		os.Exit(4)
		return
	}
	res, err := client.Do(req)
	if err != nil {
		log.Println("Error creating response: ", err)
		os.Exit(5)
		return
	}

	defer res.Body.Close()

	// decode json and store it in a map
	var result map[string]interface{}
	json.NewDecoder(res.Body).Decode(&result)

	POLLING_URL = TRANSCRIPT_URL + "/" + result["id"].(string)
	duration := time.Since(StartTime)
	fmt.Println("\n\n[ TRANSCRIBE ]"+"\nAudio URL: "+AUDIO_URL+"\nTranscript URL: "+POLLING_URL+"\nError: False"+"\nTime: ", duration)
}
