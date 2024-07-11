# Mass Calculator Helm Chart
This Helm chart deploys the Mass Calculator application to a Kubernetes cluster. The application calculates the mass of an aluminium sphere and an iron cube based on dimensions provided through HTTP endpoints.

## Application Overview
The Mass Calculator application is a simple Go program that provides HTTP endpoints to calculate the mass of:
- An aluminium sphere based on its diameter (`/aluminium/sphere`)
- An iron cube based on its side length (`/iron/cube`)

The application listens on a port specified via a command-line argument, which is made configurable through the Helm chart.


### Configuration
The following table lists the configurable parameters of the chart and their default values.

| Parameter             | Description                        | Default           |
|-----------------------|------------------------------------|-------------------|
| `replicaCount`        | Number of replicas                 | `2`               |
| `image.repository`    | Docker image repository            | `mass-calculator` |
| `image.tag`           | Docker image tag                   | `latest`          |
| `image.pullPolicy`    | Image pull policy                  | `IfNotPresent`    |
| `service.type`        | Kubernetes Service type            | `ClusterIP`       |
| `service.port`        | Service port                       | `8080`            |
| `service.nodePort`    | NodePort for development           | `30009`           |
| `env.port`            | Application port                   | `8080`            |
| `resources`           | Resource requests and limits       | `{}`              |

### Installing the Chart
#### For Development
To install the chart with the release name `my-release` for development using `NodePort`:

```sh
helm install my-release ./mass-calculator --set service.port=9090 --set env.port=9090 -f mass-calculator/values-dev.yaml
```

This command sets the service and application port to 9090.

#### For Production
To install the chart with the release name my-release for production using ClusterIP and Ingress:
```sh
helm install my-release ./mass-calculator --values ./mass-calculator/values-prod.yaml --set service.port=8080 --set env.port=8080
```
This command sets the service and application port to 8080.

#### Uninstalling the Chart
To uninstall the release:
```sh
helm uninstall my-release
```

## Using the Application

### Development Environment (NodePort)

After deploying the chart with the development configuration, you can use the following endpoints to calculate the mass of geometrical shapes:

- **Calculate the mass of an aluminium sphere**:
  ```sh
  curl "http://<minikube-ip>:30009/aluminium/sphere?dimension=<diameter>"
  ```
  Replace <minikube-ip> with the actual IP of your Minikube cluster and <diameter> with the diameter of the sphere.

- **Calculate the mass of an iron cube**:
  ```sh
  curl "http://<minikube-ip>:30009/iron/cube?dimension=<side_length>"
  ```
  Replace <minikube-ip> with the actual IP of your Minikube cluster and <side_length> with the side length of the cube.

### Production Environment (Ingress)
After deploying the chart with the production configuration, you can use the following endpoints to calculate the mass of geometrical shapes:

- **Calculate the mass of an aluminium sphere**:
  ```sh
  curl "http://<ingress-host>/aluminium/sphere?dimension=<diameter>"
  ```
  Replace <ingress-host> with the actual Ingress host configured in your values-prod.yaml and <diameter> with the diameter of the sphere.

- **Calculate the mass of an iron cube**:
  ```sh
  curl "http://<ingress-host>/iron/cube?dimension=<side_length>"
  ```
  Replace <ingress-host> with the actual Ingress host configured in your values-prod.yaml and <side_length> with the side length of the cube.

The result will be returned in grams (g).

>**Note:** Ensure you update the /etc/hosts file to map the Minikube IP or Ingress IP to the desired host name if needed.

## Notes
This chart configures resource requests and limits and includes liveness and readiness probes to ensure the application runs smoothly in a Kubernetes cluster.

---

For detailed information on the application itself and the Dockerfile used to containerize it, please refer to the respective source files (`main.go` and `Dockerfile`).