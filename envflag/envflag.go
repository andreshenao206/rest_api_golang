package envflag

import (
	"github.com/namsral/flag"
)

var EmailAddress string

func Init() {
	flag.StringVar(&EmailAddress, "email_address", "", "Email User")
	flag.Parse()

}
