# SCANOSS Geo Provenance Service API v2

Provides geographical origin analysis for software components by examining contributor locations, commit patterns, and development activity across different regions.

## GetCountryContributorsByComponents

Retrieves geographical provenance information based on contributor declared locations for software components. Analyzes repository metadata and contributor profiles to determine the geographical distribution of development activity.

### Request Format
See [Common API Types](../common/v2/README.md) for `ComponentsRequest` documentation.

### HTTP Request Example
```bash
curl -X POST 'https://api.scanoss.com/v2/geoprovenance/countries/components' \
  -H 'Content-Type: application/json' \
  -H "X-Api-Key: $SC_API_KEY" \
  -d '{
    "components": [
      {"purl": "pkg:github/scanoss/engine@5.0.0"}
    ]
  }'
```

### gRPC Request Example
```bash
grpcurl -H "X-Api-Key: $SC_API_KEY" \
  -d '{
    "components": [
      {"purl": "pkg:github/scanoss/engine@5.0.0"}
    ]
  }' \
  api.scanoss.com:443 \
  scanoss.api.geoprovenance.v2.GeoProvenance/GetCountryContributorsByComponents
```

### Response Format

The method returns geographical contributor information including:

- `components_locations` array: Contains geo-provenance data for each requested component
- `status` field: Response status indicating success or failure of the request

Each component object contains:
- `purl`: Package URL identifying the component
- `declared_locations`: Locations declared by contributors in their profiles or repository metadata
- `curated_locations`: SCANOSS-curated geographical data based on analysis

### Response Examples

#### Component with Geographic Distribution
```json
{
  "components_locations": [
    {
      "purl": "pkg:github/scanoss/engine@5.0.0",
      "declared_locations": [
        {
          "type": "owner",
          "location": "Barcelona, Spain"
        },
        {
          "type": "contributor",
          "location": "Berlin, Germany"
        }
      ],
      "curated_locations": [
        {
          "country": "Spain",
          "count": 8
        },
        {
          "country": "Germany",
          "count": 3
        },
        {
          "country": "United States",
          "count": 2
        }
      ]
    }
  ],
  "status": {
    "status": "SUCCESS",
    "message": "Geo-provenance successfully retrieved"
  }
}
```

#### Component with Limited Geographic Data
```json
{
  "components_locations": [
    {
      "purl": "pkg:npm/simple-utility@1.0.0",
      "declared_locations": [],
      "curated_locations": [
        {
          "country": "United States",
          "count": 1
        }
      ]
    }
  ],
  "status": {
    "status": "SUCCESS",
    "message": "Geo-provenance successfully retrieved"
  }
}
```

## GetCountryContributorsByComponent

Retrieves geographical provenance information for a single component based on contributor declared locations. Analyzes repository metadata and contributor profiles to determine the geographical distribution of development activity.

### Request Format
See [Common API Types](../common/v2/README.md) for `ComponentRequest` documentation.

### HTTP Request Example
```bash
curl -X GET 'https://api.scanoss.com/v2/geoprovenance/countries/component?purl=pkg:github/scanoss/engine@5.0.0' \
  -H "X-Api-Key: $SC_API_KEY"
```

### gRPC Request Example
```bash
grpcurl -H "X-Api-Key: $SC_API_KEY" \
  -d '{
    "purl": "pkg:github/scanoss/engine@5.0.0"
  }' \
  api.scanoss.com:443 \
  scanoss.api.geoprovenance.v2.GeoProvenance/GetCountryContributorsByComponent
```

### Response Format

The method returns geographical contributor information including:

- `component_locations` object: Contains geo-provenance data for the requested component
- `status` field: Response status indicating success or failure of the request

The component object contains:
- `purl`: Package URL identifying the component
- `declared_locations`: Locations declared by contributors in their profiles or repository metadata
- `curated_locations`: SCANOSS-curated geographical data based on analysis

### Response Examples

#### Component with Geographic Distribution
```json
{
  "component_locations": {
    "purl": "pkg:github/scanoss/engine@5.0.0",
    "declared_locations": [
      {
        "type": "owner",
        "location": "Barcelona, Spain"
      },
      {
        "type": "contributor",
        "location": "Berlin, Germany"
      }
    ],
    "curated_locations": [
      {
        "country": "Spain",
        "count": 8
      },
      {
        "country": "Germany",
        "count": 3
      },
      {
        "country": "United States",
        "count": 2
      }
    ]
  },
  "status": {
    "status": "SUCCESS",
    "message": "Geo-provenance successfully retrieved"
  }
}
```

#### Component with Limited Geographic Data
```json
{
  "component_locations": {
    "purl": "pkg:npm/simple-utility@1.0.0",
    "declared_locations": [],
    "curated_locations": [
      {
        "country": "United States",
        "count": 1
      }
    ]
  },
  "status": {
    "status": "SUCCESS",
    "message": "Geo-provenance successfully retrieved"
  }
}
```

## GetComponentContributors (Deprecated)

**Note: This method is deprecated and will be removed in a future version. Please use GetCountryContributorsByComponents instead.**

Legacy method for analyzing geographical contributor information using PURL-based requests. Provides the same functionality as GetCountryContributorsByComponents but with the older request format.

### Request Format
See [Common API Types](../common/v2/README.md) for `PurlRequest` documentation.

### HTTP Request Example
```bash
curl -X POST 'https://api.scanoss.com/v2/geoprovenance/countries' \
  -H 'Content-Type: application/json' \
  -H "X-Api-Key: $SC_API_KEY" \
  -d '{
    "purls": [
      {"purl": "pkg:github/scanoss/engine@5.0.0"}
    ]
  }'
```

### Migration Guide
To migrate from `GetComponentContributors` to `GetCountryContributorsByComponents`:

1. **Change the endpoint**: `/v2/geoprovenance/countries` � `/v2/geoprovenance/countries/component`
2. **Update request format**: Use `ComponentsRequest` instead of `PurlRequest`
3. **Update response handling**: Use `ComponentsContributorResponse` instead of `ContributorResponse`


## GetOriginByComponents

Retrieves geographical origin information based on contributor commit timing patterns and development activity analysis. This method examines when contributors are most active to infer their likely geographical locations.

### Request Format
See [Common API Types](../common/v2/README.md) for `ComponentsRequest` documentation.

### HTTP Request Example
```bash
curl -X POST 'https://api.scanoss.com/v2/geoprovenance/origin/components' \
  -H 'Content-Type: application/json' \
  -H "X-Api-Key: $SC_API_KEY" \
  -d '{
    "components": [
      {"purl": "pkg:github/scanoss/engine@5.0.0"}
    ]
  }'
```

### gRPC Request Example
```bash
grpcurl -H "X-Api-Key: $SC_API_KEY" \
  -d  '{
    "components": [
      {"purl": "pkg:github/scanoss/engine@5.0.0"}
    ]
  }' \
  api.scanoss.com:443 \
  scanoss.api.geoprovenance.v2.GeoProvenance/GetOriginByComponents
```

### Response Format

The method returns commit-time based geographical analysis including:

- `components_locations` array: Contains geographical origin data based on commit patterns
- `status` field: Response status indicating success or failure of the request

Each component location object contains:
- `purl`: Package URL identifying the component
- `locations`: Array of countries with contributor percentages based on commit timing analysis

### Response Examples

#### Component with Origin Analysis
```json
{
  "components_locations": [
    {
      "purl": "pkg:github/scanoss/engine@5.0.0",
      "locations": [
        {
          "name": "ES",
          "percentage": 65.5
        },
        {
          "name": "DE",
          "percentage": 20.3
        },
        {
          "name": "US",
          "percentage": 14.2
        }
      ]
    }
  ],
  "status": {
    "status": "SUCCESS",
    "message": "Geo-provenance origin successfully retrieved"
  }
}
```

#### Component with Single Origin
```json
{
  "components_locations": [
    {
      "purl": "pkg:npm/private-utility@2.1.0",
      "locations": [
        {
          "name": "US",
          "percentage": 100.0
        }
      ]
    }
  ],
  "status": {
    "status": "SUCCESS",
    "message": "Geo-provenance origin successfully retrieved"
  }
}
```


## GetOriginByComponent

Retrieves geographical origin information for a single component based on contributor commit timing patterns and development activity analysis. This method examines when contributors are most active to infer their likely geographical locations.

### Request Format
See [Common API Types](../common/v2/README.md) for `ComponentRequest` documentation.

### HTTP Request Example
```bash
curl -X GET 'https://api.scanoss.com/v2/geoprovenance/origin/component?purl=pkg:github/scanoss/engine@5.0.0' \
  -H "X-Api-Key: $SC_API_KEY"
```

### gRPC Request Example
```bash
grpcurl -H "X-Api-Key: $SC_API_KEY" \
  -d '{
    "purl": "pkg:github/scanoss/engine@5.0.0"
  }' \
  api.scanoss.com:443 \
  scanoss.api.geoprovenance.v2.GeoProvenance/GetOriginByComponent
```

### Response Format

The method returns commit-time based geographical analysis including:

- `component_locations` object: Contains geographical origin data based on commit patterns for the requested component
- `status` field: Response status indicating success or failure of the request

The component location object contains:
- `purl`: Package URL identifying the component
- `locations`: Array of countries with contributor percentages based on commit timing analysis

### Response Examples

#### Component with Origin Analysis
```json
{
  "component_locations": {
    "purl": "pkg:github/scanoss/engine@5.0.0",
    "locations": [
      {
        "name": "ES",
        "percentage": 65.5
      },
      {
        "name": "DE",
        "percentage": 20.3
      },
      {
        "name": "US",
        "percentage": 14.2
      }
    ]
  },
  "status": {
    "status": "SUCCESS",
    "message": "Geo-provenance origin successfully retrieved"
  }
}
```

#### Component with Single Origin
```json
{
  "component_locations": {
    "purl": "pkg:npm/private-utility@2.1.0",
    "locations": [
      {
        "name": "US",
        "percentage": 100.0
      }
    ]
  },
  "status": {
    "status": "SUCCESS",
    "message": "Geo-provenance origin successfully retrieved"
  }
}
```

## GetComponentOrigin (Deprecated)

**Note: This method is deprecated and will be removed in a future version. Please use GetOriginByComponents instead.**

Legacy method for analyzing geographical origin information using PURL-based requests. Provides the same functionality as GetOriginByComponents but with the older request format.

### Request Format
See [Common API Types](../common/v2/README.md) for `PurlRequest` documentation.

### HTTP Request Example
```bash
curl -X POST 'https://api.scanoss.com/v2/geoprovenance/origin' \
  -H 'Content-Type: application/json' \
  -H "X-Api-Key: $SC_API_KEY" \
  -d '{
    "purls": [
      {"purl": "pkg:github/scanoss/engine@5.0.0"}
    ]
  }'
```


### Migration Guide
To migrate from `GetComponentOrigin` to `GetOriginByComponents`:

1. **Change the endpoint**: `/v2/geoprovenance/origin` � `/v2/geoprovenance/origin/components`
2. **Update request format**: Use `ComponentsRequest` instead of `PurlRequest`
3. **Update response handling**: Use `ComponentsOriginResponse` instead of `OriginResponse`

## Echo

Standard service health check endpoint for testing connectivity and API key validation.

### HTTP Request Example
```bash
curl -X POST 'https://api.scanoss.com/v2/geoprovenance/echo' \
  -H 'Content-Type: application/json' \
  -H "X-Api-Key: $SC_API_KEY" \
  -d '{"message": "test"}'
```

### gRPC Request Example
```bash
grpcurl -H "X-Api-Key: $SC_API_KEY" \
  -d '{"message": "test"}' \
  api.scanoss.com:443 \
  scanoss.api.geoprovenance.v2.GeoProvenance/Echo
```

## Analysis Methods

The Geo Provenance service employs two complementary analysis approaches:

### Declared Location Analysis
- Examines contributor profiles and repository metadata
- Identifies locations explicitly declared by project maintainers and contributors

### Commit Timing Analysis
- Analyzes commit timestamps and patterns to infer contributor time zones
- Uses statistical models to estimate geographical distribution
- Provides data-driven insights into actual development activity locations

Both methods work together to provide comprehensive geographical intelligence for software components.