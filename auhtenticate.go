package cwauth

import (
	"github.com/coreos/go-oidc/v3/oidc"
)

func AuthURL() (string, error) {
	state, err := GenerateRandomString(16)
	if err != nil {
		return "", err
	}

	nonce, err := GenerateRandomString(16)
	if err != nil {
		return "", err
	}

	return Config.AuthCodeURL(state, oidc.Nonce(nonce)), nil
}
