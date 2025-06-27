package game

import (
	"strings"
)

type Hero struct {
	Character
	Class       Class
	Weapon      *Weapon
	Power       Spell
	PowerIsUsed bool
}

const HeroIdx = -1

func (h *Hero) String() string {
	elems := make([]string, 0, 3)
	elems = append(elems, string(h.Class))

	statusStr := h.Status.String()
	if statusStr != "" {
		elems = append(elems, statusStr)
	}

	if !h.PowerIsUsed {
		elems = append(elems, h.Power.String())
	}

	return strings.Join(elems, " | ")
}

func (h *Hero) Copy() *Hero {
	heroCopy := *h
	return &heroCopy
}

func (h *Hero) healthString() string {
	return playerBarString("Здоровье:", h.Health, h.MaxHealth, "+")
}