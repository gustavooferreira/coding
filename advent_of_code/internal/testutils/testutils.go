package testutils

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// TextInfo contains information about the text to be used in a test.
// Only set one of the 2 options.
// If both are set, the content in the file referenced by FilePath will take precedence!
type TextInfo struct {
	Content  string
	FilePath string
}

// GetTextContent returns the text content, either from the field Content or by reading the file referenced in the
// field FilePath.
func (ti *TextInfo) GetTextContent(t *testing.T) []string {
	if ti.FilePath != "" {
		return ReadFile(t, ti.FilePath)
	}

	return strings.Split(ti.Content, "\n")
}

// Function to read file and return contents as a slice of strings.
func ReadFile(t *testing.T, filePath string) []string {
	t.Helper()

	data, err := os.ReadFile(filePath)
	require.NoError(t, err)

	stringData := string(data)

	return strings.Split(stringData, "\n")
}
