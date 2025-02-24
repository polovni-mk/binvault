# BinVault

BinVault is an open-source service designed to compress and store images efficiently. It is built to be deployed in any environment and includes a built-in static server (nginx) with caching capabilities.

## Features

- **Image Compression**: Supports various image compression algorithms.
- **Storage**: Efficiently stores compressed images.
- **Built-in Nginx**: Built-in nginx for serving compressed files.
- **Buckets**: The files can be separated into multiple buckets with different visibility (public | private).

## Getting Started

### Running

1. Build the Docker image:
    ```sh
    docker build -t binvault .
    ```

2. Run the Docker container:
    ```sh
    docker run -d -p 8080:8080 binvault
    ```

3. Access the service:
    Open your browser and navigate to `http://localhost:8080` to start using BinVault.