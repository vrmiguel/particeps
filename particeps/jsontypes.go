package particeps

// UniversalResponse is the struct that all uploads return
type UniversalResponse struct {
	Status   bool
	FullURL  string
	ShortURL string
}

// AnonFilesSuccess matches the JSON successful response given by AnonFiles
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

// AnonFilesFailure matches the JSON failure response given by AnonFiles
type AnonFilesFailure struct {
	Status bool `json:"status"`
	Error  struct {
		Message string `json:"message"`
		Type    string `json:"type"`
		Code    int    `json:"code"`
	} `json:"error"`
}
