package Temperature

type TemperatureSet []Temperature

func (list *TemperatureSet) Append(object Temperature) {
	*list = append(*list, object)
}
