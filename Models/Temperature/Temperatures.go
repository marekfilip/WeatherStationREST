package Temperature

type Temperatures []Temperature

func (list *Temperatures) Append(object Temperature) {
	*list = append(*list, object)
}
