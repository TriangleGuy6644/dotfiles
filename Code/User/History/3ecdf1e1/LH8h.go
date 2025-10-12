package middleware

import(
	"errors"
	"net/http"

	"goapi/api"
	"goapi/internal/tools"
	log "github.com/sirupsen/logrus"
)

var UnAuthorizationError = errors.New("Invalid username or token.")

func Authorization(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var username string = r.URL.Query().Get("username")
		var token = r.Header.Get("Authorization")
		var err error

		if username == "" || token == "" {
			log.Error(UnAuthorizationError)
			api.RequestErrorHandler(w, UnAuthorizationError)
			return
	}

	var database *tools.DatabaseInterface
	database, err = tools.NewDatabase()
	if err != nil {
		api.InternalErrorHandler(w)
		return
	}
}