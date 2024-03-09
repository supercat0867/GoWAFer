package api_handler

// LoginRequest 登录请求体
type LoginRequest struct {
	Username  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required"`
	CaptchaID string `json:"captchaId" binding:"required"`
	Captcha   string `json:"captcha" binding:"required"`
}

// CreateIPRequest 创建IP请求
type CreateIPRequest struct {
	IPAddress    string `json:"ip-address" binding:"required"`
	Type         int    `json:"type"`
	ExpirationAt string `json:"expirationAt"`
}

// UpdateIPRequest 更新IP请求
type UpdateIPRequest struct {
	IPAddress    string `json:"ip_address" binding:"required"`
	ExpirationAt string `json:"expiration_at"`
}
