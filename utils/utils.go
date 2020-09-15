package utils

import (
	"fmt"
	"strings"
)

func QuoteIdent(s, quote string) string {
	part := strings.SplitN(s, ".", 2)
	if len(part) == 2 {
		return QuoteIdent(part[0], quote) + "." + QuoteIdent(part[1], quote)
	}
	return quote + s + quote
}

func QuoteIdents(idents []string, quoteIdent func(s string) string) string {
	var b strings.Builder

	for i, s := range idents {
		b.WriteString(quoteIdent(s))

		if i < len(idents)-1 {
			b.WriteString(",")
		}
	}
	return b.String()
}

func QuoteInterfaces(values ...interface{}) string {
	var b strings.Builder

	for i, value := range values {

		b.WriteString("'")
		b.WriteString(Addslashes(fmt.Sprintf("%v", value)))
		b.WriteString("'")

		if i < len(values)-1 {
			b.WriteString(",")
		}
	}

	return b.String()
}

func QuoteStrings(values ...string) string {
	var b strings.Builder

	for i, value := range values {

		b.WriteString("'")
		b.WriteString(Addslashes(value))
		b.WriteString("'")

		if i < len(values)-1 {
			b.WriteString(",")
		}
	}

	return b.String()
}

// get the default value of the column
func DefaultValue(value interface{}) string {
	if value == nil {
		return ""
	} else if value, ok := value.(fmt.Stringer); ok {
		return value.String()
	}
	return fmt.Sprintf("'%v'", value)
}

// addslashes() 函数返回在预定义字符之前添加反斜杠的字符串。
// 预定义字符是：
// 单引号（'）
// 双引号（"）
// 反斜杠（\）
func Addslashes(str string) string {
	var tmpRune []rune
	strRune := []rune(str)
	for _, ch := range strRune {
		switch ch {
		case []rune{'\\'}[0], []rune{'"'}[0], []rune{'\''}[0]:
			tmpRune = append(tmpRune, []rune{'\\'}[0])
			tmpRune = append(tmpRune, ch)
		default:
			tmpRune = append(tmpRune, ch)
		}
	}
	return string(tmpRune)
}

// stripslashes() 函数删除由 addslashes() 函数添加的反斜杠。
func Stripslashes(str string) string {
	var dstRune []rune
	strRune := []rune(str)
	strLenth := len(strRune)
	for i := 0; i < strLenth; i++ {
		if strRune[i] == []rune{'\\'}[0] {
			i++
		}
		dstRune = append(dstRune, strRune[i])
	}
	return string(dstRune)
}

// unique string slice
func StringsUnique(items []string) []string {
	var set map[string]struct{}
	var newItems []string

	for _, item := range items {
		if _, ok := set[item]; ok == false {
			newItems = append(newItems, item)
		}
	}

	return newItems
}

type express struct {
	content string
}

func (e *express) String() string {
	return e.content
}

func Express(content string) fmt.Stringer {
	return &express{
		content: content,
	}
}
