package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

// Status codes
const (
	timeFormat = "01-02-2006"
	success    = 200
	failure    = 400
)

// Healthz return ok without auth for testing connectivity
func Healthz(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(200)
	data := []byte("OK")
	w.Write(data)

}

// NewRecord create a new record nd store it in elastic
// TODO, after prototype want data to go elsewhere and then be sent to elastic
func NewRecord(w http.ResponseWriter, r *http.Request) {

	payload := Payload{
		Code:       success,
		W:          w,
		Error:      nil,
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		payload.processError(err, "reading body")
	}
	var records []Record
	err = json.Unmarshal(body, &records)
	if err != nil {
		payload.processError(err, "unmarshall body")
	}
	for _, record := range records {
		payload.NewRecords = append(payload.NewRecords, record)
		/*  Uncomment when read doing other indexing strategy
		currentTime := time.Now()
		indexName := "frisk-" + currentTime.Format(timeFormat)
		*/
		err := ESConnect(&payload)
		if err != nil {
			payload.processError(err, "ESConnect")
		}
		err = ESInsert(&payload)
		if err != nil {
			payload.processError(err, "ESInsert")
		}
	}
	payload.completeRequest()
}

func (p *Payload) completeRequest() {
	// Check for error message and end accordingly
	if p.Error != nil {
		p.Code = failure
		log.Println("Error: ", p.Error)
	}

	p.W.Header().Set("Content-Type", "application/json; charset=UTF-8")
	p.W.WriteHeader(p.Code)
	p.W.Write(p.Data)
}

func (p *Payload) processError(err error, msg string) {
	p.Error = errors.New(msg + ":" + err.Error())
	p.completeRequest()
}
