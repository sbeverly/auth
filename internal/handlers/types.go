package handlers

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Email string `json:"email"`
	Token string `json:"token"`
}

type VerifyRequest struct {
	Token string `json:"token"`
}

type SuccessResponse struct {
	Message string `json:"message,omitempty"`
}

type ErrorResponse struct {
	Error string `json:"error,omitempty"`
}

type PingResponse struct {
	Status string `json:"status, omitempty"`
	Uptime string `json:"uptime, omitempty"`
}
