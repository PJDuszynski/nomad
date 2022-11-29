package oidc

import (
	"testing"

	"github.com/shoenig/test/must"
)

func TestCallbackServer(t *testing.T) {

	testCallbackServer, err := NewCallbackServer()
	must.NoError(t, err)
	must.NotNil(t, testCallbackServer)

	defer func() {
		must.NoError(t, testCallbackServer.Close())
	}()
	must.StrNotEqFold(t, "", testCallbackServer.Nonce())
	must.StrNotEqFold(t, "", testCallbackServer.RedirectURI())
}
