package main

import (
	"github.com/NikosGour/mail_library/build"
	log "github.com/NikosGour/logging/log"
)

func main(){
	log.Debug("DEBUG_MODE = %t\n",build.DEBUG_MODE)
}