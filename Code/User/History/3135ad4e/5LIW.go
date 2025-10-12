package tools

import(
	log "github.com/sirupsen/logrus"
)

type LoginDetails struct{
	AuthToken string
	Username string
}
type CoinDetails struct {
	Coins int
}