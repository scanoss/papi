# SCANOSS Scanning Service API v2

Provides high-precision folder scanning capabilities to identify software components through hierarchical folder hashing, enabling accurate component detection based on project structure, file names, and content patterns.

## FolderHashScan

Performs a high-precision scan of folder structures using proximity hashing algorithms to identify matching software components. This method analyzes the hierarchical organization of files and directories, computing similarity hashes from filenames, directory names, and content to find component matches in the SCANOSS knowledge base.

### Request Format

The scan request requires a hierarchical folder structure represented as a tree of nodes:

- `root`: The root folder node containing the structure to scan
  - `path_id`: Folder path identifier (can be actual or obfuscated for privacy)
  - `sim_hash_names`: Proximity hash calculated from filenames in this folder and its children
  - `sim_hash_content`: Proximity hash calculated from file contents in this folder and its children
  - `sim_hash_dir_names`: Proximity hash calculated from directory names in this folder and its children
  - `lang_extensions`: Language extension counts (e.g., {"py": 15, "js": 8})
  - `children`: Array of sub-folder nodes with the same structure

Optional filtering parameters:
- `rank_threshold`: Only return results with rank at or below this value (lower rank = better quality)
- `category`: Filter results by package ecosystem (e.g., "github", "npm", "maven")
- `query_limit`: Maximum number of results to return per folder
- `recursive_threshold`: Minimum score threshold to consider a match
- `min_accepted_score`: Minimum score threshold to accept a match (0-1 scale)

### HTTP Request Example
```bash
curl -X POST 'https://api.scanoss.com/v2/scanning/hfh/scan' \
  -H 'Content-Type: application/json' \
  -H "X-Api-Key: $SC_API_KEY" \
  -d '{
    "root": {
      "path_id": "src",
      "sim_hash_names": "abc123def456",
      "sim_hash_content": "789ghi012jkl",
      "sim_hash_dir_names": "345mno678pqr",
      "lang_extensions": {
        "py": 25,
        "md": 3
      },
      "children": [
        {
          "path_id": "src/utils",
          "sim_hash_names": "def456ghi789",
          "sim_hash_content": "012jkl345mno",
          "sim_hash_dir_names": "678pqr901stu",
          "lang_extensions": {
            "py": 8
          },
          "children": []
        }
      ]
    },
    "rank_threshold": 5,
    "category": "github",
    "query_limit": 10,
    "min_accepted_score": 0.7
  }'
```

### gRPC Request Example
```bash
grpcurl -H "X-Api-Key: $SC_API_KEY" \
  -d '{
    "root": {
      "path_id": "src",
      "sim_hash_names": "abc123def456",
      "sim_hash_content": "789ghi012jkl",
      "sim_hash_dir_names": "345mno678pqr",
      "lang_extensions": {
        "py": 25,
        "md": 3
      },
      "children": [
        {
          "path_id": "src/utils",
          "sim_hash_names": "def456ghi789",
          "sim_hash_content": "012jkl345mno",
          "sim_hash_dir_names": "678pqr901stu",
          "lang_extensions": {
            "py": 8
          },
          "children": []
        }
      ]
    },
    "rank_threshold": 5,
    "category": "github",
    "query_limit": 10,
    "min_accepted_score": 0.7
  }' \
  api.scanoss.com:443 \
  scanoss.api.scanning.v2.Scanning/FolderHashScan
```

### Response Format

The method returns comprehensive scan results including:

- `results` array: List of folder paths with their matching components
  - `path_id` field: The folder path identifier from the request
  - `components` array: List of matching components found in this folder
- `status` field: Response status indicating success or failure

Each component object contains:
- `purl`: Package URL identifier for the component
- `name`: Component name
- `vendor`: Component vendor/organization
- `versions`: Array of matched versions with scores (0-1 scale, higher is better)
- `rank`: Component ranking from 1 to 9 (1 = official well-known repository, 9 = potentially low-quality)
- `order`: Match ordering (1 = best match, higher numbers = lower confidence)

### Response Example

#### Successful Scan with Matches
```json
{
  "results": [
    {
      "path_id": "src",
      "components": [
        {
          "purl": "pkg:github/scanoss/scanoss.py",
          "name": "scanoss-py",
          "vendor": "scanoss",
          "versions": [
            {
              "version": "1.30.0",
              "score": 0.95
            },
            {
              "version": "1.29.0",
              "score": 0.87
            }
          ],
          "rank": 1,
          "order": 1
        },
        {
          "purl": "pkg:github/example/similar-project",
          "name": "similar-project",
          "vendor": "example",
          "versions": [
            {
              "version": "2.1.0",
              "score": 0.78
            }
          ],
          "rank": 3,
          "order": 2
        }
      ]
    },
    {
      "path_id": "src/utils",
      "components": [
        {
          "purl": "pkg:pypi/requests",
          "name": "requests",
          "vendor": "psf",
          "versions": [
            {
              "version": "2.31.0",
              "score": 0.92
            }
          ],
          "rank": 1,
          "order": 1
        }
      ]
    }
  ],
  "status": {
    "status": "SUCCESS",
    "message": "Scan completed successfully"
  }
}
```

#### Scan with No Matches
```json
{
  "results": [],
  "status": {
    "status": "SUCCESS",
    "message": "Scan completed successfully"
  }
}
```

### Understanding Match Quality

**Rank System** (1-9 scale):
- Rank 1-2: Official repositories from well-known organizations
- Rank 3-4: Verified community projects
- Rank 5-6: General open-source projects
- Rank 7-9: Lower confidence matches, may include forks or derivative works

**Score System** (0-1 scale):
- 0.9-1.0: Very high confidence match
- 0.8-0.89: High confidence match
- 0.7-0.79: Good match
- Below 0.7: Lower confidence (consider using `min_accepted_score` to filter)

**Match Order**:
Components are ordered by match quality, with order=1 being the best match. Use this to prioritize results when multiple components match the same folder.
