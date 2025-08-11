# SCANOSS License Service API v2

Analyzes software components to identify licensing information. 

## GetLicenses

Retrieves license information for software components identified by Package URLs. Examines source code, license files, 
and package metadata to determine which licenses apply to each component. 
Returns license data in both individual SPDX license and SPDX expressions when determinable.

### Request Format
See [Common API Types](../common/v2/README.md) for `ComponentBatchRequest` documentation.

### Response Format
The method returns comprehensive license data in both individual SPDX license and SPDX expression formats when determinable. This addresses the same scenarios identified in the [CycloneDX specification](https://cyclonedx.org/docs/1.6/json/#components_items_licenses):
- `licenses` array: Always contains individual license objects found in the component
- `statement` field: Contains SPDX expression when licensing terms are clearly determinable from source analysis
- `version` field: Shows the specific version that was analyzed  
- `requirement` field: Echoes the client's version constraint from the request

### HTTP Request Example
```bash
curl -X POST 'https://api.scanoss.com/api/v2/licenses/purl' \
  -H 'Content-Type: application/json' \
  -H 'x-api-key: $SC_API_KEY' \
  -d '{
    "components": [
      {"purl": "pkg:github/scanoss/engine@5.0.0"},
      {"purl": "pkg:github/scanoss/scanoss.py@v1.30.0"}
    ]
  }'
```

### gRPC Request Example
```bash
grpcurl -H "x-api-key: $SC_API_KEY" \
  -d '{
    "components": [
      {"purl": "pkg:github/scanoss/engine@5.0.0"},
      {"purl": "pkg:github/scanoss/scanoss.py@v1.30.0"}
    ]
  }' \
  api.scanoss.com:443 \
  scanoss.api.licenses.v2.License/GetLicenses
```

### Response Examples

#### Licenses Array Only
Occurs when the component contains multiple license files or declarations without a unified SPDX expression that can be reliably determined.

**FFmpeg Analysis:**
FFmpeg contains separate LGPL and GPL license files throughout the codebase. The project does not provide a single SPDX expression to describe the overall licensing model, so the method returns each discovered license individually. This allows consumers to understand all licensing obligations present in the component.

```json
{
  "components": [{
    "purl": "pkg:github/ffmpeg/ffmpeg@n7.0",
    "requirement": "",
    "version": "n7.0",
    "statement": "",
    "licenses": [
      {"id": "LGPL-2.1-or-later", "full_name": "GNU Lesser General Public License v2.1 or later"},
      {"id": "GPL-2.0-or-later", "full_name": "GNU General Public License v2.0 or later"}
    ]
  }]
}
```

#### SPDX Expression with Choice (OR)
Occurs when the component's source code headers or documentation explicitly indicate a choice between licenses.

**Logback Analysis:**
Logback's source code headers explicitly state the component is "dual-licensed under either the terms of the Eclipse Public License v1.0... or... under the terms of the GNU Lesser General Public License version 2.1". This clear language allows the method to generate the SPDX expression `EPL-1.0 OR LGPL-2.1-only`, indicating users can choose either license for compliance. Individual license objects are also provided for detailed analysis.

```json
{
  "components": [{
    "purl": "pkg:maven/ch.qos.logback/logback-classic@1.5.0",
    "requirement": "",
    "version": "1.5.0",
    "statement": "EPL-1.0 OR LGPL-2.1-only",
    "licenses": [
      {"id": "EPL-1.0", "full_name": "Eclipse Public License 1.0"},
      {"id": "LGPL-2.1-only", "full_name": "GNU Lesser General Public License v2.1 only"}
    ]
  }]
}
```

#### SPDX Expression with Requirements (AND)
Occurs when the component requires compliance with multiple licenses simultaneously.

**OpenSSL Analysis:**
OpenSSL documentation states that "both the conditions of the OpenSSL License and the original SSLeay license apply to the toolkit". This indicates users must comply with both licenses, generating the SPDX expression `OpenSSL AND SSLeay`. Unlike OR expressions, there is no choice—both licenses apply simultaneously and must be satisfied.

```json
{
  "components": [{
    "purl": "pkg:github/openssl/openssl@1.1.1n",
    "requirement": "",
    "version": "1.1.1n",
    "statement": "OpenSSL AND SSLeay",
    "licenses": [
      {"id": "OpenSSL", "full_name": "OpenSSL License"},
      {"id": "SSLeay", "full_name": "Original SSLeay License"}
    ]
  }]
}
```