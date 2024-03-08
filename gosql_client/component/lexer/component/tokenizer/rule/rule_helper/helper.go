package rule_helper

import "strings"

func SanitizeTok(tok string) string {
	tok = strings.TrimSpace(tok)
	return tok
}
func AreTokEqual(tok1 string, tok2 string) bool {

	tok1 = SanitizeTok(tok1)
	tok2 = SanitizeTok(tok2)

	tok1 = strings.ToLower(tok1)
	tok2 = strings.ToLower(tok2)

	return tok1 == tok2
}
