package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strings"
	"time"
)

// HttpHandler handles requests
type HttpHandler struct {
	Dir string
}

// implement `ServeHTTP` method on `HttpHandler` struct
func (h HttpHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	url := req.URL.String()
	if url == "/cleanup" {
		fmt.Println(time.Now().UTC().Format(time.UnixDate) + ": Running e2e clean up")
		CleanUp(h, res)

		return
	}

	if strings.Index(url, "/build-exclusions?templateId=") == 0 {
		fmt.Println(time.Now().UTC().Format(time.UnixDate) + ": Building template exclusions")
		BuildExclusionTree(h, res, req)

		return
	}

	log.Fatalln("No handler found for url: " + url)
}

func CleanUp(h HttpHandler, res http.ResponseWriter) {
	cmd := exec.Command("psh", "e2e:cleanup")
	cmd.Dir = h.Dir

	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}

	// create response binary data
	data := []byte("success")
	// write `data` to response
	_, err = res.Write(data)
	if err != nil {
		log.Fatal(err)
	}
}

func BuildExclusionTree(h HttpHandler, res http.ResponseWriter, req *http.Request) {
	cmd := exec.Command("bin/console", "cupro-template:generate-tree", req.URL.Query().Get("templateId"))
	cmd.Dir = h.Dir

	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}

	// create response binary data
	data := []byte("success")
	// write `data` to response
	_, err = res.Write(data)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	pathPointer := flag.String("path", "", "Path to your shopware installation")
	flag.Parse()

	if len([]rune(*pathPointer)) == 0 {
		log.Fatal("Please specify a path with -path=/my/path")
	}

	fmt.Println("Starting in: " + *pathPointer)

	fmt.Println("Creating db dump")
	cmd := exec.Command("psh", "e2e:dump-db")
	cmd.Dir = *pathPointer

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	// create a new handler
	handler := HttpHandler{}
	handler.Dir = *pathPointer

	fmt.Println("Listening on port 8005")

	// listen and serve
	err = http.ListenAndServe(":8005", handler)
	if err != nil {
		log.Fatal(err)
	}
}
