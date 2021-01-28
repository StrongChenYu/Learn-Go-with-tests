package main

const spanish = "Spanish"
const chinese = "Chinese"
const chineseHelloPrefix = "你好, "
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "


func Hello(name, language string) string {
	if name == "" {
		name = "world"
	}
	return choseDifferentPrefix(language) + name
}

func choseDifferentPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case chinese:
		prefix = chineseHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	//
}