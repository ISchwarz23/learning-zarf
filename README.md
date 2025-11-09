# Learning Zarf

The purpose of this repo was to gain experience with zarf. For that purpose I created an example hello world app which consists of a backend which provides a greeting and a frontend which displays that greeting.  
The version 1.0.0 of this app has a hardcoded greeting. In version 1.1.0 the greeting can be configured in the backend.  

## What is Zarf?
<img align="right" alt="zarf logo" src="https://raw.githubusercontent.com/zarf-dev/zarf/main/site/src/assets/zarf-logo.png"  height="256" />

[Zarf](https://www.zarf.dev) claims to eliminates the [complexity of airgap software delivery](https://www.itopstimes.com/contain/air-gap-kubernetes-considerations-for-running-cloud-native-applications-without-the-cloud/) for Kubernetes clusters and cloud-native workloads using a declarative packaging strategy to support DevSecOps in offline and semi-connected environments.  
It is a free open source tool that enables continuous software delivery on systems that are disconnected from the internet.

## Instructions

### Setup Cluster for Air-Gapped Deployment
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

### Building a Zarf Package
To build a zarf package use the Makefile:
```shell
make package
```
This will create a `zarf-package-hello-world-app-amd64-<version>.tar.zst` in the project root.

### Deploing a Zarf Package
To deploy the zarf package to the prepared cluster run the following command:
```shell
zarf package deploy zarf-package-hello-world-app-amd64-*.tar.zst --set ENVIRONMENT_VAR_INGRESS_HOST=<your-ingress-host-name>
```
The ingress host name is required, as the hello world app wants to register routes to the ingress.

### Removing a Zarf Package
To remove the package again run:
```shell
zarf package remove zarf-package-hello-world-app-amd64-*.tar.zst
```

### Inspecting a Zarf Package
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

## Topics to Investigate
- signing of zarf packages
- overwrite image versions in zarf definition and helm charts during build
