package loop

import (
	sessionpkg "hearthstone/internal/session"
	"hearthstone/internal/setup"
	"strings"
)

func handleInput(session *sessionpkg.Session) (exit bool) {
	allArgs, err := setup.Input()
	if err != nil {
		return true
	}

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