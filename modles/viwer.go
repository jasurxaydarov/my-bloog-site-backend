package modles

import "github.com/google/uuid"

// /
type Viewer struct {
	ViewerID uuid.UUID `json:"viewer_id"`
	FullName string    `json:"full_name"`
	Username string    `json:"username"`
	Gmail    string    `json:"gmail"`
	Password string    `json:"password"`
}

type ViewerReqReg struct {
	FullName string `json:"full_name"`
	Username string `json:"username"`
	Gmail    string `json:"gmail"`
	Password string `json:"password"`
	Otp      string `json:"otp"`
}

type GetViwersResp struct {
	Viwers Viewer
	Count  int
}

type CheackViwer struct {
	Gmail string `json:"gmail"`
}

type LoginViwer struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthResp struct{
	AccessToken  string	`json:"access_token"`
}