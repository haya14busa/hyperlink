// Pacakge hyperlink is utility of hyperlink (OSC 8) feature.
package hyperlink

import (
	"fmt"
	"io"
	"os"
)

const (
	esc  = "\033"
	osc  = esc + "]"
	st   = esc + "\\"
	bell = "\a"
)

// Write writes OSC 8 sequence (hyperlink).
func Write(w io.Writer, link, text string) error {
	ew := &errWriter{w: w}
	ew.write(wrapOSC(osc8Start(link)))
	ew.write(text)
	ew.write(wrapOSC(osc8End()))
	return ew.err
}

func wrapOSC(s string) string {
	if os.Getenv("TMUX") != "" {
		return wrapOSCForTmux(s)
	}
	return s
}

func wrapOSCForTmux(s string) string {
	return esc + "Ptmux;" + esc + s + st
}

func osc8Start(link string) string {
	return osc + "8" + ";;" + link + bell
}

func osc8End() string {
	return osc + "8" + ";;" + bell
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
