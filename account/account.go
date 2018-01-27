package account

import (
	"bufio"
	"log"
	"os"
	"path"
	"strings"
)

const line_separator = " "

type TotpAccount struct {
	Name   string
	Secret string
}

func TotpAccounts() []TotpAccount {
	accounts := make([]TotpAccount, 0)
	secretsfile, err := os.Open(path.Join(os.Getenv("HOME"), "totpsecrets"))
	if err != nil {
		log.Fatalf("Could not open file. Cause: %v", err)
	}
	scanner := bufio.NewScanner(secretsfile)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), line_separator)
		if len(line) != 2 {
			log.Fatalf("Expected two values per line separated by '%v'. Instead found %v values.", line_separator, len(line))
		}
		accounts = append(accounts, TotpAccount{line[0], line[1]})
	}
	secretsfile.Close()
	return accounts
}
