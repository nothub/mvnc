package mvnc

import (
	"encoding/base64"
	"errors"
	"fmt"
)

var ErrAuthFmt = errors.New("invalid auth data")

var auth = ""

func AuthToken(token string) error {
	if token == "" {
		return ErrAuthFmt
	}

	auth = fmt.Sprintf("Authorization: Bearer %s", token)

	return nil
}

func AuthUser(user string, pass string) error {
	if user == "" || pass == "" {
		return ErrAuthFmt
	}

	raw := []byte(fmt.Sprintf("%s:%s", user, pass))
	b64 := base64.StdEncoding.EncodeToString(raw)
	auth = fmt.Sprintf("Authorization: Basic %s", b64)

	return nil
}
