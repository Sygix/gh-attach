package app

import "strings"

func canAttemptTokenUpload(token string) bool {
	// Unrecognized formats may be legacy PATs; let the server decide.
	return token != "" && !strings.HasPrefix(token, "ghs_") && !strings.HasPrefix(token, "ghr_")
}
