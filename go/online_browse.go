package onet

import (
	"context"
	"net/url"
	"strconv"
)

// BrightOutlookCategory is a bright outlook classification category.
type BrightOutlookCategory struct {
	Value string `json:"value"`
	Title string `json:"title"`
}

// BrightOutlookCategoriesResponse is returned by OnlineBrightOutlookCategories.
type BrightOutlookCategoriesResponse struct {
	Category []BrightOutlookCategory `json:"category"`
}

// BrightOutlookCategoryResponse is returned by OnlineBrightOutlookCategory.
type BrightOutlookCategoryResponse struct {
	Category   string       `json:"category"`
	Occupation []Occupation `json:"occupation"`
}

// OnlineBrightOutlookCategories returns the list of bright outlook categories.
func (c *Client) OnlineBrightOutlookCategories(ctx context.Context) (*BrightOutlookCategoriesResponse, error) {
	var out BrightOutlookCategoriesResponse
	return &out, c.do(ctx, "online/bright_outlook/", nil, &out)
}

// OnlineBrightOutlookCategory returns occupations in the given bright outlook category.
func (c *Client) OnlineBrightOutlookCategory(ctx context.Context, category string) (*BrightOutlookCategoryResponse, error) {
	var out BrightOutlookCategoryResponse
	return &out, c.do(ctx, "online/bright_outlook/"+category, nil, &out)
}

// OnlineBrightOutlookAll returns a paginated list of all bright outlook occupations.
func (c *Client) OnlineBrightOutlookAll(ctx context.Context, page PageParams) (*OccupationList, error) {
	q := url.Values{}
	page.apply(q)
	var out OccupationList
	return &out, c.do(ctx, "online/bright_outlook/all", q, &out)
}

// CareerCluster is a career and technical education cluster.
type CareerCluster struct {
	Code  string `json:"code"`
	Title string `json:"title"`
}

// CareerClustersResponse is returned by OnlineCareerClusters.
type CareerClustersResponse struct {
	CareerCluster []CareerCluster `json:"career_cluster"`
}

// CareerClusterResponse is returned by OnlineCareerCluster.
type CareerClusterResponse struct {
	CareerCluster CareerCluster `json:"career_cluster"`
	Occupation    []Occupation  `json:"occupation"`
}

// SubCluster is a sub-cluster within a career cluster.
type SubCluster struct {
	Code  string `json:"code"`
	Title string `json:"title"`
}

// SubClusterResponse is returned by OnlineCareerSubCluster.
type SubClusterResponse struct {
	SubCluster SubCluster   `json:"sub_cluster"`
	Occupation []Occupation `json:"occupation"`
}

// OnlineCareerClusters returns the list of career and technical education clusters.
func (c *Client) OnlineCareerClusters(ctx context.Context) (*CareerClustersResponse, error) {
	var out CareerClustersResponse
	return &out, c.do(ctx, "online/career_clusters/", nil, &out)
}

// OnlineCareerCluster returns occupations in the given career cluster.
func (c *Client) OnlineCareerCluster(ctx context.Context, code string) (*CareerClusterResponse, error) {
	var out CareerClusterResponse
	return &out, c.do(ctx, "online/career_clusters/"+code, nil, &out)
}

// OnlineCareerSubCluster returns occupations in the given career sub-cluster.
func (c *Client) OnlineCareerSubCluster(ctx context.Context, code string) (*SubClusterResponse, error) {
	var out SubClusterResponse
	return &out, c.do(ctx, "online/career_clusters/sub_clusters/"+code, nil, &out)
}

// OnlineCareerClustersAll returns a paginated list of all career cluster occupations.
func (c *Client) OnlineCareerClustersAll(ctx context.Context, page PageParams) (*OccupationList, error) {
	q := url.Values{}
	page.apply(q)
	var out OccupationList
	return &out, c.do(ctx, "online/career_clusters/all", q, &out)
}

// HotTechnology is a technology tool or software title tracked as in-demand.
type HotTechnology struct {
	Title string `json:"title"`
}

// HotTechnologiesResponse is returned by OnlineHotTechnologies.
type HotTechnologiesResponse struct {
	Technology []HotTechnology `json:"technology"`
}

// HotTechnologyResponse is returned by OnlineHotTechnologyOccupations.
type HotTechnologyResponse struct {
	Technology HotTechnology `json:"technology"`
	Occupation []Occupation  `json:"occupation"`
}

// OnlineHotTechnologies returns the list of hot technology titles.
func (c *Client) OnlineHotTechnologies(ctx context.Context) (*HotTechnologiesResponse, error) {
	var out HotTechnologiesResponse
	return &out, c.do(ctx, "online/hot_technology/", nil, &out)
}

// OnlineHotTechnologyOccupations returns occupations that use the given hot technology.
func (c *Client) OnlineHotTechnologyOccupations(ctx context.Context, title string) (*HotTechnologyResponse, error) {
	var out HotTechnologyResponse
	return &out, c.do(ctx, "online/hot_technology/"+title, nil, &out)
}

// Industry is an industry classification used to group occupations.
type Industry struct {
	Code  string `json:"code"`
	Title string `json:"title"`
}

// IndustriesResponse is returned by OnlineIndustries.
type IndustriesResponse struct {
	Industry []Industry `json:"industry"`
}

// IndustryResponse is returned by OnlineIndustry.
type IndustryResponse struct {
	Industry   Industry     `json:"industry"`
	Occupation []Occupation `json:"occupation"`
}

// OnlineIndustries returns the list of industries.
func (c *Client) OnlineIndustries(ctx context.Context) (*IndustriesResponse, error) {
	var out IndustriesResponse
	return &out, c.do(ctx, "online/industries/", nil, &out)
}

// OnlineIndustry returns occupations in the given industry.
func (c *Client) OnlineIndustry(ctx context.Context, code string) (*IndustryResponse, error) {
	var out IndustryResponse
	return &out, c.do(ctx, "online/industries/"+code, nil, &out)
}

// OnlineIndustriesAll returns a paginated list of all industry occupations.
func (c *Client) OnlineIndustriesAll(ctx context.Context, page PageParams) (*OccupationList, error) {
	q := url.Values{}
	page.apply(q)
	var out OccupationList
	return &out, c.do(ctx, "online/industries/all", q, &out)
}

// JobFamily is an O*NET job family grouping related occupations.
type JobFamily struct {
	Code  string `json:"code"`
	Title string `json:"title"`
}

// JobFamiliesResponse is returned by OnlineJobFamilies.
type JobFamiliesResponse struct {
	JobFamily []JobFamily `json:"job_family"`
}

// JobFamilyResponse is returned by OnlineJobFamily.
type JobFamilyResponse struct {
	JobFamily  JobFamily    `json:"job_family"`
	Occupation []Occupation `json:"occupation"`
}

// OnlineJobFamilies returns the list of O*NET job families.
func (c *Client) OnlineJobFamilies(ctx context.Context) (*JobFamiliesResponse, error) {
	var out JobFamiliesResponse
	return &out, c.do(ctx, "online/job_families/", nil, &out)
}

// OnlineJobFamily returns occupations in the given job family.
func (c *Client) OnlineJobFamily(ctx context.Context, code string) (*JobFamilyResponse, error) {
	var out JobFamilyResponse
	return &out, c.do(ctx, "online/job_families/"+code, nil, &out)
}

// OnlineJobFamiliesAll returns a paginated list of all job family occupations.
func (c *Client) OnlineJobFamiliesAll(ctx context.Context, page PageParams) (*OccupationList, error) {
	q := url.Values{}
	page.apply(q)
	var out OccupationList
	return &out, c.do(ctx, "online/job_families/all", q, &out)
}

// JobZoneListItem is a job zone entry in the zones list.
type JobZoneListItem struct {
	Number int    `json:"number"`
	Title  string `json:"title"`
}

// JobZonesListResponse is returned by OnlineJobZones.
type JobZonesListResponse struct {
	JobZone []JobZoneListItem `json:"job_zone"`
}

// JobZoneOccupationsResponse is returned by OnlineJobZone.
type JobZoneOccupationsResponse struct {
	JobZone    JobZoneListItem `json:"job_zone"`
	Occupation []Occupation    `json:"occupation"`
}

// OnlineJobZones returns the list of O*NET job zones (1–5).
func (c *Client) OnlineJobZones(ctx context.Context) (*JobZonesListResponse, error) {
	var out JobZonesListResponse
	return &out, c.do(ctx, "online/job_zones/", nil, &out)
}

// OnlineJobZone returns occupations in the given job zone (1–5).
func (c *Client) OnlineJobZone(ctx context.Context, number int) (*JobZoneOccupationsResponse, error) {
	var out JobZoneOccupationsResponse
	return &out, c.do(ctx, "online/job_zones/"+strconv.Itoa(number), nil, &out)
}

// OnlineJobZonesAll returns a paginated list of all occupations across job zones.
func (c *Client) OnlineJobZonesAll(ctx context.Context, page PageParams) (*OccupationList, error) {
	q := url.Values{}
	page.apply(q)
	var out OccupationList
	return &out, c.do(ctx, "online/job_zones/all", q, &out)
}

// STEMType is a STEM occupation type classification.
type STEMType struct {
	Value string `json:"value"`
	Title string `json:"title"`
}

// STEMTypesResponse is returned by OnlineSTEMOccupationTypes.
type STEMTypesResponse struct {
	STEMType []STEMType `json:"stem_type"`
}

// STEMTypeResponse is returned by OnlineSTEMOccupationType.
type STEMTypeResponse struct {
	STEMType   STEMType     `json:"stem_type"`
	Occupation []Occupation `json:"occupation"`
}

// OnlineSTEMOccupationTypes returns the list of STEM occupation type classifications.
func (c *Client) OnlineSTEMOccupationTypes(ctx context.Context) (*STEMTypesResponse, error) {
	var out STEMTypesResponse
	return &out, c.do(ctx, "online/stem_occupations/", nil, &out)
}

// OnlineSTEMOccupationType returns occupations of the given STEM type.
func (c *Client) OnlineSTEMOccupationType(ctx context.Context, stemType string) (*STEMTypeResponse, error) {
	var out STEMTypeResponse
	return &out, c.do(ctx, "online/stem_occupations/"+stemType, nil, &out)
}

// OnlineSTEMOccupationsAll returns a paginated list of all STEM occupations.
func (c *Client) OnlineSTEMOccupationsAll(ctx context.Context, page PageParams) (*OccupationList, error) {
	q := url.Values{}
	page.apply(q)
	var out OccupationList
	return &out, c.do(ctx, "online/stem_occupations/all", q, &out)
}
