package config

import "regexp"

var (
	RegexpSha512           = regexp.MustCompile(`^[a-fA-F0-9]{128}$`)
	RegexpHasLowercase     = regexp.MustCompile(`[a-z]`)
	RegexpHasUppercase     = regexp.MustCompile(`[A-Z]`)
	RegexpDigitsOnly       = regexp.MustCompile(`^\d+$`)
	RegexpHasDigit         = regexp.MustCompile(`\d`)
	RegexpNonNegativeInt   = regexp.MustCompile(`^(0|[1-9]\d{0,8})$`)
	RegexpUsernameChars    = regexp.MustCompile(`^[a-zA-Z0-9_.@-]+$`)
	RegexpHasSpecialChar   = regexp.MustCompile(`[@$!%*?&_\-=]`)
)
