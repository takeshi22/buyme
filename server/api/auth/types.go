package auth

type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AccessUuid   string
	RefreshUuid  string
	AtExpires    int64
	RtExpires    int64
}

type LoginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
