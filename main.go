package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/epond/totp-client/account"
	"github.com/pquerna/otp/totp"
)

func main() {
	generateCodes(account.TotpAccounts())
}

func generateCodes(accounts []account.TotpAccount) {
	codeTime := time.Now()
	fmt.Printf("It is %v\n", codeTime.Format("Jan 2, 2006 at 3:04pm (MST)"))
	for _, account := range accounts {
		code, err := totp.GenerateCode(strings.ToUpper(account.Secret), codeTime)
		if err == nil {
			fmt.Printf("%v - %v\n", account.Name, code)
		} else {
			fmt.Printf("Could not generate code for %v due to [%v]\n", account.Name, err)
		}
	}
}
