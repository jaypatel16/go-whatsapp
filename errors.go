package whatsapp

import (
	"fmt"
	"github.com/pkg/errors"
)

var (
	ErrAlreadyConnected  = errors.New("already connected")
	ErrAlreadyLoggedIn   = errors.New("already logged in")
	ErrInvalidSession    = errors.New("invalid session")
	ErrLoginInProgress   = errors.New("login or restore already running")
	ErrNotConnected      = errors.New("not connected")
	ErrInvalidWsData     = errors.New("received invalid data")
	ErrConnectionTimeout = errors.New("connection timed out")
	ErrMissingMessageTag = errors.New("no messageTag specified or to short")
	ErrInvalidHmac       = errors.New("invalid hmac")
	ErrConnectionFailed  = errors.New("connection to WhatsApp servers failed")
)

type ErrConnectionClosed struct {
	Code int
	Text string
}

func (e *ErrConnectionClosed) Error() string {
	return fmt.Sprintf("server closed connection,code: %d,text: %s", e.Code, e.Text)
}
