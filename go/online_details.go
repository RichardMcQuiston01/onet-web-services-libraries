package onet

import "context"

// OnlineDetailsTasks returns detailed task data (with scores) for the given occupation.
func (c *Client) OnlineDetailsTasks(ctx context.Context, code string) (*OccupationTasks, error) {
	var out OccupationTasks
	return &out, c.occupationResource(ctx, code, "details", "tasks", &out)
}

// OnlineDetailsTechnologySkills returns detailed technology skills for the given occupation.
func (c *Client) OnlineDetailsTechnologySkills(ctx context.Context, code string) (*OccupationTechnologySkills, error) {
	var out OccupationTechnologySkills
	return &out, c.occupationResource(ctx, code, "details", "technology_skills", &out)
}

// OnlineDetailsWorkActivities returns detailed work activities (with scores) for the given occupation.
func (c *Client) OnlineDetailsWorkActivities(ctx context.Context, code string) (*OccupationElements, error) {
	var out OccupationElements
	return &out, c.occupationResource(ctx, code, "details", "work_activities", &out)
}

// OnlineDetailsDetailedWorkActivities returns detailed work activity statements for the given occupation.
func (c *Client) OnlineDetailsDetailedWorkActivities(ctx context.Context, code string) (*OccupationDetailedWorkActivities, error) {
	var out OccupationDetailedWorkActivities
	return &out, c.occupationResource(ctx, code, "details", "detailed_work_activities", &out)
}

// OnlineDetailsWorkContext returns detailed work context factors (with scores) for the given occupation.
func (c *Client) OnlineDetailsWorkContext(ctx context.Context, code string) (*OccupationElements, error) {
	var out OccupationElements
	return &out, c.occupationResource(ctx, code, "details", "work_context", &out)
}

// OnlineDetailsJobZone returns the job zone (with description) for the given occupation.
func (c *Client) OnlineDetailsJobZone(ctx context.Context, code string) (*OccupationJobZone, error) {
	var out OccupationJobZone
	return &out, c.occupationResource(ctx, code, "details", "job_zone", &out)
}

// OnlineDetailsApprenticeship returns apprenticeship data for the given occupation.
func (c *Client) OnlineDetailsApprenticeship(ctx context.Context, code string) (*OccupationApprenticeship, error) {
	var out OccupationApprenticeship
	return &out, c.occupationResource(ctx, code, "details", "apprenticeship", &out)
}

// OnlineDetailsSkills returns detailed skills (with scores) for the given occupation.
func (c *Client) OnlineDetailsSkills(ctx context.Context, code string) (*OccupationElements, error) {
	var out OccupationElements
	return &out, c.occupationResource(ctx, code, "details", "skills", &out)
}

// OnlineDetailsKnowledge returns detailed knowledge areas (with scores) for the given occupation.
func (c *Client) OnlineDetailsKnowledge(ctx context.Context, code string) (*OccupationElements, error) {
	var out OccupationElements
	return &out, c.occupationResource(ctx, code, "details", "knowledge", &out)
}

// OnlineDetailsEducation returns detailed education requirements for the given occupation.
func (c *Client) OnlineDetailsEducation(ctx context.Context, code string) (*OccupationEducation, error) {
	var out OccupationEducation
	return &out, c.occupationResource(ctx, code, "details", "education", &out)
}

// OnlineDetailsAbilities returns detailed abilities (with scores) for the given occupation.
func (c *Client) OnlineDetailsAbilities(ctx context.Context, code string) (*OccupationElements, error) {
	var out OccupationElements
	return &out, c.occupationResource(ctx, code, "details", "abilities", &out)
}

// OnlineDetailsInterests returns detailed interests (with scores) for the given occupation.
func (c *Client) OnlineDetailsInterests(ctx context.Context, code string) (*OccupationElements, error) {
	var out OccupationElements
	return &out, c.occupationResource(ctx, code, "details", "interests", &out)
}

// OnlineDetailsWorkStyles returns detailed work styles (with scores) for the given occupation.
func (c *Client) OnlineDetailsWorkStyles(ctx context.Context, code string) (*OccupationElements, error) {
	var out OccupationElements
	return &out, c.occupationResource(ctx, code, "details", "work_styles", &out)
}

// OnlineDetailsRelatedOccupations returns related occupations for the given occupation.
func (c *Client) OnlineDetailsRelatedOccupations(ctx context.Context, code string) (*OccupationRelatedOccupations, error) {
	var out OccupationRelatedOccupations
	return &out, c.occupationResource(ctx, code, "details", "related_occupations", &out)
}

// OnlineDetailsProfessionalAssociations returns professional associations for the given occupation.
func (c *Client) OnlineDetailsProfessionalAssociations(ctx context.Context, code string) (*OccupationProfessionalAssociations, error) {
	var out OccupationProfessionalAssociations
	return &out, c.occupationResource(ctx, code, "details", "professional_associations", &out)
}

// OnlineDetailsMilitaryCareerSummaries returns military occupations linked to the given occupation.
func (c *Client) OnlineDetailsMilitaryCareerSummaries(ctx context.Context, code string) (*OccupationMilitaryCareerSummaries, error) {
	var out OccupationMilitaryCareerSummaries
	return &out, c.occupationResource(ctx, code, "details", "military_career_summaries", &out)
}
