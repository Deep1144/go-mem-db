package helper

import (
	"fmt"
	"os"
)

// AppsDataPath returns the path to the user-specific configuration directory for apps.
// It retrieves the path using os.UserConfigDir() function and returns it as a string.
// If an error occurs while getting the path, it will print the error message and exit the program.
func AppsDataPath() string {
	appConfigDir, err := os.UserConfigDir()
	if err != nil {
		fmt.Printf("Failed to get apps config directory: %s\n", err)
		os.Exit(1)
	}
	return appConfigDir
}
