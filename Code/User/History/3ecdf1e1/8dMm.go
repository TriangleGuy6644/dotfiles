package middleware

import(
	"errors"
	"net/http"

	"goapi/api"
	"goapi/internal/tools"
	log "github.com/sirupsen/logrus"
)

var UnAuthorizationError = errors.New("Invalid username or token.")

func Authorization()