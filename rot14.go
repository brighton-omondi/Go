package piscine

func Rot14(s string) string {
	var result string

	for _, char := range s {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
			var base byte
			if char >= 'a' && char <= 'z' {
				base = 'a'
			} else {
				base = 'A'
			}
			rotated := byte((int(char)-int(base)+14)%26 + int(base))
			result += string(rotated)
		} else {
			result += string(char)
		}
	}

	return result
}
