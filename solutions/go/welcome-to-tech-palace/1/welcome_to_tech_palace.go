package techpalace

import "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	stat := strings.Repeat("*", numStarsPerLine)
	return stat + "\n" + welcomeMsg + "\n" + stat
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
		result := strings.ReplaceAll(oldMsg, "*", "")
		return strings.TrimSpace(result)
	}
