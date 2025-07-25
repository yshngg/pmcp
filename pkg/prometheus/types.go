package prome

// {
// 	"status": "success" | "error",
// 	"data": <data>,

// 	// Only set if status is "error". The data field may still hold
// 	// additional data.
// 	"errorType": "<string>",
// 	"error": "<string>",

//  // Only set if there were warnings while executing the request.
//	// There will still be data in the data field.
//	"warnings": ["<string>"],
//	// Only set if there were info-level annotations while executing the request.
//	"infos": ["<string>"]
// }

type JSONResponseStatus string

const (
	JSONResponseStatusSuccess JSONResponseStatus = "success"
	JSONResponseStatusError   JSONResponseStatus = "error"
)

type JSONResponse struct {
	Status    JSONResponseStatus `json:"status"`
	Data      any                `json:"data"`
	ErrorType string             `json:"errorType"`
	Error     string             `json:"error"`
	Warnings  []string           `json:"warnings"`
	Infos     []string           `json:"infos"`
}
