package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/epond/totp-client/account"
	"github.com/pquerna/otp/totp"
)

func main() {
	fmt.Println("totp client starting...")
	generateCodes(account.TotpAccounts())
}

func generateCodes(accounts []account.TotpAccount) {
	codeTime := time.Now()
	for _, account := range accounts {
		code, err := totp.GenerateCode(strings.ToUpper(account.Secret), codeTime)
		if err == nil {
			fmt.Printf("Code for [%v] is [%v]\n", account.Name, code)
		} else {
			fmt.Printf("Could not generate code for [%v] due to [%v]\n", account.Name, err)
		}
	}
}
