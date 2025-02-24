# BinVault

**BinVault** is an open-source file storage and compression service designed for cloud-native environments. It efficiently compresses, stores, and serves images and other files, making it ideal for developers, hosting providers, and edge computing platforms.

[![forthebadge](https://forthebadge.com/images/featured/featured-gluten-free.svg)](https://forthebadge.com)
[![forthebadge](https://forthebadge.com/images/featured/featured-built-with-love.svg)](https://forthebadge.com)


![GitHub release (latest by date)](https://img.shields.io/github/v/release/kalevski/binvault?style=for-the-badge)
![Docker Pulls](https://img.shields.io/docker/pulls/kalevski/binvault?style=for-the-badge)
![GitHub issues](https://img.shields.io/github/issues/kalevski/binvault?style=for-the-badge)
![GitHub](https://img.shields.io/github/license/kalevski/binvault?style=for-the-badge)

## 🚀 Features

- **Image Compression**: Supports PNGQuant, MozJPEG, and WebP.
- **Efficient Storage**: Organizes files in "buckets" with customizable permissions.
- **Built-in Nginx Server**: Serves compressed files with caching and gzip support.
- **REST API**: Fully functional API for uploading, retrieving, and managing files.
- **Worker Queue**: Asynchronous compression using worker pools.
- **SQLite Database**: Tracks files and metadata efficiently.
- **Cloud-Native Ready**: Supports Kubernetes, Docker, and containerized deployments.

---

## 📦 Installation

### **1️⃣ Install via Docker (Recommended)**
```sh
# Clone the repository
git clone https://github.com/polovni-mk/binvault.git && cd binvault

# Build and run the container
docker build -t binvault .
docker run -p 8080:8080 -v $(pwd)/data:/app/data binvault
```

### **2️⃣ Manual Installation (Go & SQLite)**
```sh
# Install dependencies
sudo apt update && sudo apt install sqlite3

# Install Go Modules
go mod tidy

# Run the server
go run main.go
```

---

## 📌 Usage

### **Upload a File**
```sh
curl -X POST -F "file=@image.png" http://localhost:8080/api/upload
```

### **Retrieve a Compressed File**
```sh
curl http://localhost:8080/api/files/image_compressed.png
```

### **List All Files in a Bucket**
```sh
curl http://localhost:8080/api/bucket/default
```

| HTTP Method | Endpoint         | Description                      | Request Body | Response Body |
|-------------|------------------|----------------------------------|--------------|---------------|
| GET         | /api/resource    | Retrieves a list of resources    | None         | JSON array of resources |
| GET         | /api/resource/:id| Retrieves a specific resource    | None         | JSON object of the resource |
| POST        | /api/resource    | Creates a new resource           | JSON object  | JSON object of the created resource |
| PUT         | /api/resource/:id| Updates an existing resource     | JSON object  | JSON object of the updated resource |
| DELETE      | /api/resource/:id| Deletes a specific resource      | None         | JSON object with deletion status |

---

## 🔒 Authentication & Security
BinVault supports **token-based authentication**. To use it, include an `Authorization` header:
```sh
curl -H "Authorization: Bearer YOUR_API_TOKEN" http://localhost:8080/api/files
```

---

## 🛠 Configuration
BinVault uses environment variables for customization:
```sh
# Port configuration
PORT=8080

# Compression settings
PNG_QUALITY=65-80
JPEG_QUALITY=75
```

---

## 🤝 Contributing
We welcome contributions! Follow these steps:
1. Fork the repository
2. Create a feature branch (`git checkout -b feature-name`)
3. Commit changes (`git commit -m "Added new feature"`)
4. Push to your fork and submit a PR

For detailed contribution guidelines, see [CONTRIBUTING.md](CONTRIBUTING.md).

---

## 📜 License
BinVault is released under the **Apache 2.0 License**.

---

## 📧 Contact & Community
- **GitHub Issues**: [Report Bugs](https://github.com/polovni-mk/binvault/issues)
- **Discussions**: [Join the Community](https://github.com/polovni-mk/binvault/discussions)
- **Twitter**: [Follow for Updates](https://twitter.com/binvault)

---

🚀 **Start using BinVault today for efficient and optimized file storage!**
