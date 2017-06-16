package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/AgoraIO/AgoraDynamicKey/go/src/DynamicKey5"
)

func main() {
	http.HandleFunc("/", pingHandler)
	http.HandleFunc("/recording_key/", recordingKeyHandler)
	log.Println("Starting server on localhost:8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, `{"message": "OK"}`)
	return
}

func recordingKeyHandler(w http.ResponseWriter, r *http.Request) {
	appID := ""
	appCertificate := ""
	channelName := strings.TrimPrefix(r.URL.Path, "/recording_key/")

	uid := uint32(2882341273)
	randomInt := uint32(rand.Int31n(100))
	unixTs := uint32(time.Now().Unix())
	expiredTs := uint32(0)

	recordingKey, err := DynamicKey5.GenerateRecordingKey(appID, appCertificate, channelName, unixTs, randomInt, uid, expiredTs)
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, fmt.Sprintf(`{"recording_key": "%s"}`, recordingKey))
}
