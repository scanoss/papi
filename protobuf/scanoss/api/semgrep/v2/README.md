# SCANOSS Semgrep Security Analysis Service API v2

Provides comprehensive security vulnerability analysis for software components using Semgrep static analysis engine. Detects potential security issues, code quality problems, and compliance violations.

## GetComponentsIssues

Retrieves security issues detected by Semgrep analysis for multiple software components. Analyzes component source code to identify potential vulnerabilities, security anti-patterns, and code quality issues.

### Request Format
See [Common API Types](../common/v2/README.md) for `ComponentsRequest` documentation.

### HTTP Request Example
```bash
curl -X POST 'https://api.scanoss.com/v2/semgrep/issues/components' \
  -H 'Content-Type: application/json' \
  -H "X-Api-Key: $SC_API_KEY" \
  -d '{
    "components": [
      {
        "purl": "pkg:maven/org.apache.commons/commons-lang3",
        "requirement": "3.12.0"
      }
    ]
  }'
```

### Response Format

The method returns security analysis information including:

- `components` array: Contains security issue data for each requested component
- `status` field: Response status indicating success or failure of the request

Each component object contains:
- `purl`: Package URL identifying the component
- `version`: Version of the component being analyzed
- `requirement`: Version requirement specification
- `files`: Array of files containing security issues

Each file object contains:
- `fileMD5`: MD5 hash of the file content for integrity verification
- `path`: Relative path of the file within the component
- `issues`: Array of security issues found in the file

Each issue object contains:
- `ruleID`: Unique identifier for the Semgrep rule that detected the issue
- `from`: Starting line number where the issue was detected
- `to`: Ending line number where the issue was detected
- `severity`: Severity level (ERROR, WARNING, INFO)

### Response Examples

#### Component with Security Issues
```json
{
  "components": [
    {
      "purl": "pkg:maven/org.apache.commons/commons-lang3",
      "version": "3.12.0",
      "requirement": "3.12.0",
      "files": [
        {
          "fileMD5": "a1b2c3d4e5f6",
          "path": "src/main/java/org/apache/commons/lang3/StringUtils.java",
          "issues": [
            {
              "ruleID": "java.lang.security.audit.crypto.weak-hash",
              "from": "156",
              "to": "159",
              "severity": "WARNING"
            },
            {
              "ruleID": "java.lang.security.audit.sql-injection.sql-injection",
              "from": "284",
              "to": "286",
              "severity": "ERROR"
            }
          ]
        },
        {
          "fileMD5": "b2c3d4e5f6a1",
          "path": "src/main/java/org/apache/commons/lang3/Validate.java",
          "issues": [
            {
              "ruleID": "java.lang.security.audit.hardcoded-secret",
              "from": "95",
              "to": "95",
              "severity": "ERROR"
            }
          ]
        }
      ]
    }
  ],
  "status": {
    "status": "SUCCESS",
    "message": "Security analysis completed successfully"
  }
}
```

#### Component with No Issues Found
```json
{
  "components": [
    {
      "purl": "pkg:maven/org.springframework/spring-core",
      "version": "5.3.21",
      "requirement": "5.3.21",
      "files": []
    }
  ],
  "status": {
    "status": "SUCCESS",
    "message": "Security analysis completed successfully"
  }
}
```

## GetComponentIssues

Retrieves security issues detected by Semgrep analysis for a single software component. Analyzes component source code to identify potential vulnerabilities, security anti-patterns, and code quality issues.

### Request Format
See [Common API Types](../common/v2/README.md) for `ComponentRequest` documentation.

### HTTP Request Example
```bash
curl -X GET 'https://api.scanoss.com/v2/semgrep/issues/component?purl=pkg:maven/org.apache.commons/commons-lang3&requirement=3.12.0' \
  -H "X-Api-Key: $SC_API_KEY"
```

### Response Format

The method returns security analysis information including:

- `component` object: Contains security issue data for the requested component
- `status` field: Response status indicating success or failure of the request

The component object contains:
- `purl`: Package URL identifying the component
- `version`: Version of the component being analyzed
- `requirement`: Version requirement specification
- `files`: Array of files containing security issues

### Response Examples

#### Component with Security Issues
```json
{
  "component": {
    "purl": "pkg:maven/org.apache.commons/commons-lang3",
    "version": "3.12.0",
    "requirement": "3.12.0",
    "files": [
      {
        "fileMD5": "a1b2c3d4e5f6",
        "path": "src/main/java/org/apache/commons/lang3/StringUtils.java",
        "issues": [
          {
            "ruleID": "java.lang.security.audit.sql-injection.sql-injection",
            "from": "284",
            "to": "286",
            "severity": "ERROR"
          }
        ]
      }
    ]
  },
  "status": {
    "status": "SUCCESS",
    "message": "Security analysis completed successfully"
  }
}
```

#### Component with No Issues Found
```json
{
  "component": {
    "purl": "pkg:maven/org.springframework/spring-core",
    "version": "5.3.21",
    "requirement": "5.3.21",
    "files": []
  },
  "status": {
    "status": "SUCCESS",
    "message": "Security analysis completed successfully"
  }
}
```

## GetIssues (Deprecated)

**Note: This method is deprecated and will be removed in a future version. Please use GetComponentsIssues instead.**

Legacy method for analyzing security issues using PURL-based requests. Provides the same functionality as GetComponentsIssues but with the older request format.

### Request Format
See [Common API Types](../common/v2/README.md) for `PurlRequest` documentation.

### Migration Guide
To migrate from `GetIssues` to `GetComponentsIssues`:

1. **Change the request format**: Use `ComponentsRequest` instead of `PurlRequest`
2. **Update response handling**: Use `ComponentsIssueResponse` instead of `SemgrepResponse`

## Echo

Standard service health check endpoint for testing connectivity and API key validation.

### HTTP Request Example
```bash
curl -X POST 'https://api.scanoss.com/v2/semgrep/echo' \
  -H 'Content-Type: application/json' \
  -H "X-Api-Key: $SC_API_KEY" \
  -d '{"message": "test"}'
```

## Security Analysis Categories

The Semgrep service analyzes components across multiple security dimensions:

### Vulnerability Detection
- SQL injection patterns
- Cross-site scripting (XSS) vulnerabilities
- Command injection flaws
- Path traversal vulnerabilities
- Authentication and authorization bypasses

### Code Quality Issues
- Hardcoded secrets and credentials
- Unsafe cryptographic practices
- Insecure random number generation
- Improper input validation
- Resource leaks and memory management issues

### Compliance and Best Practices
- OWASP Top 10 security risks
- CWE (Common Weakness Enumeration) categories
- Language-specific security anti-patterns
- Framework-specific security misconfigurations

### Supported Languages
- Java