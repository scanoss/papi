# SCANOSS Dependencies Service API v2

Provides comprehensive dependency analysis for software components including direct and transitive dependency resolution across multiple ecosystems.

## GetTransitiveDependencies

Analyzes software components to retrieve their complete transitive dependency tree, helping identify all indirect dependencies that would be pulled into a project.

**Note:** All components in the request must be from the same ecosystem. Mixing components from different ecosystems (e.g., npm and maven) will result in a bad request error.

### Request Format
- `depth` (optional): Number of dependency layers to traverse (e.g., depth=3 goes 3 levels deep)
- `limit` (optional): Maximum number of components to return in the response
- `components`: Array of components to analyze - see [Common API Types](../common/v2/README.md) for `ComponentsRequest` documentation

### HTTP Request Example
```bash
curl -X POST 'https://api.scanoss.com/v2/dependencies/transitive/components' \
  -H 'Content-Type: application/json' \
  -H "X-Api-Key: $SC_API_KEY" \
  -d '{
    "depth": 3,
    "limit": 50,
    "components": [
      {"purl": "pkg:npm/express", "requirement": "4.18.0"},
      {"purl": "pkg:npm/lodash", "requirement": "4.17.0"}
    ]
  }'
```

### gRPC Request Example
```bash
grpcurl -H "X-Api-Key: $SC_API_KEY" \
  -d '{
    "depth": 3,
    "limit": 50,
    "components": [
      {"purl": "pkg:npm/express", "requirement": "4.18.0"},
      {"purl": "pkg:npm/lodash", "requirement": "4.17.0"}
    ]
  }' \
  api.scanoss.com:443 \
  scanoss.api.dependencies.v2.Dependencies/GetTransitiveDependencies
```

### Response Format

The method returns transitive dependency information including:

- `dependencies` array: Complete list of direct and transitive dependencies found
- `status` field: Response status indicating success or failure

Each dependency object contains:
- `purl`: Package URL identifying the dependency
- `version`: Specific version resolved for this dependency
- `requirement`: Version constraint that led to this resolution

### Response Examples

#### Complete Transitive Analysis
```json
{
  "dependencies": [
    {
      "purl": "pkg:npm/express",
      "version": "4.18.2",
      "requirement": "4.18.0"
    },
    {
      "purl": "pkg:npm/body-parser", 
      "version": "1.20.1",
      "requirement": "~1.20.1"
    },
    {
      "purl": "pkg:npm/cookie",
      "version": "0.5.0", 
      "requirement": "0.5.0"
    },
    {
      "purl": "pkg:npm/lodash",
      "version": "4.17.21",
      "requirement": "4.17.0"
    }
  ],
  "status": {
    "status": "SUCCESS",
    "message": "Transitive dependencies successfully retrieved"
  }
}
```

#### Component with No Dependencies
```json
{
  "dependencies": [],
  "status": {
    "status": "SUCCESS",
    "message": "Transitive dependencies successfully retrieved"
  }
}
```

## GetDependencies (Deprecated)

**Note: This method is deprecated and will be removed in a future version. Please use GetTransitiveDependencies instead.**

Legacy method for analyzing dependency files and extracting direct dependencies with license information.

### Request Format
- `files`: Array of dependency files to analyze
- `depth`: Depth to search when analyzing dependencies

### HTTP Request Example
```bash
curl -X POST 'https://api.scanoss.com/v2/dependencies/dependencies' \
  -H 'Content-Type: application/json' \
  -H "X-Api-Key: $SC_API_KEY" \
  -d '{
    "files": [
      {
        "file": "package.json",
        "purls": [
          {
            "purl": "pkg:npm/express",
            "requirement": "^4.18.0"
          }
        ]
      }
    ],
    "depth": 1
  }'
```

### gRPC Request Example
```bash
grpcurl -H "X-Api-Key: $SC_API_KEY" \
  -d '{
    "files": [
      {
        "file": "package.json", 
        "purls": [
          {
            "purl": "pkg:npm/express",
            "requirement": "^4.18.0"
          }
        ]
      }
    ],
    "depth": 1
  }' \
  api.scanoss.com:443 \
  scanoss.api.dependencies.v2.Dependencies/GetDependencies
```

### Response Format

**Warning: This format is deprecated and subject to change.**

The method returns file-based dependency information including:

- `files` array: Analysis results grouped by dependency file
- `status` field: Response status indicating success or failure

Each file object contains:
- `file`: Name of the analyzed dependency file
- `id`: Analysis identifier
- `status`: Processing status for this file
- `dependencies`: Array of resolved dependencies with license information

Each dependency object includes:
- `component`: Component name
- `purl`: Package URL
- `version`: Resolved version
- `licenses`: Array of license information
- `url`: Component homepage or repository URL
- `comment`: Additional analysis notes

### Response Example
```json
{
  "files": [
    {
      "file": "package.json",
      "id": "dependency",
      "status": "pending",
      "dependencies": [
        {
          "component": "express",
          "purl": "pkg:npm/express",
          "version": "4.18.2",
          "licenses": [
            {
              "name": "MIT",
              "spdx_id": "MIT",
              "is_spdx_approved": true,
              "url": "https://opensource.org/licenses/MIT"
            }
          ],
          "url": "https://www.npmjs.com/package/express",
          "comment": ""
        }
      ]
    }
  ],
  "status": {
    "status": "SUCCESS",
    "message": "Dependencies successfully retrieved"
  }
}
```

## Echo

Standard service health check endpoint for testing connectivity and API key validation.

### HTTP Request Example
```bash
curl -X POST 'https://api.scanoss.com/v2/dependencies/echo' \
  -H 'Content-Type: application/json' \
  -H "X-Api-Key: $SC_API_KEY" \
  -d '{"message": "test"}'
```

### gRPC Request Example
```bash
grpcurl -H "X-Api-Key: $SC_API_KEY" \
  -d '{"message": "test"}' \
  api.scanoss.com:443 \
  scanoss.api.dependencies.v2.Dependencies/Echo
```