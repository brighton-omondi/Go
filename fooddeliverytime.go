package piscine

type food struct {
	preptime int
}

// List of menu items and time it takes to prepare them
var menu = map[string]food{
	"burger":  {preptime: 15},
	"chips":   {preptime: 10},
	"nuggets": {preptime: 12},
}

func FoodDeliveryTime(order string) int {
	if item, available := menu[order]; available {
		return item.preptime
	}
	// Return 404 if the item is not in the menu
	return 404
}
