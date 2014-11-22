/**
 * This is the general response that will
 * encompass all other responses.
 */
package models

import (
	"encoding/json"
	"fmt"
	"io"
)

type Response struct {
	Success   bool        `json:"success"`
	ErrorCode string      `json:"error_code"`
	Message   string      `json:"message"`
	Content   interface{} `json:"content"`
}

func (response *Response) PrintJSON(writer io.Writer) {
	fmt.Fprintln(writer, response.ToJSON())
}

func (response *Response) ToJSON() string {
	json, err := json.Marshal(response)
	if err != nil {
		return "Error converting response to JSON"
	}
	return string(json)
}
