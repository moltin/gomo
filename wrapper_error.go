package gomo

// APIError is an error returned by the API so that you can include error speciifc logic in your own implementation
//
// 	if error.Status == 404 {
// 		// create something
// 	}
type APIError struct {
	Status int    `json:"status"`
	Detail string `json:"detail"`
	Title  string `json:"title"`
}
