package ascii

import (
	"fmt"
	"os"
	"strings"
)

const chunkSize int = 9

// Convert .txt file of the theme into lines.
func ThemeToLines(s string) []string {
	data, _ := os.ReadFile("../pkg/ascii/theme/" + s + ".txt")
	return strings.Split(string(data), "\n")
}

// Get text input from arguments and split it into differente lines.
func GetTextInput(input string) []string {
	inputCorrected := strings.Replace(input, "\\n", "\n", -1)
	return strings.Split(inputCorrected, "\n")
}

// Print the ascii-art.
func PrintAsciiArt(input, lines []string) string {
	res := ""
	for _, inputLine := range input {
		if inputLine == "" {
			res += "\n"
		} else {
			for line := 1; line < chunkSize; line++ {
				for char := 0; char < len(inputLine); char++ {
					characterStart := (int(inputLine[char]) - 32) * chunkSize
					res += lines[characterStart+line]
				}
				res += "\n"
			}
		}
	}
	return res
}

func SaveOutput(text, fileName string) {
	output := CreateFile(fileName)
	output.WriteString(text)
}

func CreateFile(path string) *os.File {
	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	return file
}

func CheckArgs(args []string) string {
	if len(args) <= 1 {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
		os.Exit(0)
	} else if !ThemeIsValid(args) {
		fmt.Println("Please provide a valid theme: [shadow/standard/thinkertoy]\nEX: go run . --output=<fileName.txt> something standard")
		os.Exit(0)
	} else if len(args) == 3 {
		if !strings.HasPrefix(args[0], "--output=") {
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\nEX: go run . --output=<fileName.txt> something standard")
			os.Exit(0)
		} else {
			fileName := args[0][9:]
			if !strings.HasSuffix(fileName, ".txt") {
				fileName += ".txt"
			}
			lines := ThemeToLines(args[2])
			input := GetTextInput(args[1])
			text := PrintAsciiArt(input, lines)
			SaveOutput(text, fileName)
			return text
		}
	} else if len(args) == 2 {
		lines := ThemeToLines(args[1])
		input := GetTextInput(args[0])
		return PrintAsciiArt(input, lines)
	} else {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
		os.Exit(0)
	}
	return ""
}

func ThemeIsValid(args []string) bool {
	return args[len(args)-1] == "shadow" ||
		args[len(args)-1] == "standard" ||
		args[len(args)-1] == "thinkertoy"
}
