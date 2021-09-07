package gohelper

type requestUser struct {
	Id          string `json:"id,omitempty"`
	City        string `json:"city,omitempty"`
	Email       string `json:"email,omitempty"`
	PhoneNumber string `json:"phoneNumber,omitempty"`
	Roles       string `json:"roles,omitempty"`
	AuthTime    int64  `json:"authTime,omitempty"`
	Exp         int64  `json:"exp,omitempty"`
	Iat         int64  `json:"iat,omitempty"`
	Sub         string `json:"sub,omitempty"`
	Aud         string `json:"aud,omitempty"`
	Iss         string `json:"iss,omitempty"`
}
