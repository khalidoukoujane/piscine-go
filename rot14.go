package piscine

func Rot14(s string) string {
	str := ""
	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			str += string(((char+14)-'a')%26 + 'a')
		} else if char >= 'A' && char <= 'Z' {
			str += string(((char+14)-'A')%26 + 'A')
		} else {
			str += string(char)
		}
	}
	return string(str)
}
