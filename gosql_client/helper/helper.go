package helper

import (
	"gosql_client/constant"
	"strings"
)

func IsTokenEqualIgnoringCase(this string, that string) bool {
	this = strings.TrimSpace(this)
	this = strings.ToLower(this)

	that = strings.TrimSpace(that)
	that = strings.ToLower(that)

	return this == that
}

func IsKeyword(keyword string) bool {

	//Check valid keyword
	var isKeyword bool = false
	for _, checkKeyword := range constant.RESERVED_KEYWORD {
		if IsTokenEqualIgnoringCase(keyword, checkKeyword) {
			isKeyword = true
			break
		}
	}

	return isKeyword
}
