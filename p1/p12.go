package main


func rotate_counting(count, dial int, direction byte, n int) (int, int) {
	if direction == 'R' {
		count +=  (dial + n) / 100
		dial = (dial + n) % 100
	} else {
		if n >= dial {
			count += (100 - dial + n) / 100
			if dial == 0 {
				count -= 1
			}
		}
		// Make dial always positive
		dial = ((dial - n) % 100 + 100) % 100
	}

	return dial, count
}
