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
	Success   bool
	ErrorCode string
	Message   string
	Content   interface{}
}

func (response *Response) PrintJSON(writer io.Writer) {
	fmt.Fprintln(writer, response.ToJSON())
}

func (response *Response) ToJSON() string {
	json, err := json.Marshal(response)
	if err != nil {
		return ""
	}
	return string(json)
}
