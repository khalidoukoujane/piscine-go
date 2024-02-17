package piscine

type food struct {
	preptime int
}

func FoodDeliveryTime(order string) int {
	Food := food{}
	if order == "burger" {
		Food.preptime = 15
	} else if order == "chips" {
		Food.preptime = 10
	} else if order == "nuggets" {
		Food.preptime = 12
	} else {
		Food.preptime = 404
	}
	return Food.preptime
}
