# Chat Service

## How to Run the Project

To get the project up and running on your local machine, follow these steps:

1. **Clone the repository:**
   First, clone the repository to your local machine using the following command:
   ```
   git clone https://github.com/jasanfarah/chat-service.git
   ```

2. **Start the database with Docker Compose:**
   Navigate to the root directory of the project and run the following command to start the database using Docker Compose:
   ```
   docker-compose up -d db
   ```
   This command starts the database service defined in your `docker-compose.yml` file in detached mode.

3. **Run the application:**
   Still in the root directory of the project, start the application by running:
   ```
   go run cmd/main.go
   ```
   This command compiles and runs the `main.go` file, starting your application.

By following these steps, you should have the project running locally on your machine.
