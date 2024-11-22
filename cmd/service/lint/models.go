package lint


type LintIssue struct {
	LineNumber int    `json:"line_number"`
	Message    string `json:"message"`
	Severity   string `json:"severity"`
}

type LintResponse struct {
	Issues []LintIssue `json:"issues"`
}

type Dockerfile struct {
    Content string
}