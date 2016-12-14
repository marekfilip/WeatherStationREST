package Config

const _SECRET_KEY string = "aAbskq90wenQWEa@23!#!1ASDLqwkm()"

func GetSecret() []byte {
	return []byte(_SECRET_KEY)
}
