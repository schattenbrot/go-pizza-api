package responder

import (
	"encoding/json"
	"net/http"
)

// Send is the helper function for sending back an HTTP response.
func Send(w http.ResponseWriter, status int, data ...interface{}) error {
	w.WriteHeader(status)

	if len(data) == 1 {
		w.Header().Set("Content-Type", "application/json")
		js, err := json.Marshal(data)
		if err != nil {
			return err
		}
		w.Write(js)
		return nil
	}

	w.Write([]byte{})
	return nil
}

// Send is the helper function for sending back an HTTP response.
func SendFile(w http.ResponseWriter, file []byte) {
	fileHeader := file[:512]
	fileContentType := http.DetectContentType(fileHeader)

	w.Header().Set("Content-Type", fileContentType)
	w.WriteHeader(http.StatusOK)
	w.Write(file)
}
