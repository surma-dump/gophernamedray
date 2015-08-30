package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/nu7hatch/gouuid"
	"github.com/surma/httptools"
)

const (
	Width  = 600
	Height = 600
)

type jobMap struct {
	Map map[string]*job
	sync.RWMutex
}

var jobs = &jobMap{
	Map: map[string]*job{},
}

func main() {
	var (
		listen = flag.String("listen", "localhost:5000", "Address to bind webserver to")
		static = flag.String("static", "./static", "Path to static folder")
	)

	flag.Parse()

	log.Printf("Starting webserver on %s...", *listen)
	err := http.ListenAndServe(*listen, httptools.NewRegexpSwitch(map[string]http.Handler{
		"/jobs": httptools.MethodSwitch{
			"GET":  http.HandlerFunc(listJobs),
			"POST": http.HandlerFunc(createJob),
		},
		"/jobs/[0-9a-f-]+": httptools.List{
			httptools.DiscardPathElements(1),
			httptools.SilentHandlerFunc(checkJobValidity),
			httptools.MethodSwitch{
				"GET":    http.HandlerFunc(showJob),
				"DELETE": http.HandlerFunc(deleteJob),
			},
		},
		"/": http.FileServer(http.Dir(*static)),
	}))
	if err != nil {
		log.Fatalf("Error starting webserver on %s: %s", *listen, err)
	}
}

func checkJobValidity(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/")
	valid := false
	func() {
		jobs.RLock()
		defer jobs.RUnlock()
		_, valid = jobs.Map[id]
	}()
	if !valid {
		ErrorWithMessage(w, http.StatusNotFound)
		return
	}
	r.Header.Set("X-Job-ID", id)
}

func createJob(w http.ResponseWriter, r *http.Request) {
	j := &job{
		ID:    NewUUID(),
		Start: time.Now(),
	}
	err := json.NewDecoder(r.Body).Decode(j)
	if err != nil {
		ErrorWithMessage(w, http.StatusBadRequest)
		log.Printf("Error parsing body: %s", err)
		return
	}

	go runJob(j)

	func() {
		jobs.Lock()
		defer jobs.Unlock()
		jobs.Map[j.ID] = j
	}()
	http.Error(w, j.ID, http.StatusCreated)
}

func listJobs(w http.ResponseWriter, r *http.Request) {
	var list []string
	func() {
		jobs.RLock()
		defer jobs.RUnlock()
		list = make([]string, 0, len(jobs.Map))
		for key := range jobs.Map {
			list = append(list, key)
		}
	}()
	json.NewEncoder(w).Encode(map[string]interface{}{"result": list})
}

func showJob(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("X-Job-ID")
	var j *job
	func() {
		jobs.RLock()
		defer jobs.RUnlock()
		j = jobs.Map[id]
	}()

	json.NewEncoder(w).Encode(j)
}
func deleteJob(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("X-Job-ID")
	func() {
		jobs.Lock()
		defer jobs.Unlock()
		delete(jobs.Map, id)
	}()
	ErrorWithMessage(w, http.StatusNoContent)
}

func ErrorWithMessage(w http.ResponseWriter, code int) {
	http.Error(w, http.StatusText(code), code)
}

func NewUUID() string {
	id, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}
	return id.String()
}
