package main

func QuadB(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}
	for i := 0; i < x; i++ {
		if i == 0 {
			print("/")
		} else if i == x-1 {
			print("\\")
		} else {
			print("*")
		}
	}
	println()
	for j := 1; j < y-1; j++ {
		for i := 0; i < x; i++ {
			if i == 0 || i == x-1 {
				print("*")
			} else {
				print(" ")
			}
		}
		println()
	}
	if y > 1 {
		for i := 0; i < x; i++ {
			if i == 0 {
				print("\\")
			} else if i == x-1 {
				print("/")
			} else {
				print("*")
			}
		}
		println()
	}
}
func main() {
	QuadB(9, 3)
}
