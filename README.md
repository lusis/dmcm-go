# dmcm-go
A Go library for interacting with the DMCM (Enstratius/Enstratus) API.

There isn't much to see here yet. The code is pretty ugly and this is mostly a learning exercise.

## Building
- Install go 1.1.1
- Clone and cd into the directory
- export GOPATH=`pwd`
- `make clean all`

## Environment variables
Like all the other libraries and tools for interacting with the DMCM API, the two REQUIRED environment variables:

- `ES_ACCESS_KEY`
- `ES_SECRET_KEY`

Additionally, there are a few other variables for interacting with DMCM installations other than our hosted version:

- `ES_ENDPOINT` - e.g. `http://vagrant.vm:15000` - no trailing slash. If unset, will use `https://api.enstratus.com`
- `ES_NOVERIFY_SSL` - Setting this to anything will disable SSL validation.
- `ES_DETAIL` - Control the level of detail in API responses: `none`, `basic` or `extended`

## Usage
As a library there really isn't much here yet. The main usage right now is a statically compiled API dump tool similar to `es-dump.py` included in [mixcoatl](https://github.com/enstratus/mixcoatl)

Examples:
```
[user@host]: ES_ACCESS_KEY=ABCDEFG ES_SECRET_KEY=12345 bin/dmcm-cli geography/Cloud/1

(map[string]interface {}) {
 (string) "clouds": ([]interface {}) {
  (map[string]interface {}) {
   (string) "privateCloud": (bool) false,
   (string) "cloudProviderConsoleURL": (string) "http://aws.amazon.com",
   (string) "computeDelegate": (string) "org.dasein.cloud.aws.AWSCloud",
   (string) "cloudId": (float64) 1,
   (string) "computeX509KeyLabel": (string) "AWS_PRIVATE_KEY",
   (string) "name": (string) "Amazon Web Services",
   (string) "computeEndpoint": (string) "https://ec2.us-east-1.amazonaws.com",
   (string) "computeSecretKeyLabel": (string) "AWS_SECRET_ACCESS_KEY",
   (string) "status": (string) "ACTIVE",
   (string) "computeX509CertLabel": (string) "AWS_CERTIFICATE",
   (string) "computeAccountNumberLabel": (string) "AWS_ACCOUNT_NUMBER",
   (string) "documentationLabel": (interface {}) <nil>,
   (string) "cloudProviderName": (string) "Amazon",
   (string) "computeAccessKeyLabel": (string) "AWS_ACCESS_KEY",
   (string) "cloudProviderLogoURL": (string) "/clouds/aws.gif"
  }
 }
}
```

```
[user@host]: ES_DETAIL=none ES_ACCESS_KEY=ABCDEFG ES_SECRET_KEY=12345 bin/dmcm-cli geography/Cloud/1

(map[string]interface {}) {
 (string) "clouds": ([]interface {}) {
  (map[string]interface {}) {
   (string) "cloudId": (float64) 1
  }
 }
}
```

```
[user@host] ES_ENDPOINT="https://vagrant.vm:15433" ES_NOVERIFY_SSL=true ES_ACCESS_KEY=ABCDEFG ES_SECRET_KEY=12345 bin/dmcm-cli geography/Cloud/1

(map[string]interface {}) {
 (string) "clouds": ([]interface {}) {
  (map[string]interface {}) {
   (string) "privateCloud": (bool) false,
   (string) "cloudProviderConsoleURL": (string) "http://aws.amazon.com",
   (string) "computeDelegate": (string) "org.dasein.cloud.aws.AWSCloud",
   (string) "cloudId": (float64) 1,
   (string) "computeX509KeyLabel": (string) "AWS_PRIVATE_KEY",
   (string) "name": (string) "Amazon Web Services",
   (string) "computeEndpoint": (string) "https://ec2.us-east-1.amazonaws.com",
   (string) "computeSecretKeyLabel": (string) "AWS_SECRET_ACCESS_KEY",
   (string) "status": (string) "ACTIVE",
   (string) "computeX509CertLabel": (string) "AWS_CERTIFICATE",
   (string) "computeAccountNumberLabel": (string) "AWS_ACCOUNT_NUMBER",
   (string) "documentationLabel": (interface {}) <nil>,
   (string) "cloudProviderName": (string) "Amazon",
   (string) "computeAccessKeyLabel": (string) "AWS_ACCESS_KEY",
   (string) "cloudProviderLogoURL": (string) "/clouds/aws.gif"
  }
 }
}
```
