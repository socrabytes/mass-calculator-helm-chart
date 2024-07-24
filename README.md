# Mass Calculator Helm Chart

## Introduction

The Mass Calculator is a Kubernetes-ready application that demonstrates proficiency in Go programming, Docker containerization, and Helm chart deployment. Originally conceived as a response to a DevOps challenge, this project has been expanded to showcase best practices in cloud-native application development.

Key features and technologies demonstrated:

- Efficient Go programming for a practical mass calculation application
- Docker containerization optimized for size and security
- Helm chart creation for flexible Kubernetes deployment
- Implementation of health checks and readiness probes
- Consideration for both development and production environments

This project serves as a compact yet comprehensive example of modern DevOps practices, from code to container to cloud deployment.

## Application Overview

The Mass Calculator is a Go-based service that calculates the mass of geometric shapes based on their dimensions and material density. It provides HTTP endpoints to compute the mass of:

* Aluminium sphere based on its diameter (`/aluminium/sphere`)
* Iron cube based on its side length (`/iron/cube`)

Key technical features:

1. **Flexible Port Configuration**: 
   - The application accepts a port number as a command-line argument.
   - When deployed via Helm, the port is configurable as an environment variable, set through Helm values.

2. **RESTful API**: Utilizes HTTP GET requests with query parameters for dimension input:
   * **Endpoints**:
      * `/aluminium/sphere` for calculating the mass of an aluminium sphere.
      * `/iron/cube` for calculating the mass of an iron cube.
   * **Query Parameter**: `dimension` (required, floating-point number).
   * **Success Response**: Returns the mass in grams, rounded to two decimal places.
   * **Error Response**: Returns HTTP status code 400 (Bad Request) for invalid or missing `dimension`.

3. **Precise Calculations**: Implements accurate formulas for volume and mass calculations, with results provided in grams.

4. **Error Handling**: Properly handles bad requests, returning appropriate HTTP status codes.

5. **Health Checks**: Implements `/healthz` and `/readyz` endpoints for liveness and readiness probes, enhancing reliability in Kubernetes deployments.

6. **Containerization**:
   - Multi-stage build process for optimized image size (< 100MB compressed).
   - Base image: Alpine Linux for minimal footprint.
   - Non-root user execution for enhanced security.
   - Environment variable `PORT` for flexible port configuration.
   - Dockerfile features:
     * Go dependencies management and build in the first stage.
     * Only the compiled binary copied to the final stage.
     * Creation of a non-root user (`appuser`) for running the application.
     * `EXPOSE` instruction documenting the default port.
     * `ENTRYPOINT` using the `PORT` environment variable to configure the application.

7. **Helm Deployment**: 
   - Utilizes a Helm chart for Kubernetes deployment, with configurable values for enhanced flexibility.
   - Supports environment-specific configurations through separate value files (`values.yaml`, `values-dev.yaml`, `values-prod.yaml`).
   - Allows for dynamic configuration adjustments using Helm's `--set` flags during deployment.

The application is designed to be highly configurable, allowing for easy adjustment of deployment parameters such as replica count, port settings, and resource allocation in both development and production environments.

## Technologies Used

- Go: For the main application logic
- Docker: For containerization
- Kubernetes: As the target deployment platform
- Helm: For packaging and deploying the application to Kubernetes
- Alpine Linux: As the base image for the Docker container

## Project Structure
```shell
mass-calculator-helm-chart/
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
├── src
│   ├── go.mod                   # Go module file for dependency management.
│   └── main.go                  # Go application source code.
├── .gitignore
├── Dockerfile                   # Dockerfile for building the application.
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

## Installation and Usage

### Prerequisites
Ensure you have the following installed before starting:
- **Docker**
- **Kubernetes Cluster**: Ensure you have a Kubernetes cluster set up.
    - For production, this guide assumes you are using Amazon EKS.
    - For development, you can use Minikube or any other local Kubernetes setup.
- **Helm**

#### EKS Cluster Setup (Production Only)
If you don't have an EKS cluster set up, you can create one using the following commands:

1. **Create an EKS Cluster**:
   ```sh
   eksctl create cluster --name demo-eks --region us-east-1 --nodegroup-name my-nodes --node-type t3.small --managed --nodes 2
   ```
2. **Update kubeconfig**:
   ```sh
   aws eks --region us-east-1 update-kubeconfig --name demo-eks
   ```

### Quick Start Guide

#### Clone the Repository

First, clone the repository to your local machine:
```shell
git clone https://github.com/socrates90/mass-calculator-helm-chart.git
cd mass-calculator-helm-chart
```

#### Development Environment (NodePort)

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

> **Note:** Ensure you update the `/etc/hosts` file to map the Minikube IP or Ingress IP to the desired host name if needed. 

#### Production Environment (LoadBalancer)
To install the chart for production using `ClusterIP` and `Ingress`:
```shell
helm install my-release ./mass-calculator --values ./mass-calculator/values-prod.yaml --set service.port=8080 --set env.port=8080
```
This command installs the Helm chart with the release name my-release, setting the service and application port to 8080. The values-prod.yaml file contains configuration specific to the production environment.

#### Accessing the Application in Production
1. **Get the External IP**:
   ```sh
   kubectl get svc -n default
   ```

2. **Test Application Endpoints**:
   - **Aluminium Sphere**:
      ```shell
      curl "http://<ingress-host>:8080/aluminium/sphere?dimension=<diameter>"
      ```
   - **Iron Cube**:
      ```shell
      curl "http://<ingress-host>:8080/iron/cube?dimension=<side_length>"
      ```
   Replace `<ingress-host>` with the external IP or domain of your LoadBalancer to calculate the mass of geometrical shapes:

   Replace  `<diameter>`, and `<side_length>` with appropriate values.

### Managing the Deployment
#### Uninstalling the Chart: 
To uninstall the release:
```shell
helm uninstall my-release
```

This command removes all the Kubernetes resources associated with the Helm release my-release.

#### Terminating EKS Cluster
For production environments using EKS, terminate the cluster with the following commands:

1. **Delete EKS Cluster**: 
   ```sh
   eksctl delete cluster --name demo-eks --region us-east-1
   ```
2. **Verify Deletion**:
   ```sh
   eksctl get cluster --name demo-eks --region us-east-1
   ```

## Notes
This chart configures resource requests and limits and includes liveness and readiness probes to ensure the application runs smoothly in a Kubernetes cluster.

---

For detailed information on the application itself and the Dockerfile used to containerize it, please refer to the respective source files (`src/main.go` and `Dockerfile`).

---

## Contributing
Contributions are welcome! Please open an issue or submit a pull request for any changes.

## License
This project is licensed udner the MIT License - see the LICENSE file for details.
