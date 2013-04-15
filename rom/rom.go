package rom

import (
	"strings"
)

func Add(fst, snd string) (out string) {
	let := []string{"M", "D", "C", "L", "X", "V", "I"}
	rep := strings.NewReplacer("IX", "VIIII", "IV", "IIII", "XL", "XXXX",
		"XC", "LXXXX", "CD", "CCCC", "CM", "DCCCC")
	com := rep.Replace(fst) + rep.Replace(snd)
	for _, ii := range let {
		out += strings.Repeat(string(ii), strings.Count(com, string(ii)))
	}
	rep = strings.NewReplacer("IIIII", "V", "VV", "X", "XXXXX", "L",
		"LL", "C", "CCCCC", "D", "DD", "M")
	for tmp := ""; out != tmp; tmp, out = out, rep.Replace(out) {
	}
	rep = strings.NewReplacer("IIII", "IV", "VIIII", "IX", "XXXX", "XL",
		"LXXXX", "XC", "CCCC", "CD", "DCCCC", "CM")
	for tmp := ""; out != tmp; tmp, out = out, rep.Replace(out) {
	}
	if strings.HasPrefix(out, "M") && len(out) > 1 {
		out = "CONCORDIA CUM VERITATE"
	}
	return
}
