package app

import (
	"fmt"

	database "github.com/RB-PRO/BazarakiUpdate/pkg/DataBase"
)

func Start() {
	bd, ErrorBD := database.New()
	if ErrorBD != nil {
		panic(ErrorBD)
	}
	fmt.Println(bd)
}
