# SCANOSS Cryptography Service API v2

Provides cryptographic intelligence for software components including algorithm detection, encryption hints, and cryptographic assessment across version ranges.

## GetComponentAlgorithms

Retrieves cryptographic algorithms detected in a single software component identified in a software component.

### Request Format
See [Common API Types](../common/v2/README.md) for `ComponentRequest` documentation.

### HTTP Request Example
```bash
curl -X GET 'https://api.scanoss.com/v2/cryptography/algorithms/component?purl=pkg:github/scanoss/engine&requirement=>=5.0.0' \
  -H 'x-api-key: $SC_API_KEY'
```

### gRPC Request Example
```bash
grpcurl -H "x-api-key: $SC_API_KEY" \
  -d '{
    "purl": "pkg:github/scanoss/engine",
    "requirement": ">=5.0.0"
  }' \
  api.scanoss.com:443 \
  scanoss.api.cryptography.v2.Cryptography/GetComponentAlgorithms
```

### Response Example
```json
{
  "component": {
    "purl": "pkg:github/scanoss/engine",
    "version": "5.0.0",
    "algorithms": [
      {
        "algorithm": "AES",
        "strength": "Strong"
      },
      {
        "algorithm": "RSA",
        "strength": "Strong"
      }
    ]
  },
  "status": {
    "status": "SUCCESS",
    "message": "Algorithms Successfully retrieved"
  }
}
```

## GetComponentsAlgorithms

Batch version of GetComponentAlgorithms - retrieves cryptographic algorithms for multiple software components in a single request.

### Request Format
See [Common API Types](../common/v2/README.md) for `ComponentsRequest` documentation.

### HTTP Request Example
```bash
curl -X POST 'https://api.scanoss.com/v2/cryptography/algorithms/components' \
  -H 'Content-Type: application/json' \
  -H 'x-api-key: $SC_API_KEY' \
  -d '{
    "components": [
      {"purl": "pkg:github/scanoss/engine", "requirement": ">=5.0.0"},
      {"purl": "pkg:github/scanoss/scanoss.py", "requirement": "~1.30.0"}
    ]
  }'
```

### gRPC Request Example
```bash
grpcurl -H "x-api-key: $SC_API_KEY" \
  -d '{
    "components": [
      {"purl": "pkg:github/scanoss/engine", "requirement": ">=5.0.0"},
      {"purl": "pkg:github/scanoss/scanoss.py", "requirement": "~1.30.0"}
    ]
  }' \
  api.scanoss.com:443 \
  scanoss.api.cryptography.v2.Cryptography/GetComponentsAlgorithms
```

## GetComponentAlgorithmsInRange

Analyzes a single software component across specified version ranges and returns all cryptographic algorithms detected along with the versions where they appear. This is useful for tracking cryptographic evolution across component development.

### Request Format
See [Common API Types](../common/v2/README.md) for `ComponentRequest` documentation.

### HTTP Request Example
```bash
curl -X GET 'https://api.scanoss.com/v2/cryptography/algorithms/range/component?purl=pkg:github/scanoss/engine&requirement=>=1.0.0' \
  -H 'x-api-key: $SC_API_KEY'
```

### gRPC Request Example
```bash
grpcurl -H "x-api-key: $SC_API_KEY" \
  -d '{
    "purl": "pkg:github/scanoss/engine",
    "requirement": ">=1.0.0,<6.0.0"
  }' \
  api.scanoss.com:443 \
  scanoss.api.cryptography.v2.Cryptography/GetComponentAlgorithmsInRange
```

### Response Format

The method returns comprehensive algorithm information including:

- `component` object: Contains the component analysis results
  - `purl` field: the requested component
  - `versions` array: List of component versions where cryptographic algorithms were detected
  - `algorithms` array: List of all cryptographic algorithms found across the version range
- `status` field: Response status indicating success or failure

Each algorithm object contains:
- Algorithm name and strength classification
- Detection metadata and analysis results

### Response Examples

#### Component with Algorithms Across Versions
```json
{
  "component": {
    "purl": "pkg:github/scanoss/engine",
    "versions": ["1.0.0", "2.0.0", "3.0.0", "5.0.0"],
    "algorithms": [
      {
        "algorithm": "AES",
        "strength": "Strong"
      },
      {
        "algorithm": "RSA",
        "strength": "Strong"
      },
      {
        "algorithm": "MD5",
        "strength": "Weak"
      }
    ]
  },
  "status": {
    "status": "SUCCESS",
    "message": "Algorithms in range Successfully retrieved"
  }
}
```

#### Component with No Cryptographic Algorithms
```json
{
  "component": {
    "purl": "pkg:github/example/simple-utility",
    "versions": [],
    "algorithms": []
  },
  "status": {
    "status": "SUCCESS",
    "message": "Algorithms in range Successfully retrieved"
  }
}
```

## GetComponentsAlgorithmsInRange

Batch version of GetComponentAlgorithmsInRange - analyzes multiple components across version ranges and returns cryptographic algorithms for each in a single request.

### Request Format
See [Common API Types](../common/v2/README.md) for `ComponentsRequest` documentation.

### HTTP Request Example
```bash
curl -X POST 'https://api.scanoss.com/v2/cryptography/algorithms/range/components' \
  -H 'Content-Type: application/json' \
  -H 'x-api-key: $SC_API_KEY' \
  -d '{
    "components": [
      {"purl": "pkg:github/scanoss/engine", "requirement": ">=1.0.0"},
      {"purl": "pkg:github/scanoss/scanoss.py", "requirement": "~1.30.0"}
    ]
  }'
```

### gRPC Request Example
```bash
grpcurl -H "x-api-key: $SC_API_KEY" \
  -d '{
    "components": [
      {"purl": "pkg:github/scanoss/engine", "requirement": ">=1.0.0"},
      {"purl": "pkg:github/scanoss/scanoss.py", "requirement": "~1.30.0"}
    ]
  }' \
  api.scanoss.com:443 \
  scanoss.api.cryptography.v2.Cryptography/GetComponentsAlgorithmsInRange
```

## ComponentVersionsInRange

Analyzes a software component and returns lists of versions that either contain cryptographic algorithms or don't, helping assess cryptographic presence across component evolution.

### Request Format
See [Common API Types](../common/v2/README.md) for `ComponentRequest` documentation.

### HTTP Request Example
```bash
curl -X GET 'https://api.scanoss.com/v2/cryptography/algorithms/versions/range/component?purl=pkg:github/scanoss/engine&requirement=>=1.0.0' \
  -H 'x-api-key: $SC_API_KEY'
```

### gRPC Request Example
```bash
grpcurl -H "x-api-key: $SC_API_KEY" \
  -d '{
    "purl": "pkg:github/scanoss/engine",
    "requirement": ">=1.0.0,<6.0.0"
  }' \
  api.scanoss.com:443 \
  scanoss.api.cryptography.v2.Cryptography/ComponentVersionsInRange
```

### Response Example
```json
{
  "component": {
    "purl": "pkg:github/scanoss/engine",
    "versions_with": ["2.0.0", "3.0.0", "4.0.0", "5.0.0"],
    "versions_without": ["1.0.0", "1.5.0"]
  },
  "status": {
    "status": "SUCCESS",
    "message": "Version ranges Successfully retrieved"
  }
}
```

## ComponentsVersionsInRange

Batch version of ComponentVersionsInRange - analyzes multiple components and returns version lists based on cryptographic algorithm presence for each component.

### Request Format
See [Common API Types](../common/v2/README.md) for `ComponentsRequest` documentation.

### HTTP Request Example
```bash
curl -X POST 'https://api.scanoss.com/v2/cryptography/algorithms/version/range/components' \
  -H 'Content-Type: application/json' \
  -H 'x-api-key: $SC_API_KEY' \
  -d '{
    "components": [
      {"purl": "pkg:github/scanoss/engine", "requirement": ">=1.0.0"},
      {"purl": "pkg:github/scanoss/scanoss.py", "requirement": "~1.30.0"}
    ]
  }'
```

### gRPC Request Example
```bash
grpcurl -H "x-api-key: $SC_API_KEY" \
  -d '{
    "components": [
      {"purl": "pkg:github/scanoss/engine", "requirement": ">=1.0.0"},
      {"purl": "pkg:github/scanoss/scanoss.py", "requirement": "~1.30.0"}
    ]
  }' \
  api.scanoss.com:443 \
  scanoss.api.cryptography.v2.Cryptography/ComponentsVersionsInRange
```

## ComponentHintsInRange

Retrieves cryptographic hints for a single component, providing insights about cryptographic protocols, libraries, SDKs and frameworks used by the component.

### Request Format
See [Common API Types](../common/v2/README.md) for `ComponentRequest` documentation.

### HTTP Request Example
```bash
curl -X GET 'https://api.scanoss.com/v2/cryptography/hints/component?purl=pkg:github/scanoss/engine&requirement=>=5.0.0' \
  -H 'x-api-key: $SC_API_KEY'
```

### gRPC Request Example
```bash
grpcurl -H "x-api-key: $SC_API_KEY" \
  -d '{
    "purl": "pkg:github/scanoss/engine",
    "requirement": ">=5.0.0"
  }' \
  api.scanoss.com:443 \
  scanoss.api.cryptography.v2.Cryptography/ComponentHintsInRange
```

### Response Format

The method returns comprehensive cryptographic hint information including:

- `purl` field: the requested component
- `hints` array: List of cryptographic hints detected in the component
- `version` field: Shows the specific version that was analyzed
- `status` field: Response status indicating success or failure

Each hint object contains:
- Unique identifier and name of the cryptographic entity
- Description of the detected usage or implementation
- Category classification (protocol/library/sdk/framework)
- Reference URL and PURL information

### Hint Categories

Hints are classified into the following categories:
- `protocol`: Cryptographic protocols (e.g., TLS, SSH, HTTPS)
- `library`: Cryptographic libraries (e.g., OpenSSL, Bouncy Castle)
- `sdk`: Software Development Kits with cryptographic capabilities
- `framework`: Frameworks that include cryptographic functionality

### Response Examples

#### Component with Cryptographic Hints
```json
{
  "component": {
    "purl": "pkg:github/scanoss/engine",
    "version": "5.0.0",
    "hints": [
      {
        "id": "openssl-hint-001",
        "name": "OpenSSL",
        "description": "Industry standard cryptographic library providing SSL/TLS protocols",
        "category": "library",
        "url": "https://www.openssl.org/",
        "purl": "pkg:generic/openssl@3.0.0"
      },
      {
        "id": "tls-protocol-001",
        "name": "TLS 1.3",
        "description": "Transport Layer Security protocol implementation",
        "category": "protocol",
        "url": "https://tools.ietf.org/html/rfc8446",
        "purl": ""
      }
    ]
  },
  "status": {
    "status": "SUCCESS",
    "message": "Cryptographic hints Successfully retrieved"
  }
}
```

#### Component with No Cryptographic Hints
```json
{
  "component": {
    "purl": "pkg:github/example/simple-utility",
    "version": "1.0.0",
    "hints": []
  },
  "status": {
    "status": "SUCCESS",
    "message": "Cryptographic hints Successfully retrieved"
  }
}
```

## ComponentsHintsInRange

Batch version of ComponentHintsInRange - retrieves cryptographic hints for multiple components in a single request, providing insights into cryptographic dependencies across multiple components.

### Request Format
See [Common API Types](../common/v2/README.md) for `ComponentsRequest` documentation.

### HTTP Request Example
```bash
curl -X POST 'https://api.scanoss.com/v2/cryptography/hints/components' \
  -H 'Content-Type: application/json' \
  -H 'x-api-key: $SC_API_KEY' \
  -d '{
    "components": [
      {"purl": "pkg:github/scanoss/engine", "requirement": ">=5.0.0"},
      {"purl": "pkg:github/scanoss/scanoss.py", "requirement": "~1.30.0"}
    ]
  }'
```

### gRPC Request Example
```bash
grpcurl -H "x-api-key: $SC_API_KEY" \
  -d '{
    "components": [
      {"purl": "pkg:github/scanoss/engine", "requirement": ">=5.0.0"},
      {"purl": "pkg:github/scanoss/scanoss.py", "requirement": "~1.30.0"}
    ]
  }' \
  api.scanoss.com:443 \
  scanoss.api.cryptography.v2.Cryptography/ComponentsHintsInRange
```