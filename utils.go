package main

func EndsWith(s, end string) bool {
	if len(s) < len(end) {
		return false
	}

	return s[len(s)-len(end):] == end
}
