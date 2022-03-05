package model

type SrsBaseReq struct {
	Action   string `json:"action" validate:"required"`
	ClientID string `json:"client_id" validate:"required"`
	IP       string `json:"ip" validate:"required" validate:"required"`
	VHost    string `json:"vhost" validate:"required"`
	App      string `json:"app" validate:"required"`
}

type OnConnectReq struct {
	SrsBaseReq
	TcUrl   string `json:"tcUrl" validate:"required,url"`
	PageUrl string `json:"pageUrl" validate:"required,url"`
}

type OnCloseReq struct {
	SrsBaseReq
}

type OnPublishReq struct {
	SrsBaseReq
	Stream string `json:"stream" validate:"required"`
}

type OnUnpublishReq struct {
	SrsBaseReq
	Stream string `json:"stream" validate:"required"`
}

type OnPlayReq struct {
	SrsBaseReq
	Stream string `json:"stream" validate:"required"`
}

type OnStopReq struct {
	SrsBaseReq
	Stream string `json:"stream" validate:"required"`
}

var PermGrantedRes = []byte{0}

var PermRejectedRes = []byte{1}
