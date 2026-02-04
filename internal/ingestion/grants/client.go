package grants

import "context"

// GrantsClient defines a minimal client for Grants.gov-like search and detail lookups.
type GrantsClient interface {
	SearchOpportunities(ctx context.Context, query SearchQuery) ([]byte, error)
	FetchOpportunityDetails(ctx context.Context, opportunityID string) ([]byte, error)
}

// SearchQuery captures typical filters for opportunity searches.
type SearchQuery struct {
	Keyword           string
	Agency            string
	Opportunity       string
	FundingInstrument string
}
