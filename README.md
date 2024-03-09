# Dockerized Go Web App with NGINX Load Balancer

This project demonstrates how to Dockerize a Go web application and use NGINX as a load balancer to distribute incoming traffic across multiple instances of the application.

## Prerequisites

- Docker
- Docker Compose

## Getting Started

1. Clone this repository:

    ```bash
    git clone https://github.com/yourusername/your-repo.git
    cd your-repo
    ```

2. Update the Docker Compose file (`docker-compose.yml`) and NGINX configuration file (`nginx.conf`) as per the provided examples.

3. Build and run the Docker services:

    ```bash
    docker-compose up -d
    ```

4. Access the application in your browser at `http://localhost`.

## Configuration

### Docker Compose

- The `docker-compose.yml` file defines services for the Go web app, Redis database, and NGINX load balancer.
- Update the services section to match your project's requirements.
- Ensure proper networking is set up to allow communication between services.

### NGINX Configuration

- The `nginx.conf` file configures NGINX as a reverse proxy and load balancer.
- Update the `upstream` block to include the addresses of your Go web app instances.
- Customize other NGINX settings as needed for your application.

## Usage

- Scale the Go web app service horizontally by increasing the number of container instances.
- Monitor NGINX logs and performance to optimize load balancing and server health.

## Contributing

Contributions are welcome! Fork this repository, make your changes, and submit a pull request.
