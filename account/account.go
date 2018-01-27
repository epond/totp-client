package account

type TotpAccount struct {
	Name   string
	Secret string
}

func TotpAccounts() []TotpAccount {
	return []TotpAccount{
		{"Good account", "abcde"},
		{"Bad account", "💩"},
	}
}
