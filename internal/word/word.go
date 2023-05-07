package word

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"strings"
	"unicode"
)

// 具体的5种单词模式转换逻辑
// 整个字符串的字符转大写

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

// 整个字符串的字符转小写

func ToLower(s string) string {
	return strings.ToLower(s)
}

// 下划线转大写驼峰

func UnderscoreToUpperCamelCase(s string) string {
	s = strings.Replace(s, "_", " ", -1)
	s = cases.Title(language.English).String(s)
	return strings.Replace(s, " ", "", -1)
}

// 下划线转小写驼峰

func UnderscoreToLowerCamelCase(s string) string {
	s = UnderscoreToUpperCamelCase(s)
	return string(unicode.ToLower(rune(s[0]))) + s[1:]
}

// 驼峰转下划线

func CamelCaseToUnderscore(s string) string {
	var output []rune
	for i, ch := range s {
		if i == 0 {
			output = append(output, unicode.ToLower(ch))
			continue
		}
		if unicode.IsUpper(ch) {
			output = append(output, '_')
		}
		output = append(output, unicode.ToLower(ch))
	}
	return string(output)
}
