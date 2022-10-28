package response

type APIBody struct {
	Data    string `json:"data,omitempty"`
	Status  bool   `json:"status"`
	Msg     string `json:"msg"`
	Path    string `json:"path"`
	Service string `json:"service"`
}

func APIResponse(data, msg, path string, status bool) *APIBody {
	return &APIBody{
		Data:   data,
		Msg:    msg,
		Status: status,
		Path:   path,
	}
}
