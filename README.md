# Golang API Mini Project

This project is a simple API built with Golang. It demonstrates the basic structure and functionality of a RESTful API.

## Table of Contents
- [Installation](#installation)
- [Usage](#usage)
- [Endpoints](#endpoints)
- [Contributing](#contributing)
- [License](#license)

## Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/yourusername/golang-api-mini-project.git
    ```
2. Navigate to the project directory:
    ```sh
    cd golang-api-mini-project
    ```
3. Install the dependencies:
    ```sh
    go mod tidy
    ```

## Usage

1. Run the application:
    ```sh
    go run main.go
    ```
2. The API will be available at `http://localhost:8080`.

## Endpoints

### GET /items
- Description: Retrieve a list of items.
- Response:
    ```json
    [
        {
            "id": 1,
            "name": "Item 1",
            "price": 100
        },
        {
            "id": 2,
            "name": "Item 2",
            "price": 200
        }
    ]
    ```

### POST /items
- Description: Create a new item.
- Request Body:
    ```json
    {
        "name": "New Item",
        "price": 150
    }
    ```
- Response:
    ```json
    {
        "id": 3,
        "name": "New Item",
        "price": 150
    }
    ```

## Contributing

Contributions are welcome! Please follow these steps:
1. Fork the repository.
2. Create a new branch (`git checkout -b feature-branch`).
3. Make your changes.
4. Commit your changes (`git commit -m 'Add some feature'`).
5. Push to the branch (`git push origin feature-branch`).
6. Open a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.