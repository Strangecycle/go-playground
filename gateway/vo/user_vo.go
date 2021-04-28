package vo

type UserLoginRequestVO struct {
	Phone   string `json:"phone" binding:"required,max=11"`
	Captcha string `json:"captcha" binding:"required,max=6"`
}
