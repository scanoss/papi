# SCANOSS License Service API v2

Analyzes software components to identify licensing information.

> **Breaking change:** The `error_message` and `error_code` fields on `ComponentLicenseInfo` have been removed. Component-level processing outcomes are now reported via the `info_message` and `info_code` fields. Clients must migrate to read `info_message`/`info_code`; responses no longer include `error_message`/`error_code`.

## GetComponentLicenses
Retrieves license information for a single software component identified by Package URL. 
Examines source code, license files, and package metadata to determine which licenses apply to the component. 

### Request Format
See [Common API Types](../common/v2/README.md) for `ComponentRequest` documentation.

### HTTP Request Example
```bash
curl -X GET 'https://api.scanoss.com/v2/licenses/component?purl=pkg:github/scanoss/engine@5.0.0' \
  -H "X-Api-Key: $SC_API_KEY"
```

## Response Format

The method returns license information in two complementary formats:

- **Individual licenses**: A list of all licenses found in the component (always provided)
- **SPDX expression**: A standardized expression showing how licenses combine (provided when the relationship between licenses can be determined)

The response includes these fields:
- `purl` field: the requested component
- `licenses` array: Always contains individual license objects found in the component. Each license object includes:
  - `id`: SPDX identifier or licenseRef
  - `full_name`: Human-readable license name
  - `is_spdx_approved`: Whether the license is approved by the SPDX License List
  - `url`: URL to the license text or reference page
- `statement` field: Contains SPDX expression when licensing terms are clearly determinable from source analysis
- `version` field: Shows the specific version that was analyzed  
- `url` field: URL linking to the component's source or repository page
- `requirement` field: Echoes the client's version constraint from the request
- `info_code` field: Always populated. Identifies the outcome of processing the component (e.g.`VERSION_NOT_FOUND`)
- `info_message` field: Human-readable description of the processing outcome. Populated on errors and may be present on success

### Info Codes

The `info_code` field reports the outcome of processing each component. Possible values:

| Code | Meaning |
|------|---------|
| `SUCCESS` | Component processed successfully. |
| `INVALID_PURL` | The provided Package URL (PURL) is invalid or malformed. |
| `COMPONENT_NOT_FOUND` | The requested component could not be found in the database. |
| `NO_INFO` | No license information is available for the requested component. |
| `INVALID_SEMVER` | The provided semantic version (SemVer) is invalid or malformed. |
| `VERSION_NOT_FOUND` | The specific component version could not be found. |

### Response Examples

#### Licenses Array Only
Occurs when the component contains multiple license files or declarations without a unified SPDX expression that can be reliably determined.

**FFmpeg Analysis:**
FFmpeg contains separate LGPL and GPL license files throughout the codebase. 
The project does not provide a single SPDX expression to describe the overall licensing model, 
so the method returns each discovered license individually. 

This allows consumers to understand all licensing obligations present in the component.

```json
{
  "component": {
    "purl": "pkg:github/ffmpeg/ffmpeg@n7.0",
    "url": "https://github.com/ffmpeg/ffmpeg",
    "requirement": "",
    "version": "n7.0",
    "statement": "",
    "licenses": [
      {
        "id": "LGPL-2.1-or-later",
        "full_name": "GNU Lesser General Public License v2.1 or later",
        "is_spdx_approved": true,
        "url": "https://spdx.org/licenses/LGPL-2.1-or-later.html"
      },
      {
        "id": "GPL-2.0-or-later",
        "full_name": "GNU General Public License v2.0 or later",
        "is_spdx_approved": true,
        "url": "https://spdx.org/licenses/GPL-2.0-or-later.html"
      }
    ]
  },
  "status": {
    "status": "SUCCESS",
    "message": "Licenses Successfully retrieved"
  }
}
```

#### SPDX Expression with Choice (OR)
Occurs when the component's source code headers or documentation explicitly indicate a choice between licenses.

**Logback Analysis:**
Logback's [LICENSE file](https://github.com/qos-ch/logback/blob/6beda7b9d4bf2cbab4129da11efa6a635304a478/logback-classic/LICENSE.txt#L7-L14) explicitly states the component is
_"dual-licensed under either the terms of the Eclipse Public License v1.0... or... under the terms of the GNU Lesser General Public License version 2.1"_.

This clear language allows the method to generate the SPDX expression `EPL-1.0 OR LGPL-2.1-only`,
indicating users can choose either license for compliance. 

Individual license objects are also provided for detailed analysis.

```json
{
  "component": {
    "purl": "pkg:maven/ch.qos.logback/logback-classic@1.5.0",
    "url": "https://github.com/qos-ch/logback",
    "requirement": "",
    "version": "1.5.0",
    "statement": "EPL-1.0 OR LGPL-2.1-only",
    "licenses": [
      {"id": "EPL-1.0", "full_name": "Eclipse Public License 1.0", "is_spdx_approved": true, "url": "https://spdx.org/licenses/EPL-1.0.html"},
      {"id": "LGPL-2.1-only", "full_name": "GNU Lesser General Public License v2.1 only", "is_spdx_approved": true, "url": "https://spdx.org/licenses/LGPL-2.1-only.html"}
    ]
  },
  "status": {
    "status": "SUCCESS",
    "message": "Licenses Successfully retrieved"
  }
}
```

#### SPDX Expression with Requirements (AND)
Occurs when the component requires compliance with multiple licenses simultaneously.

**OpenSSL Analysis:**
OpenSSL [LICENSE file](https://github.com/openssl/openssl/blob/d82e959e621a3d597f1e0d50ff8c2d8b96915fd7/LICENSE#L5-L7) states that 
_"both the conditions of the OpenSSL License and the original SSLeay license apply to the toolkit"_.

This indicates users must comply with both licenses, generating the SPDX expression `OpenSSL AND SSLeay`. 

```json
{
  "component": {
    "purl": "pkg:github/openssl/openssl@1.1.1n",
    "url": "https://github.com/openssl/openssl",
    "requirement": "",
    "version": "1.1.1n",
    "statement": "OpenSSL AND SSLeay",
    "licenses": [
      {"id": "OpenSSL", "full_name": "OpenSSL License", "is_spdx_approved": false, "url": "https://www.openssl.org/source/license-openssl-ssleay.txt"},
      {"id": "SSLeay", "full_name": "Original SSLeay License", "is_spdx_approved": false, "url": "https://www.openssl.org/source/license-openssl-ssleay.txt"}
    ]
  },
  "status": {
    "status": "SUCCESS",
    "message": "Licenses Successfully retrieved"
  }
}
```



#### Error in component
When a component cannot be processed, the component block reports the failure via `info_code` and `info_message`. The remaining fields (`licenses`, `statement`, `version`, `url`) will be empty since the component could not be resolved.

```json
{
  "component": {
    "purl": "pkg:github/scanoss/unknown-component",
    "url": "",
    "requirement": "",
    "version": "",
    "statement": "",
    "licenses": [],
    "info_message": "Component version not found",
    "info_code": "VERSION_NOT_FOUND"
  },
  "status": {
    "status": "SUCCESS",
    "message": "Licenses Successfully retrieved"
  }
}
```

## GetComponentsLicenses

Batch version of GetLicenses - retrieves license information for multiple components in a single request.

### Request Format
See [Common API Types](../common/v2/README.md) for `ComponentsRequest` documentation.

### HTTP Request Example
```bash
curl -X POST 'https://api.scanoss.com/v2/licenses/components' \
  -H 'Content-Type: application/json' \
  -H "X-Api-Key: $SC_API_KEY" \
  -d '{
    "components": [
      {"purl": "pkg:github/scanoss/engine@5.0.0"},
      {"purl": "pkg:github/scanoss/scanoss.py@v1.30.0"}
    ]
  }'
```
