package auth

type State struct {
	PostToken     string `json:"post_token"`
	CompletionUrl string `json:"completion_url"`
}

type UserData struct {
	PostToken        string `json:"post_token"`
	CompletionUrl    string `json:"completion_url"`
	DivarAccessToken string `json:"divar_access_token"`
	PhoneNumber      string `json:"phone_number"`
}

type InitResponse struct {
	URL string `json:"url"`
}

type CallbackResponse struct {
	AccessToken string `json:"access_token"`
}
