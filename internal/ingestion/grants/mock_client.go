package grants

import (
	"context"
	"fmt"
)

// MockGrantsClient returns hardcoded JSON payloads that mimic the Grants.gov V2 API schema.
type MockGrantsClient struct{}

// SearchOpportunities returns a static JSON response for a search query.
func (m MockGrantsClient) SearchOpportunities(ctx context.Context, query SearchQuery) ([]byte, error) {
	_ = ctx
	_ = query
	response := `{
  "hits": 2,
  "page": 1,
  "size": 10,
  "opportunities": [
    {
      "opportunityId": "GRANTS-001",
      "opportunityNumber": "HHS-2024-001",
      "opportunityTitle": "Community Health Innovation",
      "agency": "HHS",
      "agencyCode": "HHS",
      "postDate": "2024-01-15",
      "closeDate": "2024-03-01",
      "fundingInstrument": "Grant",
      "fundingCategory": "Health",
      "estimatedTotalProgramFunding": 5000000,
      "expectedNumberOfAwards": 25
    },
    {
      "opportunityId": "GRANTS-002",
      "opportunityNumber": "NSF-2024-ML",
      "opportunityTitle": "AI for Scientific Discovery",
      "agency": "NSF",
      "agencyCode": "NSF",
      "postDate": "2024-02-01",
      "closeDate": "2024-04-15",
      "fundingInstrument": "Cooperative Agreement",
      "fundingCategory": "Science",
      "estimatedTotalProgramFunding": 12000000,
      "expectedNumberOfAwards": 10
    }
  ]
}`
	return []byte(response), nil
}

// FetchOpportunityDetails returns a static JSON response for an opportunity detail lookup.
func (m MockGrantsClient) FetchOpportunityDetails(ctx context.Context, opportunityID string) ([]byte, error) {
	_ = ctx
	switch opportunityID {
	case "GRANTS-001":
		return []byte(`{
  "opportunityId": "GRANTS-001",
  "opportunityNumber": "HHS-2024-001",
  "opportunityTitle": "Community Health Innovation",
  "agency": "HHS",
  "agencyCode": "HHS",
  "postDate": "2024-01-15",
  "closeDate": "2024-03-01",
  "synopsis": "Supports community-driven health initiatives.",
  "eligibility": {
    "eligibleApplicants": ["Nonprofits", "Local Governments"],
    "additionalInfo": "Applicants must serve underserved communities."
  },
  "award": {
    "min": 50000,
    "max": 250000,
    "estimatedTotalProgramFunding": 5000000,
    "expectedNumberOfAwards": 25
  }
}`), nil
	case "GRANTS-002":
		return []byte(`{
  "opportunityId": "GRANTS-002",
  "opportunityNumber": "NSF-2024-ML",
  "opportunityTitle": "AI for Scientific Discovery",
  "agency": "NSF",
  "agencyCode": "NSF",
  "postDate": "2024-02-01",
  "closeDate": "2024-04-15",
  "synopsis": "Funds AI research collaborations with scientific labs.",
  "eligibility": {
    "eligibleApplicants": ["Universities", "Nonprofits"],
    "additionalInfo": "Partnerships encouraged."
  },
  "award": {
    "min": 200000,
    "max": 1500000,
    "estimatedTotalProgramFunding": 12000000,
    "expectedNumberOfAwards": 10
  }
}`), nil
	default:
		return nil, fmt.Errorf("unknown opportunity ID: %s", opportunityID)
	}
}
