package api

// ResponseData defines the structure of the API response
type ResponseData struct {
	Data struct {
		Result struct {
			Results []interface{} `json:"results"`
		} `json:"result"`
		Error interface{} `json:"error"`
	} `json:"data"`
	ResponseDetail string `json:"responseDetail"`
	ResponseCode   int    `json:"responseCode"`
}
