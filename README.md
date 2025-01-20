# 🎥 Home Theater

Welcome, and thank you for visiting the **Home Theater** project page! 🚀

## 🎯 Objective

The goal of this application is to provide an immersive and efficient experience for managing movie-related metadata, ratings, and more. Built with the power of **Go (Golang)**, this project follows a **microservices architecture** to ensure scalability, flexibility, and maintainability.

## 🏗️ Architecture

This application is divided into three core microservices, each with a distinct responsibility:

1. **MetaData Service**:  
   Handles all information related to movie metadata, such as title, genre, description, and release year.

2. **Movie Service**:  
   Manages movie records, including availability and details essential for user interactions.

3. **Rating Service**:  
   Focuses on user ratings, reviews, and aggregated scores for movies.

---

## 🚀 Getting Started

### 📂 Repository

Find the complete source code and setup instructions on GitHub:  
[Home Theater Repository](https://github.com/nikhil6392/Home-theater)

### 🔧 Prerequisites

To run this application, make sure you have the following installed:

- **Golang** (version 1.20+)
- **Docker** (for containerized microservices)
- **PostgreSQL** (or another database for backend data storage)

### 📦 Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/nikhil6392/Home-theater.git
   cd Home-theater
   ```

2. Build and run the services using Docker:
   ```bash
   docker-compose up --build
   ```

3. Access the application via the specified ports (check the `docker-compose.yml` for details).

---

## 🌟 Features

- **Scalable Microservices**:  
  Built using Go to handle high loads efficiently.

- **Containerized Deployment**:  
  Fully Dockerized for seamless deployment and testing.

- **RESTful APIs**:  
  Exposes endpoints for managing metadata, movies, and ratings.

- **Modular Design**:  
  Each service is independently deployable and testable.

---

## 🤝 Contributing

Contributions are welcome! Please feel free to submit issues, fork the repository, and create pull requests.

1. Fork the repo
2. Create a feature branch:
   ```bash
   git checkout -b feature/your-feature-name
   ```
3. Commit your changes:
   ```bash
   git commit -m "Add your feature description"
   ```
4. Push the branch:
   ```bash
   git push origin feature/your-feature-name
   ```
5. Open a pull request on GitHub.

---

## 📜 License

This project is licensed under the **MIT License**. See the [LICENSE](https://github.com/nikhil6392/Home-theater/blob/main/LICENSE) file for more information.

---

Happy coding! 🎬  
_Developed with ❤️ by [Nikhil6392](https://github.com/nikhil6392)._
