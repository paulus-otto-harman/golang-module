package gola

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func ClearScreen() {
	cmd := exec.Command("clear")
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func Wait(label string) {
	Input(Args(P("label", label)))
}

type Param struct {
	Key   string
	Value interface{}
}

func P(key string, value interface{}) Param {
	return Param{key, value}
}

func Args(params ...Param) map[string]interface{} {
	var parameters = make(map[string]interface{})
	for _, param := range params {
		parameters[param.Key] = param.Value
	}
	return parameters
}

func ToString(data interface{}, err error) (string, error) {
	// TODO : validasi
	return fmt.Sprintf("%v", data), err
}

func ToInt(data interface{}, err error) (int, error) {
	// TODO : validasi

	if err != nil {
		panic("Error Detected. System Halted.")
	}

	if number, ok := data.(int); !ok {
		return 0, err
	} else {
		return number, err
	}

}

func Input(params ...map[string]interface{}) (interface{}, error) {
	label := ""
	required := false

	if len(params) > 0 {
		if value, ok := params[0]["label"]; ok {
			label = fmt.Sprintf("%s ", value)
		}

		if value, ok := params[0]["required"]; ok {
			required = value.(bool)
		}

		if value, ok := params[0]["type"]; ok {
			if value == "number" {
				var inputAngka int

				if required {
					for err := errors.New(""); err != nil; {
						fmt.Print(label)
						_, err = fmt.Scanln(&inputAngka)
					}
					return inputAngka, nil
				}
				fmt.Print(label)
				_, err := fmt.Scanln(&inputAngka)
				return inputAngka, err
			}
		}

	}

	// default
	var inputTeks string
	if required {
		for err := errors.New(""); err != nil; {
			fmt.Print(label)
			_, err = fmt.Scanln(&inputTeks)
		}
		return inputTeks, nil
	}
	fmt.Print(label)
	_, err := fmt.Scanln(&inputTeks)
	return fmt.Sprintf("%v", inputTeks), err
}

const Red = 31
const Green = 32
const Yellow = 33
const Blue = 34
const Magenta = 35
const Cyan = 36
const LightGray = 37

const Gray = 90
const LightRed = 91
const LightGreen = 92
const LightYellow = 93
const LightBlue = 94
const LightMagenta = 95
const LightCyan = 96
const White = 97

const Bold = "\033[1m%s\033[0m" // ESC[1m
const Color = "\x1b[%dm%s\x1b[0m"

func Tf(mode string, teks string, warna ...int) string {
	switch {
	case mode == Bold && len(warna) > 0:
		return fmt.Sprintf(Color, warna[0], fmt.Sprintf(Bold, teks))
	case mode == Bold:
		return fmt.Sprintf(Bold, teks)
	default:
		return fmt.Sprintf(Color, warna[0], teks)
	}
}

func Test() {
	
}
