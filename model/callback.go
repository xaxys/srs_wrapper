package model

type SrsCallbackReq struct {
	Action   string `json:"action"`
	ClientID string `json:"client_id"`
	IP       string `json:"ip"`
	VHost    string `json:"vhost"`
	App      string `json:"app"`
	Param    string `json:"param"`
	TcUrl    string `json:"tcUrl"`
	PageUrl  string `json:"pageUrl"`
	Stream   string `json:"stream"`
}

var PermGrantedRes = []byte("0")

var PermRejectedRes = []byte("1")
