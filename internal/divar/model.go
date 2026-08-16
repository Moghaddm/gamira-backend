package divar

type Session struct {
	PostToken     string `json:"post_token"`
	CompletionUrl string `json:"completion_url"`
}

type StartFlowResponse struct {
	URL string `json:"url"`
}
