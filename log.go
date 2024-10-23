package gola

import "fmt"

func Logger(teks string) string {
	return fmt.Sprintf("gola : %s", teks)
}
