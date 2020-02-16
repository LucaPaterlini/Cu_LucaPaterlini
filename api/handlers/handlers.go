package handlers

import (
	"Cu_LucaPaterlini/api/exchange"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

const tickerMissing = "ticker missing"

// HandlerExchange manage the endpoint that provide the latest currency exchange value and an advice.
func HandlerExchange(w http.ResponseWriter, r *http.Request) {
	// checking the currency ticker
	vars := mux.Vars(r)
	if _, ok := vars["ticker"]; !ok {
		http.Error(w, tickerMissing, http.StatusBadRequest)
		return
	}
	// preparing and adjusting the dates
	endDate := time.Now()
	startDate := endDate.AddDate(0, 0, -7)
	switch endDate.Weekday() {
	case 6:
		endDate =endDate.AddDate(0, 0, -1)
	case 0:
		endDate =endDate.AddDate(0, 0, -2)
	}

	resp, err := exchange.GetRateByDate(vars["ticker"], startDate.Format("2006-01-02"), endDate.Format("2006-01-02"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusFailedDependency)
		return
	}
	var respB []byte
	if respB,err = json.Marshal( &resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if _, err = w.Write(respB); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// positive response
	w.Header().Set("Content-Type", "application/json")
}
