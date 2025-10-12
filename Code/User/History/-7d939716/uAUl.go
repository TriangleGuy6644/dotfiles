package api

import(
	"net/http"
	"encoding/json"
)

//coin balance params
type CoinBalanceParams struct {
	Username string
}

type CoinBalanceResponse struct {
	//success code usually 200
	Code int
	Balance int64
}

type Error struct{
	Code int
	Message string
}

func writeError(w http.ResponseWriter, message string, code int){
	resp := Error{
		Code: code,
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json")
	
}