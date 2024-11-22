package lint

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"go.uber.org/zap"
)

func FetchBestPracticesMarkdown(url string, logger *zap.SugaredLogger) (string, error) {
	logger.Infof("Fetching best practices from URL: %s", url)

	// Scarica il file Markdown
	resp, err := http.Get(url)
	if err != nil {
		logger.Errorf("Error fetching Markdown file: %v", err)
		return "", fmt.Errorf("failed to fetch Markdown file: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		logger.Errorf("Unexpected status code: %d", resp.StatusCode)
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Legge il contenuto del file
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Errorf("Error reading response body: %v", err)
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	// Analizza il contenuto Markdown
	content := string(body)
	var sections []string
	lines := strings.Split(content, "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)
		// Identifica i titoli principali
		if strings.HasPrefix(line, "#") {
			sections = append(sections, "\n"+line) // Aggiunge una nuova sezione
		} else if len(line) > 0 {
			// Aggiunge paragrafi o contenuti sotto il titolo
			if len(sections) > 0 {
				sections[len(sections)-1] += "\n" + line
			}
		}
	}

	if len(sections) == 0 {
		logger.Warn("No relevant content found in the Markdown file.")
		return "", fmt.Errorf("no relevant content found")
	}

	// Combina le sezioni in un risultato leggibile
	result := strings.Join(sections, "\n\n")
	logger.Info("Best practices fetched successfully.")
	return result, nil
}
