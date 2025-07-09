package tui

import (
	"bufio"
	"errors"
	"hearthstone/internal/session"
	"hearthstone/pkg/ui"
	"os"
	"strings"
)

func Display(session *session.Session) {
	ui.UpdateFrame(sessionString(session))
}

func Input() ([]string, error) {
	if !scanner.Scan() {
		return nil, errors.New("End of input")
	}
	input := scanner.Text()

	input = strings.ToLower(input)
	return strings.Fields(input), nil
}

var scanner = bufio.NewScanner(os.Stdin)