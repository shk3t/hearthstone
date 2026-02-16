package tui

import (
	"fmt"
	"hearthstone/internal/game"
	"hearthstone/pkg/ui"
	"strings"

	"github.com/fatih/color"
)

func cardInfo(c game.Card, fmtFunc ui.FormatFunc) string {
	builder := strings.Builder{}

	name := ui.BoldString(c.Name)
	if fmtFunc != nil {
		name = fmtFunc(name)
	}

	fmt.Fprintln(&builder, name)

	if c.Description != "" {
		fmt.Fprintln(&builder, c.Description)
	}

	fmt.Fprintf(&builder,
		"%s     %s",
		color.HiBlackString("Мана:"),
		color.BlueString("%d", c.ManaCost),
	)

	return builder.String()
}