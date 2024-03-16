package helper

import (
	"encoding/json"
	"net/http"
)


func ConvertToJson(w http.ResponseWriter, data interface{}) {
	// untuk memberitahu bahwa ini data bentuknya json
	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err := encoder.Encode(data)
	HandleErrorWithPanic(err)
}