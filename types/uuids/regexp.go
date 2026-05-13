package uuids

import "regexp"

// RegexpUUID matches a standard UUID in the form
// xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx (case-insensitive, any version).
var RegexpUUID = regexp.MustCompile(
	`^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$`,
)
