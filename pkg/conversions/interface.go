package conversions

func TrueNilInterfaceSlice[S any, I any](input []*S) []I {
	output := make([]I, len(input))
	for i, v := range input {
		if v == nil {
			output[i] = *new(I)
		} else {
			output[i] = any(v).(I)
		}
	}
	return output
}