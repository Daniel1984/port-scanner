package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/port-scanner/pkg/portscanner"
	"github.com/port-scanner/pkg/reqvalidator"
)

type openPorts struct {
	FromPort  int    `json:"from_port"`
	ToPort    int    `json:"to_port"`
	Domain    string `json:"domain"`
	OpenPorts []int  `json:"open_ports"`
}

func ScanOpenPorts(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	defer r.Body.Close()
	w.Header().Add("Content-Type", "application/json")

	queryValues := r.URL.Query()
	domain := queryValues.Get("domain")
	toPort := queryValues.Get("toPort")

	v := reqvalidator.New()
	v.Required("domain", domain)
	v.Required("toPort", toPort)
	v.ValidDecimalString("toPort", toPort)
	if !v.Valid() {
		w.WriteHeader(http.StatusForbidden)
		w.Write(v.GetErrResp())
		return
	}

	// safe to skip error check here as validator above has done that already
	port, _ := strconv.Atoi(toPort)
	op := portscanner.
		New(domain).
		ScanTo(port)

	report := openPorts{
		FromPort:  0,
		ToPort:    port,
		Domain:    domain,
		OpenPorts: op,
	}

	resp, err := json.Marshal(report)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
