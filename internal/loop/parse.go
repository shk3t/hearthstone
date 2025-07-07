package loop

import (
	"hearthstone/internal/game"
	"strconv"
	"strings"
)

// Subtracts 1 from specified position to get index.
//
// "h" means Hero, converts to -1 index.
// "0" position also can be considered as hero position.
//
// "t"/"b" for precise specifying target side (Top/Bottom).
func parsePosition(arg string) (idx int, side game.Side, err error) {
	switch {
	case strings.Contains(arg, "t"):
		arg = strings.Trim(arg, "t")
		side = game.TopSide
	case strings.Contains(arg, "b"):
		arg = strings.Trim(arg, "b")
		side = game.BotSide
	default:
		side = game.UnsetSide
	}

	if strings.Contains(arg, "h") {
		return game.HeroIdx, side, nil
	}

	pos, err := strconv.Atoi(arg)
	if err != nil {
		return 0, game.UnsetSide, err
	}

	return pos - 1, side, nil
}

func parseAllPositions(args []string) (idxes []int, sides []game.Side, errs []error) {
	lenArgs := len(args)
	idxes = make([]int, lenArgs)
	sides = make([]game.Side, lenArgs)
	errs = make([]error, lenArgs)

	for i := range args {
		idx, side, err := parsePosition(args[i])
		idxes[i] = idx
		sides[i] = side
		errs[i] = err
	}

	return idxes, sides, errs
}