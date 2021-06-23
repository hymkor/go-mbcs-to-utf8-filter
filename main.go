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
	sc   *bufio.Scanner
	text string
	ansi bool
	err  error
}

func NewFilter(r io.Reader) *Filter {
	return &Filter{
		sc: bufio.NewScanner(r),
	}
}

func (f *Filter) Scan() bool {
	if !f.sc.Scan() {
		f.err = f.sc.Err()
		return false
	}
	line := f.sc.Bytes()
	if !f.ansi && utf8.Valid(line) {
		f.text = f.sc.Text()
	} else {
		f.text, f.err = ansiToUtf8(line)
		if f.err != nil {
			return false
		}
		f.ansi = true
	}
	return true
}
func (f *Filter) Text() string {
	return f.text
}

func (f *Filter) Err() error {
	return f.err
}
