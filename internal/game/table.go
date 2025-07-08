package game

type Table [SidesCount]TableArea

func NewTable() *Table {
	return &Table{
		newTableArea(TopSide),
		newTableArea(BotSide),
	}
}

func (t *Table) CleanupDeadMinions() {
	t[TopSide].cleanupDeadMinions()
	t[BotSide].cleanupDeadMinions()
}