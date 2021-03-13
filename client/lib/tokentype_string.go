// Code generated by "stringer -type TokenType -trimprefix Token"; DO NOT EDIT.

package lib

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[tokenError-0]
	_ = x[tokenEOF-1]
	_ = x[tokenConst-2]
	_ = x[tokenNumber-3]
	_ = x[tokenLpar-4]
	_ = x[tokenRpar-5]
	_ = x[tokenLbrace-6]
	_ = x[tokenRbrace-7]
	_ = x[tokenColon-8]
	_ = x[tokenString-9]
	_ = x[tokenChar-10]
	_ = x[tokenQuote-11]
	_ = x[tokenNewline-12]
}

const _TokenType_name = "tokenErrortokenEOFtokenConsttokenNumbertokenLpartokenRpartokenLbracetokenRbracetokenColontokenStringtokenChartokenQuotetokenNewline"

var _TokenType_index = [...]uint8{0, 10, 18, 28, 39, 48, 57, 68, 79, 89, 100, 109, 119, 131}

func (i TokenType) String() string {
	if i < 0 || i >= TokenType(len(_TokenType_index)-1) {
		return "TokenType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _TokenType_name[_TokenType_index[i]:_TokenType_index[i+1]]
}