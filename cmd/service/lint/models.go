package lint

// Issue represents a single linting issue found in the Dockerfile.
type Issue struct {
	Issue    string `json:"issue"`    // Description of the issue.
	Severity string `json:"severity"` // Severity level of the issue (e.g., high, medium, low).
	Advice   string `json:"advice"`   // Recommended action to resolve the issue.
}

// LintResponse represents a list of linting issues.
type LintResponse struct {
	Issues []Issue `json:"issues"`
}

type Dockerfile struct {
	Content string
}