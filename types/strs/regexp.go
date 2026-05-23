package strs

import "regexp"

var (
	RegexpSha512         = regexp.MustCompile(`^[a-fA-F0-9]{128}$`)
	RegexpHasLowercase   = regexp.MustCompile(`[a-z]`)
	RegexpHasUppercase   = regexp.MustCompile(`[A-Z]`)
	RegexpDigitsOnly     = regexp.MustCompile(`^\d+$`)
	RegexpHasDigit       = regexp.MustCompile(`\d`)
	RegexpNonNegativeInt = regexp.MustCompile(`^(0|[1-9]\d{0,8})$`)
	RegexpUsernameChars  = regexp.MustCompile(`^[a-zA-Z0-9_.@-]+$`)
	RegexpHasSpecialChar = regexp.MustCompile(`[@$!%*?&_\-=]`)
	RegexpAlpha          = regexp.MustCompile(`^[a-zA-Z]+$`)
	RegexpAlphaNumeric   = regexp.MustCompile(`^[a-zA-Z0-9]+$`)
	RegexpHTMLTag        = regexp.MustCompile(`<[^>]*>`)
	RegexpUUID           = regexp.MustCompile(`^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$`)
	RegexpHex            = regexp.MustCompile(`^[0-9a-fA-F]+$`)
	RegexpHexColor       = regexp.MustCompile(`^#([0-9a-fA-F]{3}|[0-9a-fA-F]{6}|[0-9a-fA-F]{8})$`)
	RegexpSha256         = regexp.MustCompile(`^[a-fA-F0-9]{64}$`)
	RegexpSlug           = regexp.MustCompile(`^[a-z0-9]+(?:-[a-z0-9]+)*$`)
	RegexpDataURI        = regexp.MustCompile(`^data:[a-zA-Z0-9!#$&^_.+-]+/[a-zA-Z0-9!#$&^_.+-]+(?:;[a-zA-Z0-9!#$&^_.+-]+=[a-zA-Z0-9!#$&^_.+-]+)*(?:;base64)?,[a-zA-Z0-9!$&',()*+;=\-._~:@/?%]*$`)
)
