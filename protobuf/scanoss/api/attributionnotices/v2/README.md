# SCANOSS Attribution Notices Service API v2

Provides attribution/license notices mined for software components. Given a list of PURLs with optional version requirements, the service returns the attribution notices associated with each component.

## GetComponentsAttributionNotices

Retrieves attribution notices for a list of software components.

### Request Format
See [Common API Types](../common/v2/README.md) for `ComponentsRequest` documentation.

### HTTP Request Example
```bash
curl -X POST 'https://api.scanoss.com/v2/attribution-notices/components' \
  -H 'Content-Type: application/json' \
  -H "X-Api-Key: $SC_API_KEY" \
  -d '{
    "components": [
      {"purl": "pkg:npm/express@4.18.2"},
      {"purl": "pkg:example/example"}
    ]
  }'
```

### Response Format

The method returns attribution notices information including:

- `attributions` array: Contains attribution notices data for each requested component
- `status` field: Response status indicating success or failure of the request

Each element in `attributions` contains:
- `purl`: Package URL identifying the component
- `notices`: List of attribution notices mined for the component
  - `notice_text`: Full text of the attribution/license notice
  - `notice_md5`: MD5 hash of the notice text
  - `source`: Where the notice was mined from

### Response Examples

#### Component with Attribution Notices
```json
{
  "attributions": [
    {
      "purl": "pkg:example/example",
      "notices": [
        {
          "notice_text": "license text1",
          "notice_md5": "MD51",
          "source": "new file with more mining details"
        },
        {
          "notice_text": "license text2",
          "notice_md5": "MD52",
          "source": "new file with more mining details"
        }
      ]
    },
    {
      "purl": "pkg:npm/express@4.18.2",
      "notices": [
        {
          "notice_text": "MIT License\nPermission is hereby granted, free of charge...",
          "notice_md5": "a1b2c3d4e5f6a7b8c9d0e1f2a3b4c5d6",
          "source": "package.json dependency scan"
        }
      ]
    }
  ],
  "status": {
    "status": "SUCCESS",
    "message": "Attribution notices successfully retrieved"
  }
}
```

#### Status in component
The response reports the processing status of each component via `info_code` and `info_message`. When a component cannot be processed, the remaining fields will be empty since the component could not be resolved:
```json
{
  "attributions": [
    {
      "purl": "pkg:npm/unknown-component@1.0.0",
      "notices": [],
      "info_message": "Component not found",
      "info_code": "COMPONENT_NOT_FOUND"
    }
  ],
  "status": {
    "status": "SUCCESS",
    "message": "Attribution notices successfully retrieved"
  }
}
```

## Echo

Standard service health check endpoint for testing connectivity and API key validation.

### HTTP Request Example
```bash
curl -X POST 'https://api.scanoss.com/v2/attribution-notices/echo' \
  -H 'Content-Type: application/json' \
  -H "X-Api-Key: $SC_API_KEY" \
  -d '{"message": "test"}'
```
