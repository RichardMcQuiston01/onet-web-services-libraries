# Changelog

All notable changes to this project will be documented in this file.

The format follows [Keep a Changelog](https://keepachangelog.com/en/1.0.0/).
This project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

---

## [Unreleased]

## [0.1.0] — 2026-06-23

Initial release of the Go client library for the O*NET Web Services API.

### Added

**Core**
- `Client` struct with `NewClient(apiKey string, opts ...Option) *Client`
- `WithBaseURL(u string) Option` — override the default API base URL
- `WithHTTPClient(hc *http.Client) Option` — supply a custom HTTP client
- `APIError` type exposing `StatusCode` and `Message` for all non-2xx responses
- `PageParams` / `PageMeta` for uniform pagination across all list endpoints

**About**
- `About` — API version and O*NET release date

**O\*NET OnLine — Search & Occupation**
- `OnlineSearch` — keyword search with pagination
- `OnlineOccupations` — paginated list of all occupations
- `OnlineOccupation` — full occupation record by O*NET-SOC code

**O\*NET OnLine — Summary (16 topics)**
`OnlineSummaryTasks`, `OnlineSummaryTechnologySkills`, `OnlineSummaryWorkActivities`,
`OnlineSummaryDetailedWorkActivities`, `OnlineSummaryWorkContext`, `OnlineSummaryJobZone`,
`OnlineSummaryApprenticeship`, `OnlineSummarySkills`, `OnlineSummaryKnowledge`,
`OnlineSummaryEducation`, `OnlineSummaryAbilities`, `OnlineSummaryInterests`,
`OnlineSummaryWorkStyles`, `OnlineSummaryRelatedOccupations`,
`OnlineSummaryProfessionalAssociations`, `OnlineSummaryMilitaryCareerSummaries`

**O\*NET OnLine — Details (16 topics)**
Same topics as Summary, returning richer statistical data (standard errors, confidence intervals):
`OnlineDetailsTasks` … `OnlineDetailsMilitaryCareerSummaries`

**O\*NET OnLine — Browse**
- Bright Outlook: `OnlineBrightOutlookCategories`, `OnlineBrightOutlookCategory`, `OnlineBrightOutlookAll`
- Career Clusters: `OnlineCareerClusters`, `OnlineCareerCluster`, `OnlineCareerSubCluster`, `OnlineCareerClustersAll`
- Hot Technologies: `OnlineHotTechnologies`, `OnlineHotTechnologyOccupations`
- Industries: `OnlineIndustries`, `OnlineIndustry`, `OnlineIndustriesAll`
- Job Families: `OnlineJobFamilies`, `OnlineJobFamily`, `OnlineJobFamiliesAll`
- Job Zones: `OnlineJobZones`, `OnlineJobZone`, `OnlineJobZonesAll`
- STEM: `OnlineSTEMOccupationTypes`, `OnlineSTEMOccupationType`, `OnlineSTEMOccupationsAll`

**O\*NET OnLine — Crosswalks**
`OnlineCrosswalkMilitary`, `OnlineCrosswalkEducation`, `OnlineCrosswalkOccupationHandbook`,
`OnlineCrosswalkSOC`, `OnlineCrosswalkDOT`, `OnlineCrosswalkRAPIDS`, `OnlineCrosswalkESCO`

**O\*NET OnLine — Finders**
- Job Duties: `OnlineJobDutiesSearch`, `OnlineJobDutiesTasks`, `OnlineJobDutiesResults`
- Associations: `OnlineAssociationsSearch`, `OnlineAssociations`, `OnlineAssociationsAll`
- Related Activities: `OnlineRelatedActivitiesSearch`, `OnlineRelatedActivities`, `OnlineRelatedActivitiesResults`
- Soft Skills: `OnlineSoftSkills`, `OnlineSoftSkillsResults`
- Technology: `OnlineTechnologyExamplesSearch`, `OnlineTechnologyExample`, `OnlineTechnologyCategoriesSearch`, `OnlineTechnologyCategory`

**O\*NET OnLine — Data Browsers (8 attributes)**
List all elements and look up occupations by element ID for: Abilities, Interests, Knowledge,
Basic Skills, Cross-Functional Skills, Work Activities, Work Context, Work Styles

**My Next Move (English) — `/mnm/`**
`MNMSearch`, `MNMCareer`, `MNMCareerKnowledge`, `MNMCareerSkills`, `MNMCareerAbilities`,
`MNMCareerPersonality`, `MNMCareerTechnology`, `MNMCareerEducation`, `MNMCareerJobOutlook`,
`MNMCareerState`, `MNMIndustries`, `MNMIndustryMostPeople`, `MNMIndustrySomePeople`,
`MNMInterestProfilerQuestions`, `MNMInterestProfilerQuestions30`,
`MNMInterestProfilerResults`, `MNMInterestProfilerCareers`,
`MNMInterests`, `MNMInterestCareers`, `MNMJobZones`,
`MNMBrightOutlookCategories`, `MNMBrightOutlookCareers`,
`MNMCareerClusters`, `MNMCareerCluster`, `MNMJobPreparation`, `MNMJobPreparationZone`

**Mi Próximo Paso (Spanish) — `/mpp/`**
Same as MNM (minus personality and technology endpoints): `MPPSearch`, `MPPCareer`,
`MPPCareerKnowledge`, `MPPCareerSkills`, `MPPCareerAbilities`, `MPPCareerEducation`,
`MPPCareerJobOutlook`, `MPPCareerState`, `MPPIndustries`, `MPPIndustryMostPeople`,
`MPPIndustrySomePeople`, `MPPInterestProfilerQuestions`, `MPPInterestProfilerQuestions30`,
`MPPInterestProfilerResults`, `MPPInterestProfilerCareers`,
`MPPInterests`, `MPPInterestCareers`, `MPPJobZones`,
`MPPBrightOutlookCategories`, `MPPBrightOutlookCareers`,
`MPPCareerClusters`, `MPPCareerCluster`, `MPPJobPreparation`, `MPPJobPreparationZone`

**Veterans — `/veterans/`**
- `VeteransSearch` — search occupations by military keyword

**Taxonomy Crosswalks**
`TaxonomyFrom2010`, `TaxonomyTo2010`, `TaxonomyFrom2019`, `TaxonomyTo2019`,
`Taxonomy2010To2019`, `Taxonomy2019To2010`

**Database**
- `DatabaseTables` — list all O*NET data tables
- `DatabaseTableInfo` — column definitions for a table
- `DatabaseRows` — query rows with typed `FilterParam` (8 operators) and `SortParam`

**Project**
- `go.mod` — module `github.com/RichardMcQuiston01/onet-web-services-go`, Go 1.21, zero external dependencies
- `.gitignore` for Go projects
- `CLAUDE.md` with build commands and architecture reference

[Unreleased]: https://github.com/RichardMcQuiston01/onet-web-services-go/compare/v0.1.0...HEAD
[0.1.0]: https://github.com/RichardMcQuiston01/onet-web-services-go/releases/tag/v0.1.0
