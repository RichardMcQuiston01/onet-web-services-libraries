package onet

import (
	"context"
	"net/url"
	"strconv"
	"strings"
)

// CareerList is a paginated list of careers returned by MNM/MPP search endpoints.
type CareerList struct {
	PageMeta
	Career []Career `json:"career"`
}

// MNMEducation holds education requirements as embedded in MNMCareerDetail.
type MNMEducation struct {
	EducationUsuallyNeeded []string `json:"education_usually_needed,omitempty"`
	Apprenticeship         bool     `json:"apprenticeship,omitempty"`
}

// MNMOutlookValue is a labeled outlook or openings value.
type MNMOutlookValue struct {
	Category string `json:"category,omitempty"`
	Value    string `json:"value,omitempty"`
}

// MNMSalary holds median salary data.
type MNMSalary struct {
	AnnualMedian int     `json:"annual_median,omitempty"`
	HourlyMedian float64 `json:"hourly_median,omitempty"`
}

// MNMJobOutlookSummary holds job outlook data as embedded in MNMCareerDetail.
type MNMJobOutlookSummary struct {
	Outlook  *MNMOutlookValue `json:"outlook,omitempty"`
	Salary   *MNMSalary       `json:"salary,omitempty"`
	Openings *MNMOutlookValue `json:"openings,omitempty"`
}

// MNMCareerDetail is the full career record returned by MNMCareer/MPPCareer.
type MNMCareerDetail struct {
	Code        string                `json:"code"`
	Title       string                `json:"title"`
	AlsoCalled  []string              `json:"also_called,omitempty"`
	WhatTheyDo  string                `json:"what_they_do,omitempty"`
	OnTheJob    []string              `json:"on_the_job,omitempty"`
	Knowledge   []Element             `json:"knowledge,omitempty"`
	Skills      []Element             `json:"skills,omitempty"`
	Abilities   []Element             `json:"abilities,omitempty"`
	Personality []Element             `json:"personality,omitempty"`
	Technology  []TechnologyCategory  `json:"technology,omitempty"`
	Education   *MNMEducation         `json:"education,omitempty"`
	JobOutlook  *MNMJobOutlookSummary `json:"job_outlook,omitempty"`
}

// MNMCareerEducationResponse is returned by MNMCareerEducation/MPPCareerEducation.
type MNMCareerEducationResponse struct {
	Code                   string   `json:"code"`
	Title                  string   `json:"title"`
	EducationUsuallyNeeded []string `json:"education_usually_needed,omitempty"`
	Apprenticeship         bool     `json:"apprenticeship,omitempty"`
}

// MNMCareerJobOutlookResponse is returned by MNMCareerJobOutlook/MPPCareerJobOutlook.
type MNMCareerJobOutlookResponse struct {
	Code     string           `json:"code"`
	Title    string           `json:"title"`
	Outlook  *MNMOutlookValue `json:"outlook,omitempty"`
	Salary   *MNMSalary       `json:"salary,omitempty"`
	Openings *MNMOutlookValue `json:"openings,omitempty"`
}

// MNMStateItem is a US state entry in the career state response.
type MNMStateItem struct {
	Abbreviation string `json:"abbreviation"`
	Title        string `json:"title"`
}

// MNMCareerStateResponse is returned by MNMCareerState/MPPCareerState.
type MNMCareerStateResponse struct {
	Code  string         `json:"code"`
	Title string         `json:"title"`
	State []MNMStateItem `json:"state"`
}

// MNMIndustriesResponse is returned by MNMIndustries/MPPIndustries.
type MNMIndustriesResponse struct {
	Industry []Industry `json:"industry"`
}

// MNMIndustryCareersResponse is returned by MNMIndustryMostPeople and MNMIndustrySomePeople.
type MNMIndustryCareersResponse struct {
	Code   string   `json:"code"`
	Title  string   `json:"title"`
	Career []Career `json:"career"`
}

// InterestProfilerQuestion is one question from the O*NET Interest Profiler.
type InterestProfilerQuestion struct {
	Number int    `json:"number"`
	Area   string `json:"area"`
	Text   string `json:"text"`
}

// InterestProfilerQuestionsResponse is returned by MNMInterestProfilerQuestions variants.
type InterestProfilerQuestionsResponse struct {
	Question []InterestProfilerQuestion `json:"question"`
}

// RIASECScore is a scored RIASEC area from the interest profiler.
type RIASECScore struct {
	Area  string `json:"area"`
	Score int    `json:"score"`
}

// InterestProfilerResultsResponse is returned by MNMInterestProfilerResults.
type InterestProfilerResultsResponse struct {
	Result []RIASECScore `json:"result"`
}

// MNMInterest is a RIASEC interest area.
type MNMInterest struct {
	Code  string `json:"code"`
	Title string `json:"title"`
}

// MNMInterestsResponse is returned by MNMInterests/MPPInterests.
type MNMInterestsResponse struct {
	Interest []MNMInterest `json:"interest"`
}

// MNMInterestCareersResponse is returned by MNMInterestCareers/MPPInterestCareers.
type MNMInterestCareersResponse struct {
	Code   string   `json:"code"`
	Title  string   `json:"title"`
	Career []Career `json:"career"`
}

// MNMJobZoneItem is an entry in the MNM job zones list.
type MNMJobZoneItem struct {
	Number int    `json:"number"`
	Title  string `json:"title"`
}

// MNMJobZonesResponse is returned by MNMJobZones/MPPJobZones.
type MNMJobZonesResponse struct {
	JobZone []MNMJobZoneItem `json:"job_zone"`
}

// MNMBrightOutlookCategoriesResponse is returned by MNMBrightOutlookCategories.
type MNMBrightOutlookCategoriesResponse struct {
	Category []BrightOutlookCategory `json:"category"`
}

// MNMBrightOutlookCareersResponse is returned by MNMBrightOutlookCareers.
type MNMBrightOutlookCareersResponse struct {
	Code   string   `json:"code"`
	Title  string   `json:"title"`
	Career []Career `json:"career"`
}

// MNMCareerClustersResponse is returned by MNMCareerClusters/MPPCareerClusters.
type MNMCareerClustersResponse struct {
	CareerCluster []CareerCluster `json:"career_cluster"`
}

// MNMCareerClusterResponse is returned by MNMCareerCluster/MPPCareerCluster.
type MNMCareerClusterResponse struct {
	Code   string   `json:"code"`
	Title  string   `json:"title"`
	Career []Career `json:"career"`
}

// MNMJobPreparationResponse is returned by MNMJobPreparation/MPPJobPreparation.
type MNMJobPreparationResponse struct {
	JobZone []MNMJobZoneItem `json:"job_zone"`
}

// MNMJobPreparationZoneResponse is returned by MNMJobPreparationZone/MPPJobPreparationZone.
type MNMJobPreparationZoneResponse struct {
	Zone   int      `json:"zone"`
	Career []Career `json:"career"`
}

// ─── internal helpers ────────────────────────────────────────────────────────

func (c *Client) mnmSearch(ctx context.Context, prefix, keyword string, page PageParams) (*CareerList, error) {
	q := url.Values{"keyword": {keyword}}
	page.apply(q)
	var out CareerList
	return &out, c.do(ctx, prefix+"/search", q, &out)
}

func (c *Client) mnmCareer(ctx context.Context, prefix, code string) (*MNMCareerDetail, error) {
	var out MNMCareerDetail
	return &out, c.do(ctx, prefix+"/careers/"+code+"/", nil, &out)
}

func (c *Client) mnmCareerResource(ctx context.Context, prefix, code, resource string, out any) error {
	return c.do(ctx, prefix+"/careers/"+code+"/"+resource, nil, out)
}

func (c *Client) mnmInterestProfilerAnswers(ctx context.Context, prefix, endpoint string, answers []string, page PageParams, out any) error {
	q := url.Values{"answers": {strings.Join(answers, ",")}}
	page.apply(q)
	return c.do(ctx, prefix+"/interestprofiler/"+endpoint, q, out)
}

// ─── MNM (English) ───────────────────────────────────────────────────────────

// MNMSearch searches My Next Move careers by keyword.
func (c *Client) MNMSearch(ctx context.Context, keyword string, page PageParams) (*CareerList, error) {
	return c.mnmSearch(ctx, "mnm", keyword, page)
}

// MNMCareer returns the full career record for the given O*NET-SOC code.
func (c *Client) MNMCareer(ctx context.Context, code string) (*MNMCareerDetail, error) {
	return c.mnmCareer(ctx, "mnm", code)
}

// MNMCareerKnowledge returns knowledge areas for the given career.
func (c *Client) MNMCareerKnowledge(ctx context.Context, code string) (*OccupationElements, error) {
	var out OccupationElements
	return &out, c.mnmCareerResource(ctx, "mnm", code, "knowledge", &out)
}

// MNMCareerSkills returns skills for the given career.
func (c *Client) MNMCareerSkills(ctx context.Context, code string) (*OccupationElements, error) {
	var out OccupationElements
	return &out, c.mnmCareerResource(ctx, "mnm", code, "skills", &out)
}

// MNMCareerAbilities returns abilities for the given career.
func (c *Client) MNMCareerAbilities(ctx context.Context, code string) (*OccupationElements, error) {
	var out OccupationElements
	return &out, c.mnmCareerResource(ctx, "mnm", code, "abilities", &out)
}

// MNMCareerPersonality returns personality traits for the given career.
func (c *Client) MNMCareerPersonality(ctx context.Context, code string) (*OccupationElements, error) {
	var out OccupationElements
	return &out, c.mnmCareerResource(ctx, "mnm", code, "personality", &out)
}

// MNMCareerTechnology returns technology categories for the given career.
func (c *Client) MNMCareerTechnology(ctx context.Context, code string) (*OccupationTechnologySkills, error) {
	var out OccupationTechnologySkills
	return &out, c.mnmCareerResource(ctx, "mnm", code, "technology", &out)
}

// MNMCareerEducation returns education requirements for the given career.
func (c *Client) MNMCareerEducation(ctx context.Context, code string) (*MNMCareerEducationResponse, error) {
	var out MNMCareerEducationResponse
	return &out, c.mnmCareerResource(ctx, "mnm", code, "education", &out)
}

// MNMCareerJobOutlook returns job outlook data for the given career.
func (c *Client) MNMCareerJobOutlook(ctx context.Context, code string) (*MNMCareerJobOutlookResponse, error) {
	var out MNMCareerJobOutlookResponse
	return &out, c.mnmCareerResource(ctx, "mnm", code, "job_outlook", &out)
}

// MNMCareerState returns state-level employment data for the given career.
func (c *Client) MNMCareerState(ctx context.Context, code string) (*MNMCareerStateResponse, error) {
	var out MNMCareerStateResponse
	return &out, c.mnmCareerResource(ctx, "mnm", code, "state", &out)
}

// MNMIndustries returns the list of industries in My Next Move.
func (c *Client) MNMIndustries(ctx context.Context) (*MNMIndustriesResponse, error) {
	var out MNMIndustriesResponse
	return &out, c.do(ctx, "mnm/industries/", nil, &out)
}

// MNMIndustryMostPeople returns the most common careers in the given industry.
func (c *Client) MNMIndustryMostPeople(ctx context.Context, code string) (*MNMIndustryCareersResponse, error) {
	var out MNMIndustryCareersResponse
	return &out, c.do(ctx, "mnm/industries/"+code+"/most_people", nil, &out)
}

// MNMIndustrySomePeople returns additional careers in the given industry.
func (c *Client) MNMIndustrySomePeople(ctx context.Context, code string) (*MNMIndustryCareersResponse, error) {
	var out MNMIndustryCareersResponse
	return &out, c.do(ctx, "mnm/industries/"+code+"/some_people", nil, &out)
}

// MNMInterestProfilerQuestions returns the full 60-question interest profiler.
func (c *Client) MNMInterestProfilerQuestions(ctx context.Context) (*InterestProfilerQuestionsResponse, error) {
	var out InterestProfilerQuestionsResponse
	return &out, c.do(ctx, "mnm/interestprofiler/questions", nil, &out)
}

// MNMInterestProfilerQuestions30 returns the short 30-question interest profiler.
func (c *Client) MNMInterestProfilerQuestions30(ctx context.Context) (*InterestProfilerQuestionsResponse, error) {
	var out InterestProfilerQuestionsResponse
	return &out, c.do(ctx, "mnm/interestprofiler/questions_30", nil, &out)
}

// MNMInterestProfilerResults returns RIASEC scores for the given answers.
// answers is a slice of response values (one per question, e.g. "1" through "5").
func (c *Client) MNMInterestProfilerResults(ctx context.Context, answers []string) (*InterestProfilerResultsResponse, error) {
	var out InterestProfilerResultsResponse
	return &out, c.mnmInterestProfilerAnswers(ctx, "mnm", "results", answers, PageParams{}, &out)
}

// MNMInterestProfilerCareers returns careers matching the given interest profiler answers.
func (c *Client) MNMInterestProfilerCareers(ctx context.Context, answers []string, page PageParams) (*CareerList, error) {
	var out CareerList
	return &out, c.mnmInterestProfilerAnswers(ctx, "mnm", "careers", answers, page, &out)
}

// MNMInterests returns the list of RIASEC interest areas.
func (c *Client) MNMInterests(ctx context.Context) (*MNMInterestsResponse, error) {
	var out MNMInterestsResponse
	return &out, c.do(ctx, "mnm/interestprofiler/interests/", nil, &out)
}

// MNMInterestCareers returns careers in the given RIASEC interest area.
func (c *Client) MNMInterestCareers(ctx context.Context, code string) (*MNMInterestCareersResponse, error) {
	var out MNMInterestCareersResponse
	return &out, c.do(ctx, "mnm/interestprofiler/interests/"+code, nil, &out)
}

// MNMJobZones returns the list of My Next Move job zones.
func (c *Client) MNMJobZones(ctx context.Context) (*MNMJobZonesResponse, error) {
	var out MNMJobZonesResponse
	return &out, c.do(ctx, "mnm/interestprofiler/job_zones", nil, &out)
}

// MNMBrightOutlookCategories returns the bright outlook categories in My Next Move.
func (c *Client) MNMBrightOutlookCategories(ctx context.Context) (*MNMBrightOutlookCategoriesResponse, error) {
	var out MNMBrightOutlookCategoriesResponse
	return &out, c.do(ctx, "mnm/bright_outlook/", nil, &out)
}

// MNMBrightOutlookCareers returns careers in the given bright outlook category.
func (c *Client) MNMBrightOutlookCareers(ctx context.Context, category string) (*MNMBrightOutlookCareersResponse, error) {
	var out MNMBrightOutlookCareersResponse
	return &out, c.do(ctx, "mnm/bright_outlook/"+category, nil, &out)
}

// MNMCareerClusters returns the list of career clusters in My Next Move.
func (c *Client) MNMCareerClusters(ctx context.Context) (*MNMCareerClustersResponse, error) {
	var out MNMCareerClustersResponse
	return &out, c.do(ctx, "mnm/career_clusters/", nil, &out)
}

// MNMCareerCluster returns careers in the given career cluster.
func (c *Client) MNMCareerCluster(ctx context.Context, code string) (*MNMCareerClusterResponse, error) {
	var out MNMCareerClusterResponse
	return &out, c.do(ctx, "mnm/career_clusters/"+code, nil, &out)
}

// MNMJobPreparation returns all job zones with career preparation info.
func (c *Client) MNMJobPreparation(ctx context.Context) (*MNMJobPreparationResponse, error) {
	var out MNMJobPreparationResponse
	return &out, c.do(ctx, "mnm/job_preparation/", nil, &out)
}

// MNMJobPreparationZone returns careers in the given job preparation zone (1–5).
func (c *Client) MNMJobPreparationZone(ctx context.Context, zone int) (*MNMJobPreparationZoneResponse, error) {
	var out MNMJobPreparationZoneResponse
	return &out, c.do(ctx, "mnm/job_preparation/"+strconv.Itoa(zone), nil, &out)
}

// ─── MPP (Spanish — Mi Próximo Paso) ─────────────────────────────────────────
// MPP mirrors MNM exactly except it has no personality or technology endpoints.

// MPPSearch searches Mi Próximo Paso careers by keyword.
func (c *Client) MPPSearch(ctx context.Context, keyword string, page PageParams) (*CareerList, error) {
	return c.mnmSearch(ctx, "mpp", keyword, page)
}

// MPPCareer returns the full career record for the given code.
func (c *Client) MPPCareer(ctx context.Context, code string) (*MNMCareerDetail, error) {
	return c.mnmCareer(ctx, "mpp", code)
}

// MPPCareerKnowledge returns knowledge areas for the given career.
func (c *Client) MPPCareerKnowledge(ctx context.Context, code string) (*OccupationElements, error) {
	var out OccupationElements
	return &out, c.mnmCareerResource(ctx, "mpp", code, "knowledge", &out)
}

// MPPCareerSkills returns skills for the given career.
func (c *Client) MPPCareerSkills(ctx context.Context, code string) (*OccupationElements, error) {
	var out OccupationElements
	return &out, c.mnmCareerResource(ctx, "mpp", code, "skills", &out)
}

// MPPCareerAbilities returns abilities for the given career.
func (c *Client) MPPCareerAbilities(ctx context.Context, code string) (*OccupationElements, error) {
	var out OccupationElements
	return &out, c.mnmCareerResource(ctx, "mpp", code, "abilities", &out)
}

// MPPCareerEducation returns education requirements for the given career.
func (c *Client) MPPCareerEducation(ctx context.Context, code string) (*MNMCareerEducationResponse, error) {
	var out MNMCareerEducationResponse
	return &out, c.mnmCareerResource(ctx, "mpp", code, "education", &out)
}

// MPPCareerJobOutlook returns job outlook data for the given career.
func (c *Client) MPPCareerJobOutlook(ctx context.Context, code string) (*MNMCareerJobOutlookResponse, error) {
	var out MNMCareerJobOutlookResponse
	return &out, c.mnmCareerResource(ctx, "mpp", code, "job_outlook", &out)
}

// MPPCareerState returns state-level employment data for the given career.
func (c *Client) MPPCareerState(ctx context.Context, code string) (*MNMCareerStateResponse, error) {
	var out MNMCareerStateResponse
	return &out, c.mnmCareerResource(ctx, "mpp", code, "state", &out)
}

// MPPIndustries returns the list of industries in Mi Próximo Paso.
func (c *Client) MPPIndustries(ctx context.Context) (*MNMIndustriesResponse, error) {
	var out MNMIndustriesResponse
	return &out, c.do(ctx, "mpp/industries/", nil, &out)
}

// MPPIndustryMostPeople returns the most common careers in the given industry.
func (c *Client) MPPIndustryMostPeople(ctx context.Context, code string) (*MNMIndustryCareersResponse, error) {
	var out MNMIndustryCareersResponse
	return &out, c.do(ctx, "mpp/industries/"+code+"/most_people", nil, &out)
}

// MPPIndustrySomePeople returns additional careers in the given industry.
func (c *Client) MPPIndustrySomePeople(ctx context.Context, code string) (*MNMIndustryCareersResponse, error) {
	var out MNMIndustryCareersResponse
	return &out, c.do(ctx, "mpp/industries/"+code+"/some_people", nil, &out)
}

// MPPInterestProfilerQuestions returns the full 60-question interest profiler (Spanish).
func (c *Client) MPPInterestProfilerQuestions(ctx context.Context) (*InterestProfilerQuestionsResponse, error) {
	var out InterestProfilerQuestionsResponse
	return &out, c.do(ctx, "mpp/interestprofiler/questions", nil, &out)
}

// MPPInterestProfilerQuestions30 returns the short 30-question interest profiler (Spanish).
func (c *Client) MPPInterestProfilerQuestions30(ctx context.Context) (*InterestProfilerQuestionsResponse, error) {
	var out InterestProfilerQuestionsResponse
	return &out, c.do(ctx, "mpp/interestprofiler/questions_30", nil, &out)
}

// MPPInterestProfilerResults returns RIASEC scores for the given answers (Spanish).
func (c *Client) MPPInterestProfilerResults(ctx context.Context, answers []string) (*InterestProfilerResultsResponse, error) {
	var out InterestProfilerResultsResponse
	return &out, c.mnmInterestProfilerAnswers(ctx, "mpp", "results", answers, PageParams{}, &out)
}

// MPPInterestProfilerCareers returns careers matching the given interest profiler answers (Spanish).
func (c *Client) MPPInterestProfilerCareers(ctx context.Context, answers []string, page PageParams) (*CareerList, error) {
	var out CareerList
	return &out, c.mnmInterestProfilerAnswers(ctx, "mpp", "careers", answers, page, &out)
}

// MPPInterests returns the list of RIASEC interest areas (Spanish).
func (c *Client) MPPInterests(ctx context.Context) (*MNMInterestsResponse, error) {
	var out MNMInterestsResponse
	return &out, c.do(ctx, "mpp/interestprofiler/interests/", nil, &out)
}

// MPPInterestCareers returns careers in the given RIASEC interest area (Spanish).
func (c *Client) MPPInterestCareers(ctx context.Context, code string) (*MNMInterestCareersResponse, error) {
	var out MNMInterestCareersResponse
	return &out, c.do(ctx, "mpp/interestprofiler/interests/"+code, nil, &out)
}

// MPPJobZones returns the list of Mi Próximo Paso job zones.
func (c *Client) MPPJobZones(ctx context.Context) (*MNMJobZonesResponse, error) {
	var out MNMJobZonesResponse
	return &out, c.do(ctx, "mpp/interestprofiler/job_zones", nil, &out)
}

// MPPBrightOutlookCategories returns the bright outlook categories in Mi Próximo Paso.
func (c *Client) MPPBrightOutlookCategories(ctx context.Context) (*MNMBrightOutlookCategoriesResponse, error) {
	var out MNMBrightOutlookCategoriesResponse
	return &out, c.do(ctx, "mpp/bright_outlook/", nil, &out)
}

// MPPBrightOutlookCareers returns careers in the given bright outlook category (Spanish).
func (c *Client) MPPBrightOutlookCareers(ctx context.Context, category string) (*MNMBrightOutlookCareersResponse, error) {
	var out MNMBrightOutlookCareersResponse
	return &out, c.do(ctx, "mpp/bright_outlook/"+category, nil, &out)
}

// MPPCareerClusters returns the list of career clusters in Mi Próximo Paso.
func (c *Client) MPPCareerClusters(ctx context.Context) (*MNMCareerClustersResponse, error) {
	var out MNMCareerClustersResponse
	return &out, c.do(ctx, "mpp/career_clusters/", nil, &out)
}

// MPPCareerCluster returns careers in the given career cluster (Spanish).
func (c *Client) MPPCareerCluster(ctx context.Context, code string) (*MNMCareerClusterResponse, error) {
	var out MNMCareerClusterResponse
	return &out, c.do(ctx, "mpp/career_clusters/"+code, nil, &out)
}

// MPPJobPreparation returns all job zones with career preparation info (Spanish).
func (c *Client) MPPJobPreparation(ctx context.Context) (*MNMJobPreparationResponse, error) {
	var out MNMJobPreparationResponse
	return &out, c.do(ctx, "mpp/job_preparation/", nil, &out)
}

// MPPJobPreparationZone returns careers in the given job preparation zone (1–5, Spanish).
func (c *Client) MPPJobPreparationZone(ctx context.Context, zone int) (*MNMJobPreparationZoneResponse, error) {
	var out MNMJobPreparationZoneResponse
	return &out, c.do(ctx, "mpp/job_preparation/"+strconv.Itoa(zone), nil, &out)
}

// ─── Veterans ────────────────────────────────────────────────────────────────

// VeteransOccupationList is a paginated list of occupations returned by VeteransSearch.
type VeteransOccupationList struct {
	PageMeta
	Occupation []Occupation `json:"occupation"`
}

// VeteransSearch searches occupations by military keyword or code.
func (c *Client) VeteransSearch(ctx context.Context, keyword string, page PageParams) (*VeteransOccupationList, error) {
	q := url.Values{"keyword": {keyword}}
	page.apply(q)
	var out VeteransOccupationList
	return &out, c.do(ctx, "veterans/military", q, &out)
}
