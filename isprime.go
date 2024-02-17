package piscine

func IsPrime(nb int) bool {
	prime := true
	if nb <= 1 {
		return false
	}
	if nb == 2 {
		return true
	} else {
		for i := 2; i <= nb/i; i++ {
			if nb%i == 0 {
				return false
			}
		}
		return prime
	}
}
