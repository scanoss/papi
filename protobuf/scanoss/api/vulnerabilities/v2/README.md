# SCANOSS Vulnerability Service API v2

Provides vulnerability intelligence for software components including CPE enumeration and vulnerability analysis.

## GetComponentCpes

Retrieves Common Platform Enumeration (CPE) identifiers for a single software component identified by Package URL. 
CPEs are used to identify IT platforms in vulnerability databases and enable vulnerability scanning and assessment.

### Request Format
See [Common API Types](../common/v2/README.md) for `ComponentRequest` documentation.

### HTTP Request Example
```bash
curl -X GET 'https://api.scanoss.com/v2/vulnerabilities/cpes/component?purl=pkg:github/scanoss/engine&requirement=>=5.0.0' \
  -H 'X-Api-Key: $SC_API_KEY'
```

### gRPC Request Example
```bash
grpcurl -H "X-Api-Key: $SC_API_KEY" \
  -d '{
    "purl": "pkg:github/scanoss/engine",
    "requirement": ">=5.0.0"
  }' \
  api.scanoss.com:443 \
  scanoss.api.vulnerabilities.v2.Vulnerabilities/GetComponentCpes
```

### Response Example
```json
{
  "component": {
    "purl": "pkg:github/scanoss/engine",
    "requirement": ">=5.0.0",
    "version": "5.0.0",
    "cpes": [
      "cpe:2.3:a:scanoss:engine:1.0.0:*:*:*:*:*:*:*"
    ]
  },
  "status": {
    "status": "SUCCESS",
    "message": "CPEs Successfully retrieved"
  }
}
```

## GetComponentsCpes

Batch version of GetComponentCpes - retrieves CPE identifiers for multiple components in a single request.

### Request Format
See [Common API Types](../common/v2/README.md) for `ComponentsRequest` documentation.

### HTTP Request Example
```bash
curl -X POST 'https://api.scanoss.com/v2/vulnerabilities/cpes/components' \
  -H 'Content-Type: application/json' \
  -H 'X-Api-Key: $SC_API_KEY' \
  -d '{
    "components": [
      {"purl": "pkg:github/scanoss/engine", "requirement": ">=5.0.0"},
      {"purl": "pkg:github/scanoss/scanoss.py", "requirement": "~1.30.0"}
    ]
  }'
```

### gRPC Request Example
```bash
grpcurl -H "X-Api-Key: $SC_API_KEY" \
  -d '{
    "components": [
      {"purl": "pkg:github/scanoss/engine", "requirement": ">=5.0.0"},
      {"purl": "pkg:github/scanoss/scanoss.py", "requirement": "~1.30.0"}
    ]
  }' \
  api.scanoss.com:443 \
  scanoss.api.vulnerabilities.v2.Vulnerabilities/GetComponentsCpes
```

## GetComponentVulnerabilities

Analyzes a single software component and returns known vulnerabilities including CVE details, severity scores, publication dates, and other security metadata. Vulnerability data is sourced from various security databases and feeds.

### Request Format
See [Common API Types](../common/v2/README.md) for `ComponentRequest` documentation.

### HTTP Request Example
```bash
curl -X GET 'https://api.scanoss.com/v2/vulnerabilities/component?purl=pkg:github/scanoss/engine&requirement=>=5.0.0' \
  -H 'X-Api-Key: $SC_API_KEY'
```

### gRPC Request Example
```bash
grpcurl -H "X-Api-Key: $SC_API_KEY" \
  -d '{
    "purl": "pkg:github/scanoss/engine",
    "requirement": ">=5.0.0"
  }' \
  api.scanoss.com:443 \
  scanoss.api.vulnerabilities.v2.Vulnerabilities/GetComponentVulnerabilities
```

## Response Format

The method returns comprehensive vulnerability information including:

- `purl` field: the requested component
- `vulnerabilities` array: List of known vulnerabilities affecting the component
- `version` field: Shows the specific version that was analyzed  
- `requirement` field: Echoes the client's version constraint from the request

Each vulnerability object contains:
- CVE identifier and reference URL
- Severity classification and CVSS information
- Publication and modification dates
- Summary description
- Source database information
- CVSS array with detailed scoring information (vector, score, and severity)

### CVSS Information

The `cvss` field is an array of CVSS (Common Vulnerability Scoring System) objects, allowing for multiple CVSS versions or sources. Each CVSS object contains:

- `cvss`: The CVSS vector string (e.g., "CVSS:3.1/AV:N/AC:L/PR:N/UI:N/S:U/C:H/I:H/A:H")
- `cvss_score`: The numerical CVSS score (0.0 to 10.0)
- `cvss_severity`: The severity rating based on the score ("None", "Low", "Medium", "High", "Critical")

### Response Examples

#### Component with Vulnerabilities
```json
{
  "component": {
    "purl": "pkg:github/scanoss/engine",
    "requirement": ">=5.0.0",
    "version": "5.0.0",
    "vulnerabilities": [
      {
        "id": "CVE-2024-12345",
        "cve": "CVE-2024-12345",
        "url": "https://nvd.nist.gov/vuln/detail/CVE-2024-12345",
        "summary": "Buffer overflow vulnerability in input processing",
        "severity": "High",
        "published": "2024-01-15T10:30:00Z",
        "modified": "2024-01-16T14:20:00Z",
        "source": "NVD",
        "cvss": [
          {
            "cvss": "CVSS:3.1/AV:N/AC:L/PR:N/UI:N/S:U/C:H/I:H/A:H",
            "cvss_score": 7.5,
            "cvss_severity": "High"
          }
        ]
      }
    ]
  },
  "status": {
    "status": "SUCCESS",
    "message": "Vulnerabilities Successfully retrieved"
  }
}
```

#### Component with No Known Vulnerabilities
```json
{
  "component": {
    "purl": "pkg:github/scanoss/scanoss.py",
    "requirement": ">1.30.0",
    "version": "1.31.0",
    "vulnerabilities": []
  },
  "status": {
    "status": "SUCCESS",
    "message": "Vulnerabilities Successfully retrieved"
  }
}
```

## GetComponentsVulnerabilities

Batch version of GetComponentVulnerabilities - analyzes multiple components and returns vulnerability information for each in a single request.

### Request Format
See [Common API Types](../common/v2/README.md) for `ComponentsRequest` documentation.

### HTTP Request Example
```bash
curl -X POST 'https://api.scanoss.com/v2/vulnerabilities/components' \
  -H 'Content-Type: application/json' \
  -H 'X-Api-Key: $SC_API_KEY' \
  -d '{
    "components": [
      {"purl": "pkg:github/scanoss/engine", "requirement": ">=5.0.0"},
      {"purl": "pkg:github/scanoss/scanoss.py", "requirement": "~1.30.0"}
    ]
  }'
```

### gRPC Request Example
```bash
grpcurl -H "X-Api-Key: $SC_API_KEY" \
  -d '{
    "components": [
      {"purl": "pkg:github/scanoss/engine", "requirement": ">=5.0.0"},
      {"purl": "pkg:github/scanoss/scanoss.py", "requirement": "~1.30.0"}
    ]
  }' \
  api.scanoss.com:443 \
  scanoss.api.vulnerabilities.v2.Vulnerabilities/GetComponentsVulnerabilities
```