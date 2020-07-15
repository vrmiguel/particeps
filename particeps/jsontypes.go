package particeps

import "time"

// UniversalResponse is the struct that all uploads return
type UniversalResponse struct {
	Status   bool
	FullURL  string
	ShortURL string
}

// FilebinSuccess matches the successful JSON response given by Filebin
type FilebinSuccess struct {
	Filename string    `json:"filename"`
	Bin      string    `json:"bin"`
	Bytes    int       `json:"bytes"`
	Mime     string    `json:"mime"`
	Created  time.Time `json:"created"`
	Links    []struct {
		Rel  string `json:"rel"`
		Href string `json:"href"`
	} `json:"links"`
	Datetime time.Time `json:"datetime"`
}

// AnonFilesSuccess matches the successful JSON response given by AnonFiles
type AnonFilesSuccess struct {
	Status bool `json:"status"`
	Data   struct {
		File struct {
			URL struct {
				Full  string `json:"full"`
				Short string `json:"short"`
			} `json:"url"`
			Metadata struct {
				ID   string `json:"id"`
				Name string `json:"name"`
				Size struct {
					Bytes    int    `json:"bytes"`
					Readable string `json:"readable"`
				} `json:"size"`
			} `json:"metadata"`
		} `json:"file"`
	} `json:"data"`
}

// AnonFilesFailure matches the failure JSON response given by AnonFiles
type AnonFilesFailure struct {
	Status bool `json:"status"`
	Error  struct {
		Message string `json:"message"`
		Type    string `json:"type"`
		Code    int    `json:"code"`
	} `json:"error"`
}
