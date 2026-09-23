package onet

import "context"

// OccupationElements is returned by element-based summary and detail endpoints
// (skills, knowledge, abilities, interests, work styles, work activities, work context).
type OccupationElements struct {
	Code    string    `json:"code"`
	Title   string    `json:"title"`
	Element []Element `json:"element"`
}

// Task is a work task statement for an occupation.
type Task struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	DWA  []struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"dwa,omitempty"`
}

// OccupationTasks is returned by OnlineSummaryTasks.
type OccupationTasks struct {
	Code  string `json:"code"`
	Title string `json:"title"`
	Task  []Task `json:"task"`
}

// TechnologyExample is a software or tool example within a technology skill category.
type TechnologyExample struct {
	HotTechnology bool   `json:"hot_technology,omitempty"`
	Name          string `json:"name"`
}

// TechnologyCategory groups technology examples under a label.
type TechnologyCategory struct {
	Title   string              `json:"title"`
	Example []TechnologyExample `json:"example"`
}

// OccupationTechnologySkills is returned by OnlineSummaryTechnologySkills.
type OccupationTechnologySkills struct {
	Code     string               `json:"code"`
	Title    string               `json:"title"`
	Category []TechnologyCategory `json:"category"`
}

// DetailedWorkActivity is a detailed work activity statement.
type DetailedWorkActivity struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// OccupationDetailedWorkActivities is returned by OnlineSummaryDetailedWorkActivities.
type OccupationDetailedWorkActivities struct {
	Code     string                 `json:"code"`
	Title    string                 `json:"title"`
	Activity []DetailedWorkActivity `json:"activity"`
}

// JobZoneInfo describes an O*NET job zone.
type JobZoneInfo struct {
	Value       int    `json:"value"`
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
	Education   string `json:"education,omitempty"`
}

// OccupationJobZone is returned by OnlineSummaryJobZone.
type OccupationJobZone struct {
	Code    string       `json:"code"`
	Title   string       `json:"title"`
	JobZone *JobZoneInfo `json:"job_zone,omitempty"`
}

// RapidsProgram is an apprenticeship program registered with RAPIDS.
type RapidsProgram struct {
	Code  string `json:"code"`
	Title string `json:"title"`
}

// OccupationApprenticeship is returned by OnlineSummaryApprenticeship.
type OccupationApprenticeship struct {
	Code   string          `json:"code"`
	Title  string          `json:"title"`
	Rapids []RapidsProgram `json:"rapids"`
}

// EducationCategory is a category in the education-level distribution for an occupation.
type EducationCategory struct {
	Name       string  `json:"name"`
	Percentage float64 `json:"percentage,omitempty"`
}

// EducationLevelRequired describes the typical education level for an occupation.
type EducationLevelRequired struct {
	Category []EducationCategory `json:"category"`
}

// OccupationEducation is returned by OnlineSummaryEducation.
type OccupationEducation struct {
	Code          string                  `json:"code"`
	Title         string                  `json:"title"`
	LevelRequired *EducationLevelRequired `json:"level_required,omitempty"`
}

// OccupationRelatedOccupations is returned by OnlineSummaryRelatedOccupations.
type OccupationRelatedOccupations struct {
	Code       string       `json:"code"`
	Title      string       `json:"title"`
	Occupation []Occupation `json:"occupation"`
}

// Association is a professional organization associated with an occupation.
type Association struct {
	Name string `json:"name"`
	URL  string `json:"url,omitempty"`
}

// OccupationProfessionalAssociations is returned by OnlineSummaryProfessionalAssociations.
type OccupationProfessionalAssociations struct {
	Code        string        `json:"code"`
	Title       string        `json:"title"`
	Association []Association `json:"association"`
}

// MilitaryCareer links a military occupation to a civilian O*NET occupation.
type MilitaryCareer struct {
	BranchID    string `json:"branch_id"`
	BranchTitle string `json:"branch_title"`
	Code        string `json:"code"`
	Title       string `json:"title"`
}

// OccupationMilitaryCareerSummaries is returned by OnlineSummaryMilitaryCareerSummaries.
type OccupationMilitaryCareerSummaries struct {
	Code     string           `json:"code"`
	Title    string           `json:"title"`
	Military []MilitaryCareer `json:"military"`
}

// OnlineSummaryTasks returns the task statements for the given occupation.
func (c *Client) OnlineSummaryTasks(ctx context.Context, code string) (*OccupationTasks, error) {
	var out OccupationTasks
	return &out, c.occupationResource(ctx, code, "summary", "tasks", &out)
}

// OnlineSummaryTechnologySkills returns the technology skills for the given occupation.
func (c *Client) OnlineSummaryTechnologySkills(ctx context.Context, code string) (*OccupationTechnologySkills, error) {
	var out OccupationTechnologySkills
	return &out, c.occupationResource(ctx, code, "summary", "technology_skills", &out)
}

// OnlineSummaryWorkActivities returns the work activities for the given occupation.
func (c *Client) OnlineSummaryWorkActivities(ctx context.Context, code string) (*OccupationElements, error) {
	var out OccupationElements
	return &out, c.occupationResource(ctx, code, "summary", "work_activities", &out)
}

// OnlineSummaryDetailedWorkActivities returns the detailed work activities for the given occupation.
func (c *Client) OnlineSummaryDetailedWorkActivities(ctx context.Context, code string) (*OccupationDetailedWorkActivities, error) {
	var out OccupationDetailedWorkActivities
	return &out, c.occupationResource(ctx, code, "summary", "detailed_work_activities", &out)
}

// OnlineSummaryWorkContext returns the work context factors for the given occupation.
func (c *Client) OnlineSummaryWorkContext(ctx context.Context, code string) (*OccupationElements, error) {
	var out OccupationElements
	return &out, c.occupationResource(ctx, code, "summary", "work_context", &out)
}

// OnlineSummaryJobZone returns the job zone for the given occupation.
func (c *Client) OnlineSummaryJobZone(ctx context.Context, code string) (*OccupationJobZone, error) {
	var out OccupationJobZone
	return &out, c.occupationResource(ctx, code, "summary", "job_zone", &out)
}

// OnlineSummaryApprenticeship returns apprenticeship programs for the given occupation.
func (c *Client) OnlineSummaryApprenticeship(ctx context.Context, code string) (*OccupationApprenticeship, error) {
	var out OccupationApprenticeship
	return &out, c.occupationResource(ctx, code, "summary", "apprenticeship", &out)
}

// OnlineSummarySkills returns the skills for the given occupation.
func (c *Client) OnlineSummarySkills(ctx context.Context, code string) (*OccupationElements, error) {
	var out OccupationElements
	return &out, c.occupationResource(ctx, code, "summary", "skills", &out)
}

// OnlineSummaryKnowledge returns the knowledge areas for the given occupation.
func (c *Client) OnlineSummaryKnowledge(ctx context.Context, code string) (*OccupationElements, error) {
	var out OccupationElements
	return &out, c.occupationResource(ctx, code, "summary", "knowledge", &out)
}

// OnlineSummaryEducation returns the education requirements for the given occupation.
func (c *Client) OnlineSummaryEducation(ctx context.Context, code string) (*OccupationEducation, error) {
	var out OccupationEducation
	return &out, c.occupationResource(ctx, code, "summary", "education", &out)
}

// OnlineSummaryAbilities returns the abilities for the given occupation.
func (c *Client) OnlineSummaryAbilities(ctx context.Context, code string) (*OccupationElements, error) {
	var out OccupationElements
	return &out, c.occupationResource(ctx, code, "summary", "abilities", &out)
}

// OnlineSummaryInterests returns the interests for the given occupation.
func (c *Client) OnlineSummaryInterests(ctx context.Context, code string) (*OccupationElements, error) {
	var out OccupationElements
	return &out, c.occupationResource(ctx, code, "summary", "interests", &out)
}

// OnlineSummaryWorkStyles returns the work styles for the given occupation.
func (c *Client) OnlineSummaryWorkStyles(ctx context.Context, code string) (*OccupationElements, error) {
	var out OccupationElements
	return &out, c.occupationResource(ctx, code, "summary", "work_styles", &out)
}

// OnlineSummaryRelatedOccupations returns occupations related to the given occupation.
func (c *Client) OnlineSummaryRelatedOccupations(ctx context.Context, code string) (*OccupationRelatedOccupations, error) {
	var out OccupationRelatedOccupations
	return &out, c.occupationResource(ctx, code, "summary", "related_occupations", &out)
}

// OnlineSummaryProfessionalAssociations returns professional associations for the given occupation.
func (c *Client) OnlineSummaryProfessionalAssociations(ctx context.Context, code string) (*OccupationProfessionalAssociations, error) {
	var out OccupationProfessionalAssociations
	return &out, c.occupationResource(ctx, code, "summary", "professional_associations", &out)
}

// OnlineSummaryMilitaryCareerSummaries returns military occupations linked to the given occupation.
func (c *Client) OnlineSummaryMilitaryCareerSummaries(ctx context.Context, code string) (*OccupationMilitaryCareerSummaries, error) {
	var out OccupationMilitaryCareerSummaries
	return &out, c.occupationResource(ctx, code, "summary", "military_career_summaries", &out)
}
