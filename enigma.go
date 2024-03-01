package piscine

func Enigma(a ***int, b *int, c *******int, d ****int) {
	// Store original values
	tempA := ***a
	tempC := *******c
	tempD := ****d

	// Move values around to hide them
	***a = *b
	*b = tempD
	*******c = tempA
	****d = tempC
}
