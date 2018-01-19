package main

import (
	"github.com/pquerna/otp/totp"
	"fmt"
	"strings"
	"time"
)

type TotpAccount struct {
	name string
	secret string
}

func main() {
	fmt.Println("totp client starting...")
	accounts := []TotpAccount{
		{"Good account", "abcde"},
		{"Bad account", "💩"},
	}
	generateCodes(accounts)
}

func generateCodes(accounts []TotpAccount) {
	codeTime := time.Now() // TODO use parameter to determine time
	for _, account := range accounts {
		code, err := totp.GenerateCode(strings.ToUpper(account.secret), codeTime)
	    if err == nil {
			fmt.Printf("Code for [%v] is [%v]\n", account.name, code)
	    } else {
	        fmt.Printf("Could not generate code for [%v] due to [%v]\n", account.name, err)
	    }
	}
}
