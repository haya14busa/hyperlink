package hyperlink

import (
	"fmt"
	"io"
	"os"
)

const (
	ESC  = "\033"
	OSC  = ESC + "]"
	ST   = "\033\\"
	BELL = "\a"
)

func Write(w io.Writer, link, text string) error {
	ew := &errWriter{w: w}
	ew.write(wrapOSC(osc8Start(link)))
	ew.write(text)
	ew.write(wrapOSC(osc8End()))
	return ew.err
}

func wrapOSC(esc string) string {
	if os.Getenv("TMUX") != "" {
		return wrapOSCForTmux(esc)
	}
	return esc
}

func wrapOSCForTmux(esc string) string {
	return ESC + "Ptmux;" + ESC + esc + ST
}

func osc8Start(link string) string {
	return OSC + "8" + ";;" + link + BELL
}

func osc8End() string {
	return OSC + "8" + ";;" + BELL
}

// REF: https://blog.golang.org/errors-are-values
type errWriter struct {
	w   io.Writer
	err error
}

func (ew *errWriter) write(str string) {
	if ew.err != nil {
		return
	}
	_, ew.err = fmt.Fprint(ew.w, str)
}
