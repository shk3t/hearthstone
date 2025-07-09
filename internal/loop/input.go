package loop

import (
	"bufio"
	sessionpkg "hearthstone/internal/session"
	"os"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

func handleInput(session *sessionpkg.Session) (exit bool) {
	var err error

	if !scanner.Scan() {
		return true
	}
	input := scanner.Text()

	input = strings.ToLower(input)
	allArgs := strings.Fields(input)

	var command string
	var args []string
	if len(allArgs) > 0 {
		command, args = allArgs[0], allArgs[1:]
	}

	session.Hint = ""

	err = Actions.ShortHelp.Do(args, session) // Display short help by default
	for _, action := range actionList {
		if strings.HasPrefix(command, action.shortcut) || command == action.name {
			err = action.Do(args, session)
			break
		}
	}

	if err != nil {
		session.Hint = err.Error()
	}
	return false
}