package core

import "strconv"

func parseIndexFromPosition(arg string) (int, error) {
	pos, err := strconv.Atoi(arg)
	if err != nil {
		return -1, err
	}
	return pos - 1, nil
}