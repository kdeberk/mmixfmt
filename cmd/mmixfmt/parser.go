package main

import (
	"strings"
	"unicode"
)

type lineFragments struct {
	cols    []string
	comment string
}

func readFragments(str string) lineFragments {
	upperLettersOnly := func(s string) bool {
		if strings.ToUpper(s) != s {
			return false
		}

		for _, chr := range s {
			if !unicode.IsLetter(chr) {
				return false
			}
		}
		return true
	}

	str = strings.TrimRight(str, " ")

	var seenop bool
	var frags []string
	var comment string
L:
	for idx := 0; ; idx++ {
		if "" == str {
			break L
		}

		str = strings.TrimLeft(str, " ")
		pos := strings.Index(str, " ")
		switch str[0] {
		case '"':
			// Seek to next ", and seek to next space.
			qpos := 1 + strings.Index(str[1:], `"`)
			spos := strings.Index(str[qpos:], " ")
			if -1 == spos {
				pos = len(str)
			} else {
				pos = qpos + spos
			}
		case '%':
			comment = str
			break L
		}

		if -1 == pos {
			pos = len(str)
		}

		frag := str[:pos]
		frags, str = append(frags, frag), str[pos:]

		switch {
		case upperLettersOnly(frag):
			seenop = true
		case seenop && 0 < len(str):
			comment = strings.TrimLeft(str, " ")
			break L
		}
	}
	return lineFragments{cols: frags, comment: comment}
}

type line struct {
	label   string
	op      string
	args    string
	comment string
}

func readLine(str string) (l line) {
	str = strings.Trim(str, " ")
	if 0 == len(str) {
		return
	}

	frags := readFragments(str)

	switch len(frags.cols) {
	case 2:
		l.op, l.args = frags.cols[0], frags.cols[1]
	case 3:
		l.label, l.op, l.args = frags.cols[0], frags.cols[1], frags.cols[2]
	}
	l.comment = frags.comment
	return
}
