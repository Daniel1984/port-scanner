// cmd/server/handlers/scanports.go
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
	v.ValidStringInt("toPort", toPort)
	if !v.Valid() {
		w.WriteHeader(http.StatusForbidden)
		w.Write(v.GetErrResp())
		return
	}

	port, _ := strconv.Atoi(toPort)
	ps := portscanner.New(domain, 200)
	op := ps.ScanTo(port)
	resp := openPorts{
		FromPort:  0,
		ToPort:    port,
		Domain:    domain,
		OpenPorts: op,
	}

	b, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
