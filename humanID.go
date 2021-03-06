package humanID

import (
	// "errors"
	payload "github.com/bluenumberfoundation/humanid-golang-sdk/payload"
	httpHelper "github.com/bluenumberfoundation/humanid-golang-sdk/helpers/http"
)

const SERVER_HOST = "https://core.human-id.org/v0.0.3/server/"

type HumanID interface {
	Login(lang string, clientSecret string) (*payload.LoginResp, error)
	VerifyToken(exchangeToken string) (*payload.VerifyTokenResp, error)
}

type HumanIDImpl struct {
	serverClientID     string
	serverClientSecret string
}

func New(
	serverClientID string, 
	serverClientSecret string,
) HumanID {
	return &HumanIDImpl{
		serverClientID: serverClientID,
		serverClientSecret: serverClientSecret,
	}
}

func (h HumanIDImpl) Login(lang string, clientSecret string) (*payload.LoginResp, error) {
	loginResp := &payload.LoginResp{}
	err := httpHelper.PostNoBody(
		SERVER_HOST + "users/web-login?lang=" + lang + "&clientSecret=" + clientSecret,
		loginResp,
		h.serverClientID,
		h.serverClientSecret,
	)
	if err != nil {
		return nil, err
	}
	return loginResp, nil
}

func (h HumanIDImpl) VerifyToken(exchangeToken string) (*payload.VerifyTokenResp, error) {
	verifyTokenReq := &payload.VerifyTokenReq{
		ExchangeToken: exchangeToken,
	}
	verifyTokenResp := &payload.VerifyTokenResp{}

	err := httpHelper.Post(
		SERVER_HOST + "users/exchange",
		verifyTokenReq,
		verifyTokenResp,
		h.serverClientID,
		h.serverClientSecret,
	)
	if err != nil {
		return nil, err
	}
	return verifyTokenResp, nil
}