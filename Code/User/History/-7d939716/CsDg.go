package api

import(
	"net/http"
	"encoding/json"
)

//coin balance params
type CoinBalanceParams struct {
	Username string
}

