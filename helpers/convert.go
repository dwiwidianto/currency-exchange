package helpers

import "strconv"

func StringToInt(str string) (int, error) {
	convInt, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}
	return int(convInt), nil
}
