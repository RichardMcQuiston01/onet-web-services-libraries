package onet

import (
	"context"
	"net/url"
	"strings"
)

// JobDutiesTask is a task offered for selection in the Job Duties finder.
type JobDutiesTask struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// JobDutiesTasksResponse is returned by OnlineJobDutiesTasks.
type JobDutiesTasksResponse struct {
	Code  string          `json:"code"`
	Title string          `json:"title"`
	Task  []JobDutiesTask `json:"task"`
}

// OccupationMatch is an occupation returned by finder result endpoints,
// optionally including a relevance score.
type OccupationMatch struct {
	Code  string  `json:"code"`
	Title string  `json:"title"`
	Score float64 `json:"score,omitempty"`
}

// JobDutiesResultsResponse is returned by OnlineJobDutiesResults.
type JobDutiesResultsResponse struct {
	Occupation []OccupationMatch `json:"occupation"`
}

// OnlineJobDutiesSearch searches occupations by job duty keyword.
func (c *Client) OnlineJobDutiesSearch(ctx context.Context, keyword string, page PageParams) (*OccupationList, error) {
	q := url.Values{"keyword": {keyword}}
	page.apply(q)
	var out OccupationList
	return &out, c.do(ctx, "online/job_duties/search", q, &out)
}

// OnlineJobDutiesTasks returns tasks for a given occupation for use in the Job Duties finder.
func (c *Client) OnlineJobDutiesTasks(ctx context.Context, code string) (*JobDutiesTasksResponse, error) {
	var out JobDutiesTasksResponse
	return &out, c.do(ctx, "online/job_duties/tasks/"+code, nil, &out)
}

// OnlineJobDutiesResults returns occupations that match the selected task IDs for the given occupation.
// taskIDs is a list of task IDs from OnlineJobDutiesTasks.
func (c *Client) OnlineJobDutiesResults(ctx context.Context, code string, taskIDs []string) (*JobDutiesResultsResponse, error) {
	q := url.Values{"tasks": {strings.Join(taskIDs, ",")}}
	var out JobDutiesResultsResponse
	return &out, c.do(ctx, "online/job_duties/results/"+code, q, &out)
}

// AssociationList is a paginated list of professional associations.
type AssociationList struct {
	PageMeta
	Association []Association `json:"association"`
}

// OnlineAssociationsSearch searches occupations by professional association keyword.
func (c *Client) OnlineAssociationsSearch(ctx context.Context, keyword string, page PageParams) (*OccupationList, error) {
	q := url.Values{"keyword": {keyword}}
	page.apply(q)
	var out OccupationList
	return &out, c.do(ctx, "online/associations/search", q, &out)
}

// OnlineAssociations returns professional associations for the given occupation.
func (c *Client) OnlineAssociations(ctx context.Context, code string) (*OccupationProfessionalAssociations, error) {
	var out OccupationProfessionalAssociations
	return &out, c.do(ctx, "online/associations/"+code, nil, &out)
}

// OnlineAssociationsAll returns a paginated list of all professional associations.
func (c *Client) OnlineAssociationsAll(ctx context.Context, page PageParams) (*AssociationList, error) {
	q := url.Values{}
	page.apply(q)
	var out AssociationList
	return &out, c.do(ctx, "online/associations/all", q, &out)
}

// RelatedActivity is a work activity item used in the Related Activities finder.
type RelatedActivity struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// OccupationActivities is returned by OnlineRelatedActivities.
type OccupationActivities struct {
	Code     string            `json:"code"`
	Title    string            `json:"title"`
	Activity []RelatedActivity `json:"activity"`
}

// RelatedActivitiesResultsResponse is returned by OnlineRelatedActivitiesResults.
type RelatedActivitiesResultsResponse struct {
	Occupation []OccupationMatch `json:"occupation"`
}

// OnlineRelatedActivitiesSearch searches occupations by work activity keyword.
func (c *Client) OnlineRelatedActivitiesSearch(ctx context.Context, keyword string, page PageParams) (*OccupationList, error) {
	q := url.Values{"keyword": {keyword}}
	page.apply(q)
	var out OccupationList
	return &out, c.do(ctx, "online/related_activities/search", q, &out)
}

// OnlineRelatedActivities returns work activities for the given occupation for use in the finder.
func (c *Client) OnlineRelatedActivities(ctx context.Context, code string) (*OccupationActivities, error) {
	var out OccupationActivities
	return &out, c.do(ctx, "online/related_activities/activities/"+code, nil, &out)
}

// OnlineRelatedActivitiesResults returns occupations that match the selected activity IDs.
// activityIDs is a list of activity IDs from OnlineRelatedActivities.
func (c *Client) OnlineRelatedActivitiesResults(ctx context.Context, code string, activityIDs []string) (*RelatedActivitiesResultsResponse, error) {
	q := url.Values{"activities": {strings.Join(activityIDs, ",")}}
	var out RelatedActivitiesResultsResponse
	return &out, c.do(ctx, "online/related_activities/results/"+code, q, &out)
}

// SoftSkill is a soft skill available for selection in the Soft Skills finder.
type SoftSkill struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// SoftSkillsListResponse is returned by OnlineSoftSkills.
type SoftSkillsListResponse struct {
	Skill []SoftSkill `json:"skill"`
}

// OnlineSoftSkills returns the list of soft skills available for the Soft Skills finder.
func (c *Client) OnlineSoftSkills(ctx context.Context) (*SoftSkillsListResponse, error) {
	var out SoftSkillsListResponse
	return &out, c.do(ctx, "online/soft_skills/", nil, &out)
}

// OnlineSoftSkillsResults returns occupations that match the given soft skill IDs.
// skillIDs is a list of skill IDs from OnlineSoftSkills.
func (c *Client) OnlineSoftSkillsResults(ctx context.Context, skillIDs []string, page PageParams) (*OccupationList, error) {
	q := url.Values{"skills": {strings.Join(skillIDs, ",")}}
	page.apply(q)
	var out OccupationList
	return &out, c.do(ctx, "online/soft_skills/results", q, &out)
}

// TechnologyExampleList is a paginated list of technology examples.
type TechnologyExampleList struct {
	PageMeta
	Technology []HotTechnology `json:"technology"`
}

// TechnologyExampleResponse is returned by OnlineTechnologyExample.
type TechnologyExampleResponse struct {
	Technology HotTechnology `json:"technology"`
	Occupation []Occupation  `json:"occupation"`
}

// TechnologyCategoryBrief is a technology category as returned by search results.
type TechnologyCategoryBrief struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

// TechnologyCategorySearchList is a paginated list of technology categories.
type TechnologyCategorySearchList struct {
	PageMeta
	Category []TechnologyCategoryBrief `json:"category"`
}

// TechnologyCategoryOccupationsResponse is returned by OnlineTechnologyCategory.
type TechnologyCategoryOccupationsResponse struct {
	Category   TechnologyCategoryBrief `json:"category"`
	Occupation []Occupation            `json:"occupation"`
}

// OnlineTechnologyExamplesSearch searches technology tool and software examples by keyword.
func (c *Client) OnlineTechnologyExamplesSearch(ctx context.Context, keyword string, page PageParams) (*TechnologyExampleList, error) {
	q := url.Values{"keyword": {keyword}}
	page.apply(q)
	var out TechnologyExampleList
	return &out, c.do(ctx, "online/technology/examples/search", q, &out)
}

// OnlineTechnologyExample returns occupations that use the given technology example.
func (c *Client) OnlineTechnologyExample(ctx context.Context, title string) (*TechnologyExampleResponse, error) {
	var out TechnologyExampleResponse
	return &out, c.do(ctx, "online/technology/examples/"+title, nil, &out)
}

// OnlineTechnologyCategoriesSearch searches technology categories by keyword.
func (c *Client) OnlineTechnologyCategoriesSearch(ctx context.Context, keyword string, page PageParams) (*TechnologyCategorySearchList, error) {
	q := url.Values{"keyword": {keyword}}
	page.apply(q)
	var out TechnologyCategorySearchList
	return &out, c.do(ctx, "online/technology/categories/search", q, &out)
}

// OnlineTechnologyCategory returns occupations that use technology in the given category.
func (c *Client) OnlineTechnologyCategory(ctx context.Context, id string) (*TechnologyCategoryOccupationsResponse, error) {
	var out TechnologyCategoryOccupationsResponse
	return &out, c.do(ctx, "online/technology/categories/"+id, nil, &out)
}
