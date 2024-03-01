package piscine

func ReverseMenuIndex(menu []string) []string {
	reversedMenu := make([]string, len(menu)) // Create a new slice with the same length as the menu

	for i := range menu {
		reversedMenu[len(menu)-1-i] = menu[i] // Fill the new slice with menu items in reverse order
	}

	return reversedMenu
}
