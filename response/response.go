package response

type JsonMeta struct {
	Paragraphs int `json:"numParagraphs"`
	Lines      int `json:"linesPerParagraph"`
}
type JsonResponse struct {
	Character string   `json:"character"`
	Text      []string `json:"text"`
	Meta      JsonMeta `json:"meta"`
}
