package tui

import (
	"bufio"
	"context"
	"fmt"
	"hearthstone/internal/game"
	"hearthstone/internal/loop"
	"hearthstone/pkg/helper"
	"hearthstone/pkg/ui"
	"os"
	"strings"
)

type gameIO struct {
	game      game.Game
	errorChan chan error
	scanner   *bufio.Scanner
	inputChan chan loop.ActionEntry
}

func NewGameIO() *gameIO {
	return &gameIO{
		scanner:   bufio.NewScanner(os.Stdin),
		errorChan: make(chan error, 64),
		inputChan: make(chan loop.ActionEntry, 64),
	}
}

func (io *gameIO) Run(ctx context.Context) {
	for {
		if ctx.Err() != nil {
			return
		}
		if !io.scanner.Scan() {
			return
		}

		io.handleInput()
	}
}

func (io *gameIO) handleInput() {
	input := io.scanner.Text()
	input = strings.ToLower(input)
	allArgs := strings.Fields(input)

	command, args := "", []string{}
	if len(allArgs) > 0 {
		command, args = allArgs[0], allArgs[1:]
	}

	for _, action := range actionList {
		if !action.matches(command) {
			continue
		}

		idxes, sides, errs := parseAllPositions(args)
		if helper.FirstError(errs) != nil {
			io.SetErrors(NewInvalidArgumentsError(action))
			io.redraw()
			return
		}

		if action.errFunc != nil {
			io.SetErrors(action.errFunc())
			io.redraw()
			return
		}

		if action.prepareArgs != nil {
			idxes, sides = action.prepareArgs(idxes, sides)
		}

		if idxes == nil || sides == nil {
			io.SetErrors(NewInvalidArgumentsError(action))
			io.redraw()
			return
		}

		io.inputChan <- loop.ActionEntry{
			Name:  action.name,
			Idxes: idxes,
			Sides: sides,
		}
		return
	}

	io.SetErrors(NewShortHelpError())
}

func (io *gameIO) Input() <-chan loop.ActionEntry {
	return io.inputChan
}

func (io *gameIO) SetErrors(errs ...error) {
	for _, err := range errs {
		io.errorChan <- err
	}
}

func (io *gameIO) Redraw(g game.Game) {
	io.game = g
	io.redraw()
}

func (io *gameIO) redraw() {
	builder := strings.Builder{}

errLoop:
	for {
		select {
		case err := <-io.errorChan:
			if err != nil {
				fmt.Fprintln(&builder, tuiError(err))
			}
		default:
			break errLoop
		}
	}

	ui.UpdateFrame(
		gameString(io.game, builder.String()),
	)
}
