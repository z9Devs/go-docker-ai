package lint

type LintResponse struct {
	Content string `json:"content"`
}

type Dockerfile struct {
	Content string
}
