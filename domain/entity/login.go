package entity

type SmsReq struct {
	Mobile string `json:"mobile"`
	Valid  bool   `json:"valid"`
}

type SmsResult struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Vcode   string `json:"vcode"`
}

type LoginReq struct {
	Code     string    `json:"code"`
	UserInfo LoginUser `json:"userInfo"`
}

type LoginUser struct {
	Mobile string `json:"mobile"`
	Valid  bool   `json:"valid"`
}

type UserInfo struct {
	Avatar        string `json:"avatar"`
	Birthday      int    `json:"birthday"`
	Gender        int    `json:"gender"`
	Id            int    `json:"id"`
	LastLoginIp   string `json:"last_login_ip"`
	LastLoginTime int64  `json:"last_login_time"`
	Mobile        string `json:"mobile"`
	Nickname      string `json:"nickname"`
	Password      string `json:"password"`
	RegisterIp    string `json:"register_ip"`
	RegisterTime  int64  `json:"register_time"`
	Username      string `json:"username"`
}
