# Curl commands to query SCANOSS PAPI
## Export control
### Get crypto algorithms in range

```
curl --location 'https://api.scanoss.com/api/v2/cryptography/algorithmsInRange' \
--header 'Content-Type: application/json' \
--header "X-Api-Key: <Your access token>" \
--data '{
    "purls": [
        {
            "purl": "pkg:maven/org.jetbrains.kotlin/kotlin-gradle-plugin",
            "requirement": ">0.5"
        }
    ]
}'
```

### Get Algorithms

```
curl --location 'https://api.scanoss.com/api/v2/cryptography/algorithms' \
--header 'Content-Type: application/json' \
--header "X-Api-Key: <Your access token>" \
--data '{"purls": [
    {
      "purl": "pkg:github/torvalds/linux",
      "requirement": "v4.0"
    }, 
     {
      "purl": "pkg:github/scanoss/engine",
      "requirement": "4.3"
    }
  ]
}'
```

## Provenance
### Get Countries of origin

```
curl --location 'https://api.scanoss.com/api/v2/provenance/countries' \
--header 'Content-Type: application/json' \
--header "X-Api-Key: <Your access token>" \
--data '{
  "purls": [
    {
      "purl": "pkg:github/scanoss/engine"
     
    }
  ]
}'
```

## Components
### search for a component
```
curl --location 'https://api.scanoss.com/api/v2/components/search' \
--header 'Content-Type: application/json' \
--header "X-Api-Key: <Your access token>" \
--data '{"search" : "scanoss", "limit": 10 }'
```

### Search for versions of a component
```
curl --location 'https://api.scanoss.com/api/v2/components/versions' \
--header 'Content-Type: application/json' \
--header "X-Api-Key: <Your access token>" \
--data '{"purl": "pkg:github/scanoss/scanoss.js" }'
```
## Dependencies
```
curl --location 'https://api.scanoss.com/api/v2/dependencies/dependencies' \
--header 'Content-Type: application/json' \
--header "X-Api-Key: <Your access token>" \
--data '{"files": [{ "purls": [ { "purl":"pkg:npm/wrap-ansi-cjs"} ]  } ]}'
```
## Vulnerabilities
### Get CPEs
```
curl --location 'https://api.scanoss.com/api/v2/vulnerabilities/cpes' \
--header 'Content-Type: application/json' \
--header "X-Api-Key: <Your access token>" \
--data-raw '{"purls": [
    {
    "purl": "pkg:npm/lodash@4.17.21"
      
    }
  ]
}'
```
### Get vulnerabilities
```
curl --location 'https://api.scanoss.com/api/v2/vulnerabilities/vulnerabilities' \
--header 'Content-Type: application/json' \
--header "X-Api-Key: <Your access token>" \
--data-raw '{"purls": [
      {
     "purl": "pkg:npm/lodash@4.17.21"
    },
    {
      "purl": "pkg:npm/@ava/typescript"
    }
  ]
}'
```


## Ignoring SSL certificate errors
To ignore SSL certificate errors when using curl, you can use the -k or --insecure option. This allows curl to perform "insecure" SSL connections. 
The command would look like this:

```curl -k https://example.com```
or 
```curl --insecure https://example.com```

### Example error when not using --insecure

If you don't use the --insecure option and encounter a self-signed certificate or an invalid SSL certificate, you might see an error like this:

```
curl: (60) SSL certificate problem: self-signed certificate in certificate chain
More details here: https://curl.se/docs/sslcerts.html

curl failed to verify the legitimacy of the server and therefore could not
establish a secure connection to it. To learn more about this situation and
how to fix it, please visit the web page mentioned above.
```

## Using Proxies
To use a proxy with curl, you can use the -x or --proxy option. The general format is:
```
curl -x [protocol://]host[:port] URL
```

* Using HTTP Proxy: `curl -x http://proxy.example.com:8080 https://destination.com`
* Using HTTPS Proxy: `curl -x https://proxy.example.com:8080 https://destination.com`
* Specifying a username and password for the proxy: `curl -x http://username:password@proxy.example.com:8080 http://destination.com`

## Using Proxy and SSL interception
When working in an environment with SSL interception (common in corporate settings), y
ou often need to combine the use of a proxy with the ability to ignore SSL certificate errors. 
Here's how you can do this:

1. Obtain the SSL certificate of your proxy server (provided by your IT department).
2. Use the following curl command:

```
curl -x https://proxy.example.com:8080 \
     --proxy-cacert /path/to/proxy/certificate.crt \
     --insecure \
     https://destination.com
```
