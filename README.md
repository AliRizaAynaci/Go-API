# Music API

This project provides a simple API for managing music albums. The API allows you to perform CRUD operations on albums.

## API Documentation

To view the API documentation, you can import the `api.yaml` file into tools like Postman or Swagger Editor.

## Installation

1. Clone the project:

    ```bash
    git clone https://github.com/user/project.git
    ```

2. Navigate to the project directory:

    ```bash
    cd project
    ```

3. Install the required dependencies:

    ```bash
    go mod tidy
    ```

4. Run the project:

    ```bash
    go run main.go
    ```

5. Open your browser and go to [http://localhost:8080](http://localhost:8080) to access the API.

## Usage

You can send HTTP requests to the following routes to interact with the API:

- `GET /albums`: Retrieve all albums.
- `POST /albums`: Add a new album.
- `GET /albums/{id}`: Retrieve a specific album.
- `PUT /albums/{id}`: Update a specific album.
- `DELETE /albums/{id}`: Delete a specific album.

## Contributing

If you'd like to contribute to this project, please open an Issue or submit a Pull Request.