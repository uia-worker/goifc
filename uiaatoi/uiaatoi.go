package uiaatoi

import (
	"fmt"
	"strconv"
)

type UiaAtoiError struct {
	Message string
}

func (e *UiaAtoiError) Error() string {
	return fmt.Sprintf("%s", e.Message)
}

func Uiaatoi(ascii string) (int, error) {
	i, err := strconv.Atoi(ascii)
	if err != nil {
		return 0, &UiaAtoiError{"FEIL: problemer med å konvertere ascii til integer."}
	}
	return i, nil
}
