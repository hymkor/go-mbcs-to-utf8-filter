package mbcs

import (
	"bufio"
	"io"
	"unicode/utf8"

	"golang.org/x/sys/windows"
)

const _ACP = 0

func ansiToUtf8(mbcs []byte) (string, error) {
	// query buffer's size
	size, _ := windows.MultiByteToWideChar(
		_ACP, 0, &mbcs[0], int32(len(mbcs)), nil, 0)
	if size <= 0 {
		return "", windows.GetLastError()
	}

	// convert ansi to utf16
	utf16 := make([]uint16, size)
	rc, _ := windows.MultiByteToWideChar(
		_ACP, 0, &mbcs[0], int32(len(mbcs)), &utf16[0], size)
	if rc == 0 {
		return "", windows.GetLastError()
	}
	// convert utf16 to utf8
	return windows.UTF16ToString(utf16), nil
}

type Filter struct {
	*bufio.Scanner
	text string
	ansi bool
	err  error
}

func NewFilter(r io.Reader) *Filter {
	return &Filter{
		Scanner: bufio.NewScanner(r),
	}
}

func (sc *Filter) Text() string {
	return sc.text
}

func (sc *Filter) Scan() bool {
	if !sc.Scanner.Scan() {
		sc.err = sc.Scanner.Err()
		return false
	}
	line := sc.Bytes()
	if !sc.ansi && utf8.Valid(line) {
		sc.text = sc.Scanner.Text()
	} else {
		sc.text, sc.err = ansiToUtf8(line)
		if sc.err != nil {
			return false
		}
		sc.ansi = true
	}
	return true
}
