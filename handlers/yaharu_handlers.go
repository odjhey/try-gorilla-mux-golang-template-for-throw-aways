package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func GetYaharu(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Yaharu!")
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "welcome to the homepage <br>")
	fmt.Fprint(w, `go check the <a href="/swagger/index.html">swagger UI</a>`)
	fmt.Println("endpoint hit")
}

func ReturnSSS(w http.ResponseWriter, r *http.Request) {
	const asdf = `asdfasdf`
	fmt.Fprint(w, asdf)
}

func ReturnHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `{"yo": "brow"}`)
}

type TPingResponse struct {
	Duration int    `json:"duration"`
	Message  string `json:"message"`
}
type PingBody struct {
	Timeout int    `json:"timeout"`
	Echo    string `json:"echo"`
}

// Ping godoc
// @Summary ping
// @Description ping where you can specify timeout before response
// @Description built to use in testing async requests
// @Tags ping
// @Accept  json
// @Produce  json
// @Param body body PingBody true "input payload"
// @Success 200 {object} TPingResponse
// @Router /ping [get]
func Ping(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var pingBody PingBody
	json.Unmarshal(reqBody, &pingBody)

	ch1 := make(chan TPingResponse)
	go func() {
		time.Sleep(time.Duration(pingBody.Timeout) * time.Second)
		d := TPingResponse{Duration: pingBody.Timeout, Message: pingBody.Echo}
		ch1 <- d
	}()

	w.Header().Set("Content-Type", "application/json")
	select {
	case res := <-ch1:
		json.NewEncoder(w).Encode(res)
	case <-time.After(10 * time.Second):
		res := TPingResponse{Duration: 10, Message: "Max wait reached."}
		json.NewEncoder(w).Encode(res)

	}
}
