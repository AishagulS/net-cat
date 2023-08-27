package net_cat

func NameCheck(r rune) bool {
	if (r < 'A' || r > 'Z') && (r < 'a' || r > 'z') && (r < '1' || r > '9') {
		return false
	}
	return true
}

func MsgCheck(msg string) bool {
	for _, ch := range msg {
		if ch < 32 || ch > 126 {
			return false
		}
	}
	return true
}
