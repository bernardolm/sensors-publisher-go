package database

import (
	"time"

	_ "time/tzdata"
)

func init() {
	time.Local, _ = time.LoadLocation("America/Sao_Paulo")
}
