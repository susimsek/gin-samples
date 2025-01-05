package banner

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// ANSI color codes
const (
	AnsiColorRed     = "\033[31m" // Red text
	AnsiColorBlue    = "\033[34m" // Blue text
	AnsiColorGreen   = "\033[32m" // Green text
	AnsiColorDefault = "\033[0m"  // Default terminal color
)

// PrintBanner reads the banner file, processes placeholders, and prints it
func PrintBanner(filePath string) {
	// Read the banner file
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Failed to load banner file: %v", err)
	}

	// Convert the content to string
	banner := string(data)

	// Replace color placeholders with ANSI codes
	banner = processColors(banner)

	// Replace ${ENV} with the current environment
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "dev" // Default to "dev" if APP_ENV is not set
	}
	banner = strings.ReplaceAll(banner, "${ENV}", env)

	// Print the processed banner
	fmt.Println(banner)
}

// processColors replaces color placeholders with actual ANSI codes
func processColors(banner string) string {
	banner = strings.ReplaceAll(banner, "${AnsiColor.RED}", AnsiColorRed)
	banner = strings.ReplaceAll(banner, "${AnsiColor.BLUE}", AnsiColorBlue)
	banner = strings.ReplaceAll(banner, "${AnsiColor.GREEN}", AnsiColorGreen)
	banner = strings.ReplaceAll(banner, "${AnsiColor.DEFAULT}", AnsiColorDefault)
	return banner
}
