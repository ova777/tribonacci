package main

import (
	"./lib"
	"encoding/json"
	"github.com/go-martini/martini"
	"net/http"
	"strconv"
)

// Response represents a general response api format
type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data"`
}

// Params represents a common set of parameters for forming the response
type Params struct {
	data  interface{}
	error string
	code  int
}

// main starts rest server on port 3000
func main() {
	m := martini.Classic()
	m.Get("/tribonacci/:n", GetTribonacci)
	m.NotFound(func(w http.ResponseWriter) (int, string) {
		return FormatResponse(w, Params{error: "Not found response", code: 404})
	})
	m.Run()
}

// FormatResponse formats json response from parameters
func FormatResponse(w http.ResponseWriter, params Params) (int, string) {
	if params.code == 0 {
		params.code = 200
	}

	data := Response{Success: true}

	if params.error != "" {
		data.Message = params.error
		data.Success = false
	} else {
		data.Data = params.data
	}

	w.Header().Set("Content-Type", "application/json")

	jsonStr, _ := json.Marshal(data)
	return params.code, string(jsonStr)
}

// GetTribonacci handles enpoint calculation of the tribonacci number
func GetTribonacci(params martini.Params, w http.ResponseWriter) (int, string) {
	gotN := params["n"]
	n, err := strconv.ParseUint(gotN, 10, 0)

	if err != nil || n <= 0 {
		return FormatResponse(w, Params{error: `"` + gotN + `" is not a natural number`, code: 400})
	}

	return FormatResponse(w, Params{data: lib.Tribonacci(uint(n))})
}
