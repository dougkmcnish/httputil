package response

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//ResponseBody is a simple struct suitable for
//use as a response to a HTTP request.
type ResponseBody struct {
	Message []string
	Errors  uint
	Failed  bool
}

//NewResponseBody creates a new ResponseBody
func NewResponseBody() ResponseBody {
	return ResponseBody{}
}

//Append pushes a message onto a []string of response
//messages.
func (r *ResponseBody) Append(message string) {
	r.Message = append(r.Message, message)
}

//Error increments ResponseBody.Errors to indicate
//something went wrong. Whether the error was fatal
//is up to the client implementation.
func (r *ResponseBody) Error(message string) {
	r.Errors++
	r.Append(message)
}

//Fatal sets ResponseBody.Failed to true and
//calls ResponseBody.Error() to set errored state.
//Fatal should always indicate an unrecoverable
//failure.
func (r *ResponseBody) Fatal(message string) {
	r.Failed = true
	r.Error(message)
}

//Send marshalls the ResponseBody into JSON and
//sends the response back to the client.
//It takes a http.ResponseWriter and a HTTP status code as arguments.
func (r *ResponseBody) Send(w http.ResponseWriter, status int) {
	body, _ := json.Marshal(r)
	w.WriteHeader(status)
	fmt.Fprint(w, string(body))
}
