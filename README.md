# Learning Zarf

The purpose of this repo was to gain experience with zarf. For that purpose I created an example hello world app which consists of a backend which provides a greeting and a frontend which displays that greeting.  
The version 1.0.0 of this app has a hardcoded greeting. In version 1.1.0 the greeting can be configured in the backend.  

## Setup Cluster for air-gapped deployment
I used minicube for my tests. I did set up minikube "cluster" using the following command:
```shell
minikube start --cpus 4 --memory 8192
minikube addons enable ingress
```
To prepare the minikube cluster for air-gapped deployment I did run the following command:
```shell
zarf init
```
I did skip the k3s installation, but installed the registry and git-repo.

## Building zarf package
To build a zarf package use the Makefile:
```shell
make package
```
This will create a `zarf-package-hello-world-app-amd64-<version>.tar.zst` in the project root.

## Deploy zarf package
To deploy the zarf package to the prepared cluster run the following command:
```shell
zarf package deploy zarf-package-hello-world-app-amd64-*.tar.zst --set ENVIRONMENT_VAR_INGRESS_HOST=<your-ingress-host-name>
```
The ingress host name is required, as the hello world app wants to register routes to the ingress.

## Remove zarf package
To remove the package again run:
```shell
zarf package remove zarf-package-hello-world-app-amd64-*.tar.zst
```

## Inspecting a package
To get the definition of the zarf package run
```shell
zarf package inspect definition zarf-package-hello-world-app-amd64-*.tar.zst
```

Or to get the variables directly run
```shell
zarf package inspect definition zarf-package-hello-world-app-amd64-*.tar.zst | sed 's/\x1b\[[0-9;]*m//g' | yq '.components[0].charts[0].variables'
```

To extract the SBOM call
```shell
zarf package inspect sbom zarf-package-hello-world-app-amd64-*.tar.zst
```

## This to investigate
- signing of zarf packages
- overwrite image versions in zarf definition and helm charts during build
