package Brightness

type BrightnessSet []Brightness

func (list *BrightnessSet) Append(object Brightness) {
	*list = append(*list, object)
}
