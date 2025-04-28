package game

type CardPickError struct {
}
type EmptyHandError struct {
}
type FullTableAreaError struct {
}

func (err CardPickError) Error() string {
	return "Invalid card pick"
}
func (err EmptyHandError) Error() string {
	return "Hand is empty"
}
func (err FullTableAreaError) Error() string {
	return "Table is full"
}

func NewCardPickError() CardPickError {
	return CardPickError{}
}
func NewEmptyHandError() EmptyHandError {
	return EmptyHandError{}
}
func NewFullTableAreaError() FullTableAreaError {
	return FullTableAreaError{}
}