package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"runtime/debug"

	"github.com/gorilla/mux"
	"gocv.io/x/gocv"
)

var httpClient = http.Client{}

func main() {
	router := mux.NewRouter()
	router.Path("/detect").HandlerFunc(DetectHandler).Methods("GET")
	log.Print("server listening on 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func DetectHandler(w http.ResponseWriter, request *http.Request) {

	url := request.URL.Query().Get("url")

	log.Print("downloading image from ", url)

	imageResponse, err := httpClient.Get(url)

	if err != nil || imageResponse.StatusCode != 200 {
		//return 400
	}

	classifier := gocv.NewCascadeClassifier()

	if !classifier.Load("./classifiers/haarcascade_frontalface_default.xml") {
		//return 400
	}

	imageBytes, _ := ioutil.ReadAll(imageResponse.Body)

	log.Print("detecting faces ")

	imageMat, _ := gocv.IMDecode(imageBytes, gocv.IMReadUnchanged)

	rects := classifier.DetectMultiScale(imageMat)

	message := fmt.Sprintf("{\"message\": \" detected %d faces \"}", len(rects))

	w.Header().Add("Content-Type", "application/json")
	w.Write([]byte(message))

	classifier.Close()
	debug.FreeOSMemory()
}
