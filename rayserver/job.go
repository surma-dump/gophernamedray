package main

import (
	"log"
	"time"
)

type job struct {
	ID     string    `json:"id"`
	Scene  string    `json:"scene"`
	Shader string    `json:"shader"`
	Start  time.Time `json:"start"`
	End    time.Time `json:"end"`
}

func runJob(j *job) {
	defer func() {
		j.End = time.Now()
		if x := recover(); x != nil {
			log.Printf("Invalid: %s", x)
		}
	}()

	scene, err := runSceneScript(j.Scene)
	log.Printf("%#v\n> err=%s", scene, err)
}
