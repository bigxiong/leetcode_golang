package main

func lemonadeChange(bills []int) bool {
	var five, ten int
	for _, b := range bills {
		if b == 5 {
			five++
		} else if b == 10 {
			if five == 0 {
				return false
			} else {
				five--
			}
			ten++
		} else {
			if five > 0 && ten > 0 {
				five--
				ten--
			} else if five > 3 {
				five -= 3
			} else {
				return false
			}
		}
	}
	return true
}

func main() {

}
