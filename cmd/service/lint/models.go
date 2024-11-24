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
						//"number_of_row": map[string]string{"type": "string"},
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
	Type       				string     			`json:"type"`
	Properties 				[]SchemaProperties 	`json:"properties"`
	Required             	[]string          	`json:"required"`
	AdditionalProperties 	bool              	`json:"additionalProperties"`
}

type SchemaProperties struct {
	Issues IssuesProperty `json:"issues"`
}

type IssuesProperty struct {
	Type  string          `json:"type"`
	Items IssuesItems     `json:"items"`
}

type IssuesItems struct {
	Type                 string           `json:"type"`
	Properties           ItemProperties   `json:"properties"`
	Required             []string         `json:"required"`
	AdditionalProperties bool             `json:"additionalProperties"`
}

type ItemProperties struct {
	//NumberOfRow PropertyType `json:"number_of_row"`
	Issue       PropertyType `json:"issue"`
	Severity    PropertyType `json:"severity"`
	Advice      PropertyType `json:"advice"`
}

type PropertyType struct {
	Type string `json:"type"`
}
