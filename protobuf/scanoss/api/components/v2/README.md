# SCANOSS Components Service API v2

Provides component intelligence services including component search, version information, and usage statistics.

## SearchComponents

Searches for software components across multiple package ecosystems by name, vendor, or PURL patterns. Returns matching components with basic metadata.

### Request Format

The search supports multiple search modes:
- **General search**: Use the `search` parameter for free-text component search
- **Vendor/component search**: Use `vendor` and `component` parameters for targeted searches
- **Package type filtering**: Use `package` parameter to filter by ecosystem (github, maven, npm, etc.)

### HTTP Request Example
```bash
curl -X GET 'https://api.scanoss.com/v2/components/search?search=scanoss&limit=10' \
  -H 'x-api-key: $SC_API_KEY'
```

### gRPC Request Example
```bash
grpcurl -H "x-api-key: $SC_API_KEY" \
  -d '{
    "search": "scanoss",
    "limit": 10
  }' \
  api.scanoss.com:443 \
  scanoss.api.components.v2.Components/SearchComponents
```

### Response Example
```json
{
  "components": [
    {
      "name": "scanoss-py",
      "purl": "pkg:github/scanoss/scanoss.py",
      "url": "https://github.com/scanoss/scanoss.py",
      "component": "scanoss-py"
    },
    {
      "name": "engine",
      "purl": "pkg:github/scanoss/engine",
      "url": "https://github.com/scanoss/engine",
      "component": "engine"
    }
  ],
  "status": {
    "status": "SUCCESS",
    "message": "Components successfully retrieved"
  }
}
```

**Note**: The `component` field is deprecated and will be removed in future versions. Use the `name` field instead, which provides the same information.

## GetComponentVersions

Retrieves all available versions for a specific software component identified by Package URL, including version metadata and licensing information.

### Request Format

Requires a valid PURL (Package URL) to identify the component. Optional `limit` parameter controls the number of versions returned.

### HTTP Request Example
```bash
curl -X GET 'https://api.scanoss.com/v2/components/versions?purl=pkg:github/scanoss/engine&limit=20' \
  -H 'x-api-key: $SC_API_KEY'
```

### gRPC Request Example
```bash
grpcurl -H "x-api-key: $SC_API_KEY" \
  -d '{
    "purl": "pkg:github/scanoss/engine",
    "limit": 20
  }' \
  api.scanoss.com:443 \
  scanoss.api.components.v2.Components/GetComponentVersions
```

### Response Format

The method returns comprehensive version information including:

- `component` object: Contains the component details and version list
  - `name` field: Component name
  - `purl` field: Package URL identifier
  - `url` field: Component repository or homepage URL
  - `versions` array: List of available versions with metadata
- `status` field: Response status indicating success or failure

Each version object contains:
- Version identifier and release date
- Associated licenses with SPDX information
- SPDX approval status for each license

### Response Example
```json
{
  "component": {
    "name": "engine",
    "purl": "pkg:github/scanoss/engine",
    "url": "https://github.com/scanoss/engine",
    "versions": [
      {
        "version": "5.0.0",
        "licenses": [
          {
            "name": "GNU General Public License v2.0",
            "spdx_id": "GPL-2.0",
            "is_spdx_approved": true,
            "url": "https://spdx.org/licenses/GPL-2.0.html"
          }
        ],
        "date": "2024-01-15T10:30:00Z"
      },
      {
        "version": "4.5.0",
        "licenses": [
          {
            "name": "GNU General Public License v2.0",
            "spdx_id": "GPL-2.0",
            "is_spdx_approved": true,
            "url": "https://spdx.org/licenses/GPL-2.0.html"
          }
        ],
        "date": "2023-12-10T14:20:00Z"
      }
    ],
    "component": "engine"
  },
  "status": {
    "status": "SUCCESS",
    "message": "Component versions successfully retrieved"
  }
}
```

**Note**: The `component` field within the component object is deprecated and will be removed in future versions. Use the `name` field instead, which provides the same information.

## GetComponentStatistics

Analyzes software components to provide code statistics including file counts, line counts, and programming language breakdown. Useful for understanding component complexity and composition.

### Request Format
See [Common API Types](../common/v2/README.md) for `PurlRequest` documentation.

### HTTP Request Example
```bash
curl -X POST 'https://api.scanoss.com/v2/components/statistics' \
  -H 'Content-Type: application/json' \
  -H 'x-api-key: $SC_API_KEY' \
  -d '{
    "purls": [
      "pkg:github/scanoss/engine@5.0.0",
      "pkg:github/scanoss/scanoss.py@1.30.0"
    ]
  }'
```

### gRPC Request Example
```bash
grpcurl -H "x-api-key: $SC_API_KEY" \
  -d '{
    "purls": [
      "pkg:github/scanoss/engine@5.0.0",
      "pkg:github/scanoss/scanoss.py@1.30.0"
    ]
  }' \
  api.scanoss.com:443 \
  scanoss.api.components.v2.Components/GetComponentStatistics
```

### Response Format

The method returns comprehensive code statistics including:

- `components` array: List of analyzed components with their statistics
- `status` field: Response status indicating success or failure

Each component statistics object contains:
- Component identification (PURL and version)
- Overall code metrics (total files, lines, blank lines)
- Language breakdown with file counts per programming language

### Response Examples

#### Components with Code Statistics
```json
{
  "component_statistics": [
    {
      "purl": "pkg:github/scanoss/engine@5.0.0",
      "version": "5.0.0",
      "statistics": {
        "total_source_files": 156,
        "total_lines": 25430,
        "total_blank_lines": 3420,
        "languages": [
          {
            "name": "C",
            "files": 89
          },
          {
            "name": "C Header",
            "files": 45
          },
          {
            "name": "Makefile",
            "files": 12
          },
          {
            "name": "Shell",
            "files": 10
          }
        ]
      }
    },
    {
      "purl": "pkg:github/scanoss/scanoss.py@1.30.0",
      "version": "1.30.0",
      "statistics": {
        "total_source_files": 23,
        "total_lines": 4520,
        "total_blank_lines": 680,
        "languages": [
          {
            "name": "Python",
            "files": 20
          },
          {
            "name": "Markdown",
            "files": 2
          },
          {
            "name": "YAML",
            "files": 1
          }
        ]
      }
    }
  ],
  "status": {
    "status": "SUCCESS",
    "message": "Component statistics successfully retrieved"
  }
}
```

#### Component with No Available Statistics
```json
{
  "component_statistics": [
    {
      "purl": "pkg:github/example/unknown-component@1.0.0",
      "version": "1.0.0",
      "statistics": {
        "total_source_files": 0,
        "total_lines": 0,
        "total_blank_lines": 0,
        "languages": []
      }
    }
  ],
  "status": {
    "status": "SUCCESS",
    "message": "Component statistics successfully retrieved"
  }
}
```