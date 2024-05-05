package scanner

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

// ScanImage scans an image and saves it to a JPG file, returning the filename.

func ScanImage() (string, error) {
	cmd := exec.Command("scanimage", "--format=jpeg", "--resolution=300")
	log.Printf("Command scanimage executed\n")

	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	log.Printf("Output stored in var\n")

	// Extracting image bytes from the output
	parts := strings.Split(string(output), "\n")
	startIndex := -1
	for i, part := range parts {
		if part == "##end-of-file" {
			startIndex = i
			break
		}
	}
	log.Printf("Extracted image bytes from output\n")

	if startIndex == -1 {
		log.Printf("Image not found in output\n")
		return "", nil // Image not found in output
	}

	imageBytes := []byte(strings.Join(parts[startIndex+1:], "\n"))
	log.Printf("Scanned %d bytes\n", len(imageBytes))
	fileName := "/tmp/scanned_image.jpg"

	// Write image bytes to a JPG file
	err = os.WriteFile(fileName, imageBytes, 0644)
	if err != nil {
		return "", err
	}
	log.Printf("image file created\n")

	return fileName, nil
}
