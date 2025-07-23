package game

type NextAction struct {
	Do        func(idxes []int, sides Sides) error
	OnSuccess func()
	OnFail    func()
}