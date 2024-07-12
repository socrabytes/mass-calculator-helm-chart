# Mass Calculator Helm Chart

This repository contains a Helm chart for deploying a mass calculator application to a Kubernetes cluster. The application calculates the mass of an aluminium sphere and an iron cube based on dimensions provided through HTTP endpoints.

## Application Overview

The Mass Calculator application is a simple Go program that provides HTTP endpoints to calculate the mass of:
- An aluminium sphere based on its diameter (`/aluminium/sphere`)
- An iron cube based on its side length (`/iron/cube`)

The application listens on a port specified via a command-line argument, which is made configurable through the Helm chart.

## Project Structure
```shell
teamviewer_assignment/
├── src
│   ├── go.mod                   # Go module file for dependency management.
│   └── main.go                  # Go application source code.
├── docker
│   └── Dockerfile               # Dockerfile for building the application.
├── helm
│   ├── mass-calculator          # Directory containing the Helm chart.
│   │   ├── Chart.yaml           # Metadata about Helm chart
│   │   ├── README.md            # Helm chart specific README
│   │   ├── templates            # Kubernetes resource templates
│   │   │   ├── NOTES.txt        # Instructions shown after chart installation
│   │   │   ├── _helpers.tpl     # Template helpers
│   │   │   ├── deployment.yaml  # Deployment resource template
│   │   │   ├── ingress.yaml     # Ingress resource template
│   │   │   └── service.yaml     # Service resource template
│   │   ├── values-dev.yaml      # Development environment values
│   │   ├── values-prod.yaml     # Production environment values
│   │   └── values.yaml          # Default environment values
├── .gitignore
└── README.md

```

## Configuration
The following table lists the configurable parameters of the chart and their default values.


| Parameter          | Description                  | Default           |
| ------------------ | ---------------------------- | ----------------- |
| `replicaCount`     | Number of replicas           | `2`               |
| `image.repository` | Docker image repository      | `mass-calculator` |
| `image.tag`        | Docker image tag             | `latest`          |
| `image.pullPolicy` | Image pull policy            | `IfNotPresent`    |
| `service.type`     | Kubernetes Service type      | `ClusterIP`       |
| `service.port`     | Service port                 | `8080`            |
| `service.nodePort` | NodePort for development     | `30009`           |
| `env.port`         | Application port             | `8080`            |
| `resources`        | Resource requests and limits | `{}`              |

# Installation and Usage

### Prerequisites
Ensure you have the following installed before starting:
- Docker
- Kubernetes
- Helm

## Quick Start Guide

Clone the Repository

First, clone the repository to your local machine:
```shell
git clone https://github.com/socrates90/mass-calculator-helm-chart.git
cd mass-calculator-helm-chart
```

### Development Environment (NodePort)

To install the chart for development using `NodePort`:
```shell
helm install my-release ./mass-calculator --set service.port=9090 --set env.port=9090 -f mass-calculator/values-dev.yaml
```
This command installs the Helm chart with the release name my-release, setting the service and application port to 9090. The values-dev.yaml file contains configuration specific to the development environment.

After deployment, you can use the following endpoints to calculate the mass of geometrical shapes:
- **Aluminium Sphere**:
    ```shell
    curl "http://<minikube-ip>:30009/aluminium/sphere?dimension=<diameter>"
    ```
- **Iron Cube**:
    ```shell
    curl "http://<minikube-ip>:30009/iron/cube?dimension=<side-length>" 
    ```
Replace `<minikube-ip>`, `<diameter>`, and `<side_length>` with appropriate values.

### Production Environment (Ingress)
To install the chart for production using `ClusterIP` and `Ingress`:
```shell
helm install my-release ./mass-calculator --values ./mass-calculator/values-prod.yaml --set service.port=8080 --set env.port=8080
```
This command installs the Helm chart with the release name my-release, setting the service and application port to 8080. The values-prod.yaml file contains configuration specific to the production environment.

After deployment, you can use the following endpoints to calculate the mass of geometrical shapes:
- **Aluminium Sphere**:
    ```shell
    curl "http://<ingress-host>/aluminium/sphere?dimension=<diameter>"
    ```
- **Iron Cube**:
    ```shell
    curl "http://<ingress-host>/iron/cube?dimension=<side_length>"
    ```
Replace `<ingress-host>`, `<diameter>`, and `<side_length>` with appropriate values.

> **Note:** Ensure you update the `/etc/hosts` file to map the Minikube IP or Ingress IP to the desired host name if needed. 

### Uninstalling the Chart
To uninstall the release:
```shell
helm uninstall my-release
```

This command removes all the Kubernetes resources associated with the Helm release my-release.

## Notes
This chart configures resource requests and limits and includes liveness and readiness probes to ensure the application runs smoothly in a Kubernetes cluster.

---

For detailed information on the application itself and the Dockerfile used to containerize it, please refer to the respective source files (`src/main.go` and `docker/Dockerfile`).

---

## Contributing
Contributions are welcome! Please open an issue or submit a pull request for any changes.

## License
This project is licensed udner the MIT License - see the LICENSE file for details.