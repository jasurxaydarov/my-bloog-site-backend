package modles

type Owner struct {
	FullName    string `json:"full_name"`
	Password    string `json:"password"`
	Role        string `json:"role"`
	PhoneNumber string `json:"phone_number"`
	Gmail       string `json:"gmail"`
	LinkedIn    string `json:"linked_in"`
	Telegram    string `json:"telegram"`
	Github      string `json:"github"`
	Leetcode    string `json:"leetcode"`
	AboutMe     string `json:"about_me"`
}

type LoginOwn struct {
	PhoneNumber string `json:"phone_number"`
	Gmail       string `json:"gmail"`
	Password    string `json:"password"`
	Role        string `json:"role"`
}
