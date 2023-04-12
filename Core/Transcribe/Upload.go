package Transcribe

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	Client "github.com/XanaOG/Subtitles/Core/Client"
)

func Upload() {
	Config := Client.GetConfig(Client.ConfigFile)
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(Config.Keys))
	StartTime := time.Now()
	ExitCheck()
	data, err := ioutil.ReadFile("./Assets/Videos/Output/" + os.Args[2])
	if err != nil {
		log.Println("Error reading file to upload: ", err)
		os.Exit(2)
		return
	}
	API_KEY = Config.Keys[index]
	// Setup HTTP client and set header
	client := &http.Client{}
	req, _ := http.NewRequest("POST", UPLOAD_URL, bytes.NewBuffer(data))
	req.Header.Set("authorization", API_KEY)
	res, err := client.Do(req)
	if err != nil {
		log.Println("Error uploading file: ", err)
		os.Exit(3)
		return
	}
	// decode json and store it in a map
	var result map[string]interface{}
	json.NewDecoder(res.Body).Decode(&result)
	AUDIO_URL = result["upload_url"].(string)
	duration := time.Since(StartTime)
	fmt.Println("\n\n[ UPLOAD ] "+"\nUpload URL: "+UPLOAD_URL+"\nError: False"+"\nKey: "+API_KEY+"\nTime: ", duration)
}
