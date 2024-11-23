package lint

// Issue represents a single linting issue found in the Dockerfile.
type Issue struct {
	NumberOfRow string `json:"number_of_row"` // The line number(s) where the issue is found.
	Issue       string `json:"issue"`         // Description of the issue.
	Severity    string `json:"severity"`      // Severity level of the issue (e.g., high, medium, low).
	Advice      string `json:"advice"`        // Recommended action to resolve the issue.
}

// LintResponse represents a list of linting issues.
type LintResponse struct {
	Issues []Issue `json:"issues"`
}

type Dockerfile struct {
	Content string
}



// Define the JSON schema for the response
/*
	schema := map[string]interface{}{
		"type": "object",
		"properties": map[string]interface{}{
			"issues": map[string]interface{}{
				"type": "array",
				"items": map[string]interface{}{
					"type": "object",
					"properties": map[string]interface{}{
						"number_of_row": map[string]string{"type": "string"},
						"issue":         map[string]string{"type": "string"},
						"severity":      map[string]string{"type": "string"},
						"advice":        map[string]string{"type": "string"},
					},
					"required":             []string{"number_of_row", "issue", "severity", "advice"},
					"additionalProperties": false,
				},
			},
		},
		"required":             []string{"issues"},
		"additionalProperties": false,
	}*/

type LinterChatGPTSchema struct {
	Type       string     `json:"type"`
	Properties []Property `json:"properties"`
}

type Property struct {
	Required             []string   `json:"required"`
	AdditionalProperties bool       `json:"additionalProperties"`
	Issues               []GPTIssue `json:"issues"`
}

type GPTIssue struct {
	//map[string]string{"type": "string"},
	NumberOfRow string `json:"number_of_row"`
	Type string `json:"type"`
}
