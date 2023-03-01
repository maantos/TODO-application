package http

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func respond(rw http.ResponseWriter, r *http.Request, resp *Response) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(resp.Code)
	b, err := json.Marshal(resp)
	if err != nil {
		fmt.Println(fmt.Errorf("marshalling response to json failed, err: %v", err))
		return
	}
	_, err = rw.Write(b)
	if err != nil {
		fmt.Println(fmt.Errorf("writing response to response-writer failed, err: %v", err))
	}
}
