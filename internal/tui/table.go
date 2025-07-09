package tui

import (
	"fmt"
	"hearthstone/internal/game"
	"strings"
)

func tableString(t *game.Table) string {
	builder := strings.Builder{}
	fmt.Fprintln(&builder)
	fmt.Fprintln(&builder, strings.Repeat("=", 50))
	fmt.Fprintln(&builder, tableAreaString(t[game.TopSide]))
	fmt.Fprintln(&builder, strings.Repeat("-", 50))
	fmt.Fprintln(&builder, tableAreaString(t[game.BotSide]))
	fmt.Fprintln(&builder, strings.Repeat("=", 50))
	fmt.Fprintln(&builder)
	return builder.String()
}