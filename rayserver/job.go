package main

import (
	"fmt"
	"time"

	"github.com/surma-dump/gophernamedray/gnr"
)

type job struct {
	ID           string            `json:"id"`
	Scene        string            `json:"scene"`
	Composite    string            `json:"composite"`
	Start        time.Time         `json:"start"`
	End          time.Time         `json:"end"`
	Error        string            `json:"error"`
	Compositions map[string][]byte `json:"-"`
}

func (j *job) run() {
	defer func() {
		j.End = time.Now()
		if x := recover(); x != nil {
			switch t := x.(type) {
			case fmt.Stringer:
				j.Error = t.String()
			case error:
				j.Error = t.Error()
			case string:
				j.Error = t
			default:
				j.Error = fmt.Sprintf("%#v", x)
			}
		}
	}()

	scene, err := j.runSceneScript()
	if err != nil {
		j.Error = err.Error()
		return
	}
	irs := make([][]*gnr.InteractionResult, 0, Width*Height)
	for y := uint64(0); y < Height; y++ {
		for x := uint64(0); x < Width; x++ {
			ir := scene.TracePixel(x, y)
			irs = append(irs, ir)
		}
	}
	j.runCompScript(irs)
}
