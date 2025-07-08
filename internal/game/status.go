package game

type CharacterStatus struct {
	sleep       bool
	freezeTurns int
}

func (cs *CharacterStatus) IsSleep() bool {
	return cs.sleep
}

func (cs *CharacterStatus) SetSleep(value bool) {
	cs.sleep = value
}

func (cs *CharacterStatus) IsFreeze() bool {
	return cs.freezeTurns > 0
}

func (cs *CharacterStatus) SetFreeze(value bool) {
	if value {
		cs.freezeTurns = 2
	} else {
		cs.freezeTurns = 0
	}
}

func (cs *CharacterStatus) Unfreeze() {
	cs.freezeTurns = max(cs.freezeTurns-1, 0)
}
