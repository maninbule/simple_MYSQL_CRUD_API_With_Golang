package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func ParseBody(r *http.Request, save_data interface{}) bool {
	if err := json.NewDecoder(r.Body).Decode(save_data); err != nil {
		log.Fatalf("r.Body 解析失败")
		return false
	}
	return true
	//if body, err := io.ReadAll(r.Body); err != nil {
	//	json.NewDecoder(body)
	//}
}
