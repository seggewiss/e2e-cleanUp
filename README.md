# Shopware e2e cleanup
This golang tool allows you to execute the Shopware e2e plugin suites 8005 calls in any shopware directory you want.

# Installation
```
$ git clone git@github.com:seggewiss/e2e-cleanUp.git
```

# Usage
```
$ go mod download

$ make build

$ ./proxy -path=/path/to/your/shopware/instance
```