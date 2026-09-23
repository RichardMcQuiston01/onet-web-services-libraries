package onet

// Occupation is a reference to an O*NET occupation, used across many responses.
type Occupation struct {
	Href  string          `json:"href,omitempty"`
	Code  string          `json:"code"`
	Title string          `json:"title"`
	Tags  *OccupationTags `json:"tags,omitempty"`
}

// OccupationTags holds boolean flags that may appear on an Occupation.
type OccupationTags struct {
	BrightOutlook  bool `json:"bright_outlook,omitempty"`
	Green          bool `json:"green,omitempty"`
	Apprenticeship bool `json:"apprenticeship,omitempty"`
}

// Career is the My Next Move variant of an occupation reference.
type Career struct {
	Href  string `json:"href,omitempty"`
	Code  string `json:"code"`
	Title string `json:"title"`
}

// Element represents a scored occupational attribute (skill, knowledge, ability, etc.).
type Element struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description,omitempty"`
	Score       *Score  `json:"score,omitempty"`
}

// Score holds a numeric rating for an Element, as returned by summary and detail endpoints.
type Score struct {
	Value         float64 `json:"value"`
	Scale         string  `json:"scale,omitempty"`
	N             int     `json:"n,omitempty"`
	StandardError float64 `json:"standard_error,omitempty"`
	LowerCI       float64 `json:"lower_ci_bound,omitempty"`
	UpperCI       float64 `json:"upper_ci_bound,omitempty"`
}

// NamedItem is a simple id+name pair used in several list responses.
type NamedItem struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name"`
}
