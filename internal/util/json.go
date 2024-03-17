package util

import "encoding/json"

func WriteJSON(data any) []byte {
	response, err := json.Marshal(data)
	if err != nil {
		return nil
	}
	
	return response
}