package banner

import (
	"strings"
)

func Inline(input string) string {
	input = strings.TrimSpace(input)
	if len(input) == 0 {
		return ""
	}

	lines := small.Get(rune(input[0])).lines()
	height := len(lines)
	if len(input) > 1 {
		for _, r := range input[1:] {
			switch r {
			case ' ':
				for i := 0; i < height; i++ {
					lines[i] += "  "
				}
			default:
				letter := small.Get(r).lines()
				for i := 0; i < height; i++ {
					lines[i] += letter[i]
				}
			}
		}
	}

	for i := 0; i < height; i++ {
		lines[i] = strings.TrimRight(lines[i], " ")
	}
	if lines[height-1] == "" {
		lines = lines[:height-1]
	}
	if lines[0] == "" {
		lines = lines[1:]
	}
	return strings.Join(lines, "\n")
}

type font map[rune]letter

func (f font) Get(key rune) letter {
	letter, found := f[key]
	if found {
		return letter
	}
	return f['?']
}

type letter string

func (l letter) String() string {
	return strings.Join(l.lines(), "\n")
}

func (l letter) lines() []string {
	trim := string(l[1 : len(l)-2])
	return strings.Split(trim, "@\n")
}
