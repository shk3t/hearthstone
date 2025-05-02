package core

import (
	"strconv"
	"strings"
)

// Subtracts 1 from specified position to get index.
// "h" means Hero, converts to -1 index.
// "0" position also can be considered as hero position.
func parseIndexFromPosition(arg string) (int, error) {
	if strings.HasPrefix(arg, "h") {
		return -1, nil
	}
	pos, err := strconv.Atoi(arg)
	if err != nil {
		return 0, err
	}
	return pos - 1, nil
}