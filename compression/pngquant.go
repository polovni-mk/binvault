package compression

import (
	"binvault/models"
	"fmt"
	"os"
	"os/exec"
)

func compress(file *models.File) {
	// no return needed
}

func compressPNG(inputPath, outputPath string) error {
	cmd := exec.Command("pngquant/pngquant", "--quality=10", "--output", outputPath, "--force", inputPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func main() {
	inputFile := "input.png"
	outputFile := "compressed.png"

	err := compressPNG(inputFile, outputFile)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Compression successful:", outputFile)
	}
}