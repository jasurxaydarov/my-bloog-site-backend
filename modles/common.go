package modles

type GetList struct {
	Limit  int32
	Pge    int32
	Search string
}

type Common struct {
	TableName  string `json:"tanle_name"`
	ColumnName string `json:"clomn_name"`
	ExpValue   any    `json:"exp_value"`
}

type CheckExists struct {
	IsExists bool
	Status   string
}

type OtpData struct {
	Otp   string `json:"otp"`
	Gmail string `json:"gmail"`
}

type CheckOtpRep struct {
	IsRight bool `json:"is_right"`
}
