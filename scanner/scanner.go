package scanner

import (
	"os/exec"
	"strings"
)

// ScanImage uses scanimage in the system to get the raw bytes
func ScanImage() ([]byte, error) {
	cmd := exec.Command("scanimage", "--format=jpeg", "--resolution=300")

	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	// The output of scanimage will contain the scanned image in JPEG format.
	// You can do any necessary processing here, such as saving the image to disk,
	// resizing it, etc. For simplicity, we're just returning the raw bytes.

	// Extracting image bytes from the output
	parts := strings.Split(string(output), "\n")
	startIndex := -1
	for i, part := range parts {
		if part == "##end-of-file" {
			startIndex = i
			break
		}
	}

	if startIndex == -1 {
		return nil, nil // Image not found in output
	}

	imageBytes := []byte(strings.Join(parts[startIndex+1:], "\n"))
	return imageBytes, nil
}
