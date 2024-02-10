package iteration

const repeatCount = 5

func Repeat(character string, repeat int) string {
	var repeated string
	if (repeat <= 0) {
		repeat = repeatCount
	}
	for i := 0; i < repeat; i++ {
		repeated += character
	}
	return repeated
}

