package Extract

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

var (
	InputPath  = "./Assets/Videos/Input/" + os.Args[1]
	OutputPath = "./Assets/Videos/Output/" + os.Args[2]
)

func Audio() {
	if _, err := os.Stat(OutputPath); err == nil {
		// Delete the file
		err := os.Remove(OutputPath)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	StartTime := time.Now()
	cmd := exec.Command("ffmpeg", "-i", InputPath, "-vn", "-acodec", "copy", OutputPath)
	err := cmd.Run()
	if err != nil {
		log.Println("Failed to extract audio. \nInput: " + InputPath + "\nOutput: " + OutputPath + "\nError:" + err.Error())
		return
	}
	duration := time.Since(StartTime)
	fmt.Println("[ AUDIO ]\n"+"Audio extracted successfully. \nInput: "+InputPath+"\nOutput: "+OutputPath+"\nError: False"+"\nTime: ", duration)
}
