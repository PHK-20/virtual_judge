package remote_account

type RemoteAccount struct {
	Hdu struct {
		Accounts []struct {
			Username string `json:"username"`
			Password string `json:"password"`
		} `json:"accounts"`
	} `json:"hdu"`
	Hnust struct {
		Accounts []struct {
			Username string `json:"username"`
			Password string `json:"password"`
		} `json:"accounts"`
	} `json:"hnust"`
}
