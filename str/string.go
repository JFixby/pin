package str

import (
	"github.com/jfixby/pin/lang"
	"strings"
)

func PreFillString(size int, char string) string {
	s := ""
	for i := 0; i < size; i++ {
		s = s + char
	}
	return s
}

func IndexOf(str, substr string, offsetFrom int) int {
	if len(substr) > len(str) {
		return -1
	}
	N := len(str) - len(substr)
	for i := offsetFrom; i <= N; i++ {
		if match(str, substr, i) {
			return i
		}
	}
	return -1
}

func ReplaceLines(original, replacement []string, fromIndex, toIndex int) []string {
	if fromIndex == -1 {
		lang.ReportErr("fromIndex is -1, data:\n" + strings.Join(original, "\n"))
	}
	if toIndex < fromIndex {
		lang.ReportErr("toIndex(%v) is below fromIndex(%v)", toIndex, fromIndex)
	}

	result := []string{}
	i := 0
	k := 0
	for {
		if i >= len(original) {
			break
		}
		if i > toIndex || i < fromIndex {
			result = append(result, original[i])
			i++
			continue
		}
		result = append(result, replacement[k])
		k++
		if k >= len(replacement) {
			i = toIndex + 1
		}
	}
	return result
}

func match(str string, substr string, offset int) bool {
	k := offset
	for i := 0; i < len(substr); i++ {
		a := str[k : k+1]
		b := substr[i : i+1]
		if a != b {
			return false
		}
		k++
	}
	return true
}

func DeleteLine(data string, prefix string, fromIndex int) string {
	result := []string{}
	lines := strings.Split(data, "\n")
	searching := true
	for i := fromIndex; i < len(lines); i++ {
		e := lines[i]
		if strings.Index(e, prefix) == 0 && searching {
			searching = false
			continue
		}
		result = append(result, e)
	}
	if searching {
		lang.ReportErr("Prefix not found: <%v>", prefix)
	}
	return strings.Join(result, "\n")
}

func ReplaceLine(data string, prefix string, replacement string, fromIndex int) string {
	lines := strings.Split(data, "\n")
	i := IndexOfLine(lines, prefix, fromIndex)
	if i == -1 {
		lang.ReportErr("Prefix not found <%v>\n",
			prefix,
			data,
		)
	}
	lines[i] = replacement
	return strings.Join(lines, "\n")
}

func EndsWith(str string, postfix string) bool {
	fileExt := str[len(str)-len(postfix):]
	return fileExt == postfix
}

func InsertLineAt(index int, lines []string, line string) []string {
	i := index
	if i < 0 || i > len(lines) {
		lang.ReportErr("Invalid index <%v>\n",
			i,
		)
	}
	result := []string{}
	for i, s := range lines {
		result = append(result, s)
		if i == index {
			result = append(result, line)
		}
	}
	return result
}

func IndexOfLine(lines []string, prefix string, fromIndex int) int {
	for i := fromIndex; i < len(lines); i++ {
		s := lines[i]
		if strings.Index(s, prefix) == 0 {
			return i
		}
	}
	return -1
}
