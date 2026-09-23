# onet-web-services-go

A Go client library for the [O*NET Web Services API](https://services.onetcenter.org/reference/start/overview).

## Overview

O*NET Web Services provides occupational data — skills, tasks, wages, and more — for hundreds of occupations. This library wraps the REST API with a consistent, idiomatic Go interface.

**Design goals:**

- Uniform request and response types across all endpoints
- Single `Error` type that preserves the API's HTTP status and message
- Zero third-party dependencies — uses only `net/http` and `encoding/json`
- Thread-safe; all state is scoped to the `Client` struct

## Requirements

- Go 1.21 or later
- An O*NET Web Services API key — register at [services.onetcenter.org](https://services.onetcenter.org/reference/start/authorization)

## Install

```sh
go get github.com/RichardMcQuiston01/onet-web-services-go
```

## Configuration

Create a client with your API key:

```go
import onet "github.com/RichardMcQuiston01/onet-web-services-go"

client := onet.NewClient("your-api-key")
```

**Options:**

```go
// Override the base URL (e.g. for testing against a mock server)
client := onet.NewClient("your-api-key",
    onet.WithBaseURL("https://your-mock-server/"),
)

// Supply a custom HTTP client (e.g. to set a timeout)
import "net/http"
import "time"

client := onet.NewClient("your-api-key",
    onet.WithHTTPClient(&http.Client{Timeout: 10 * time.Second}),
)
```

## Usage

### Check API version

```go
info, err := client.About(ctx)
// info.APIVersion, info.OnetVersion, info.ReleaseDate
```

### Search occupations

```go
results, err := client.OnlineSearch(ctx, "software developer", onet.PageParams{})
for _, occ := range results.Occupation {
    fmt.Println(occ.Code, occ.Title)
}

// Paginate: fetch results 21–40
results, err = client.OnlineSearch(ctx, "software developer", onet.PageParams{Start: 21, End: 40})
fmt.Printf("%d total results\n", results.Total)
```

### Fetch occupation details

```go
occ, err := client.OnlineOccupation(ctx, "15-1252.00")
fmt.Println(occ.WhatTheyDo)

skills, err := client.OnlineSummarySkills(ctx, "15-1252.00")
for _, s := range skills.Element {
    fmt.Printf("%s: %.1f\n", s.Name, s.Score.Value)
}
```

### Browse by category

```go
// All bright outlook occupations
all, err := client.OnlineBrightOutlookAll(ctx, onet.PageParams{})

// Occupations in a career cluster
cluster, err := client.OnlineCareerCluster(ctx, "11")
for _, occ := range cluster.Occupation {
    fmt.Println(occ.Title)
}
```

### Crosswalk to other classification systems

```go
// Find O*NET occupations matching a military keyword
results, err := client.OnlineCrosswalkMilitary(ctx, "infantry", onet.PageParams{})

// Translate a 2019 taxonomy code to the active code
mapping, err := client.TaxonomyFrom2019(ctx, "15-1252.00")
```

### Query the raw database

```go
rows, err := client.DatabaseRows(ctx, "abilities",
    []onet.FilterParam{
        {Column: "onetsoc_code", Op: onet.FilterEq, Value: "15-1252.00"},
        {Column: "scale_id",     Op: onet.FilterEq, Value: "LV"},
    },
    []onet.SortParam{
        {Column: "data_value", Descending: true},
    },
    onet.PageParams{},
)
```

### Error handling

All errors from API failures are returned as `*onet.APIError`:

```go
results, err := client.OnlineSearch(ctx, "developer", onet.PageParams{})
if err != nil {
    var apiErr *onet.APIError
    if errors.As(err, &apiErr) {
        fmt.Printf("API error %d: %s\n", apiErr.StatusCode, apiErr.Message)
        if apiErr.StatusCode == 429 {
            // Rate limited — back off and retry
        }
    }
}
```

HTTP 429 (rate limit) is returned as an error; the library does not retry automatically.

## Building and Testing

```sh
# Build
go build ./...

# Run all unit tests (no API key needed)
go test ./... -short

# Run integration tests against the live API
ONET_API_KEY=your-api-key go test ./... -run Integration -v

# Static analysis
go vet ./...
```

## API Coverage

| Service | Methods |
|---|---|
| O\*NET OnLine — search & occupation | `OnlineSearch`, `OnlineOccupations`, `OnlineOccupation` |
| O\*NET OnLine — summary (16 topics) | `OnlineSummary{Skills,Knowledge,Abilities,Interests,…}` |
| O\*NET OnLine — details (16 topics) | `OnlineDetails{Skills,Knowledge,Abilities,Interests,…}` |
| O\*NET OnLine — browse | BrightOutlook, CareerClusters, Industries, JobFamilies, JobZones, STEM, HotTechnologies |
| O\*NET OnLine — crosswalks | Military, Education, OccupationHandbook, SOC, DOT, RAPIDS, ESCO |
| O\*NET OnLine — finders | JobDuties, Associations, RelatedActivities, SoftSkills, Technology |
| O\*NET OnLine — data browsers | 8 attributes (abilities, interests, knowledge, skills, work activities/context/styles) |
| My Next Move (English) | `MNM*` — search, career details, interest profiler, industries, clusters |
| Mi Próximo Paso (Spanish) | `MPP*` — same as MNM minus personality and technology |
| Veterans | `VeteransSearch` |
| Taxonomy crosswalks | 2010 ↔ active, 2019 ↔ active, 2010 ↔ 2019 |
| Database | `DatabaseTables`, `DatabaseTableInfo`, `DatabaseRows` |
| About | `About` |

## Support

If this library saved you some reverse-engineering, consider [buying me a coffee](https://donate.stripe.com/00w5kD3Gj1Xo9v7gVOcs800). ☕

## Resources

- [O*NET Web Services Overview](https://services.onetcenter.org/reference/start/overview)
- [API Reference](https://services.onetcenter.org/reference/apis)
- [Authorization](https://services.onetcenter.org/reference/start/authorization)

## License

MIT © 2026 Richard McQuiston — see [LICENSE](LICENSE).
