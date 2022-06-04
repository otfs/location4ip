package main

import (
	"encoding/json"
	"location4ip/location4ip"
	"log"
	"net"
	"net/http"
	"strings"
)

//
// GetIpLocationHandle 获取Ip位置信息
//
func GetIpLocationHandle(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	ip := strings.TrimSpace(query.Get("ip"))
	if ip == "" {
		errorResult := Result{Code: "400", Msg: "required parameter ip"}
		writeResponse(w, http.StatusBadRequest, errorResult)
		return
	}
	ipInfo := net.ParseIP(ip)
	if ipInfo == nil {
		errorResult := Result{Code: "400", Msg: "invalid ip address"}
		writeResponse(w, http.StatusBadRequest, errorResult)
		return
	}

	location, err := location4ip.GetIpLocation(ip)
	if err != nil {
		errorResult := Result{Code: "500", Msg: "get ip location error"}
		writeResponse(w, http.StatusInternalServerError, errorResult)
		return
	}

	result := Result{Code: "0", Msg: "OK", Data: location}
	writeResponse(w, http.StatusOK, result)
}

func writeResponse(w http.ResponseWriter, httpStatus int, result interface{}) {
	w.WriteHeader(httpStatus)
	w.Header().Set("Content-Type", "application/json")

	resultJson, err := json.Marshal(result)
	if err != nil {
		log.Printf("write response error: %v", err)
	}
	_, err = w.Write(resultJson)
	if err != nil {
		log.Printf("write response error: %v", err)
	}
}

type Result struct {
	Code string      `json:"code"` // error code
	Msg  string      `json:"msg"`  // error message
	Data interface{} `json:"data"` // data
}
