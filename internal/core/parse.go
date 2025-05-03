package core

import (
	gamepkg "hearthstone/internal/game"
	"strconv"
	"strings"
)

// Subtracts 1 from specified position to get index.
//
// "h" means Hero, converts to -1 index.
// "0" position also can be considered as hero position.
//
// "t"/"b" for precise specifying target side (Top/Bottom).
func ParsePositions(arg string) (idx int, side gamepkg.Side, err error) {
	switch {
	case strings.Contains(arg, "t"):
		arg = strings.Trim(arg, "t")
		side = gamepkg.TopSide
	case strings.Contains(arg, "b"):
		arg = strings.Trim(arg, "b")
		side = gamepkg.BotSide
	default:
		side = gamepkg.UnsetSide
	}

	if strings.Contains(arg, "h") {
		return -1, side, nil
	}

	pos, err := strconv.Atoi(arg)
	if err != nil {
		return 0, gamepkg.UnsetSide, err
	}

	return pos - 1, side, nil
}