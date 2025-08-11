# SCANOSS Common API Types v2

This package defines reusable message types used across SCANOSS API services. **These are not callable endpoints** - for actual API usage and examples, see individual service documentation.
## Message Types

### ComponentRequest
Represents a software component to be analyzed by SCANOSS API services. Combines a Package URL for component identification with optional version constraints for resolution.

```protobuf
message ComponentRequest {
  string purl = 1;        // Package URL identifying the component
  string requirement = 2; // Version constraint for resolution
}
```
#### Fields

- **`purl:`** [Package URL](https://github.com/package-url/purl-spec) identifying the component. Examples: `pkg:github/scanoss/engine@1.0.0`, `pkg:github/scanoss/engine`

- **`requirement:`** [Version constraint](https://github.com/Masterminds/semver#checking-version-constraints) for component resolution. Examples: `^1.2.3`, `>=2.0.0`, `1.2.3`



#### Resolution Logic

SCANOSS API services apply the following logic when processing `ComponentRequest` messages:

| PURL                              | Requirement | Result |
|-----------------------------------|-------------|--------|
| `pkg:github/scanoss/engine@1.0.0` | `""`        | Analyze that exact version |
| `pkg:github/scanoss/engine`       | `""`        | Use latest version |
| `pkg:github/scanoss/engine`       | `^1.0.0`    | Find version matching constraint |
| `pkg:github/scanoss/engine`       | `1.0.0`     | Analyze that exact version |

### ComponentBatchRequest
Represents a list of software component to be analyzed by SCANOSS API services. 
Allows analysis of multiple software components in a single API call, improving performance over individual requests.

```protobuf
message ComponentBatchRequest {
  repeated ComponentRequest components = 1;
}
```

#### Fields

- **`components:`** Array of individual component requests
  - Each component analyzed independently
  - Batch processing improves performance for multiple components
  - Supports mixed PURL formats and requirement types


## Usage

`ComponentRequest` and `ComponentBatchRequest` are used by SCANOSS API services for component analysis. See individual service documentation for specific endpoints and examples.

**Note**: Some legacy services still use the deprecated `PurlRequest` message type, which will be migrated to `ComponentRequest` in future versions.

## Resources

- [Protocol Buffer Definition](./scanoss-common.proto): Complete message specifications
- [PURL Specification](https://github.com/package-url/purl-spec): Official Package URL standard
- [SCANOSS Documentation](https://docs.scanoss.com): Full platform documentation

## Issues & Support

- API Issues: [GitHub Issues](https://github.com/scanoss/papi/issues)
- General Support: support@scanoss.com
- Documentation: [docs.scanoss.com](https://docs.scanoss.com)