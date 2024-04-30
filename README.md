## Local Setup Instructions

This guide provides instructions on how to set up and run the server for the "Gin Library" project on a local machine.

### Prerequisites

- [Go](https://golang.org/dl/) (version 1.14 or higher)
- [Git](https://git-scm.com/downloads) (for cloning the repository)

### Installation

1. **Clone the Repository**

   Use Git to clone the repository to your local machine:

   ```bash
   git clone https://github.com/yourusername/ginLibrary.git
   cd ginLibrary
   ```

2. **Set Environment Variables**

   Create a `.env` file in the root directory of the project and set the necessary environment variables:

   ```plaintext
   JWT_SECRET_KEY=your_secret_key_here
   TOKEN_EXPIRATION_HOURS=72
   ```

   Replace `your_secret_key_here` with a strong secret key.

3. **Install Dependencies**

   Run the following command to download and install the project dependencies:

   ```bash
   go mod tidy
   ```

### Running the Server

To start the server, run the following command in the root directory of the project:

```bash
go run main.go
```

The server will start running on `http://localhost:8080`.

### Interacting with the API

#### Creating a JWT Token

Before interacting with the API, you need to authenticate and obtain a JWT token. Use this curl command to login:

```bash
curl --location --request POST 'http://localhost:8080/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "Username": "admin",
    "Password": "admin123"
}'
```

This will return a JWT token that you need to use as a Bearer token for subsequent requests.

#### Using the JWT Token

To access authenticated routes, include the JWT token obtained from the login step in the Authorization header of your requests:

##### Get All Books

```bash
curl --location --request GET 'http://localhost:8080/home' \
--header 'Authorization: Bearer <your_token_here>'
```

##### Add a Book

```bash
curl --location --request POST 'http://localhost:8080/addBook' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer <your_token_here>' \
--data-raw '{
    "bookName": "New Book",
    "author": "John Doe",
    "publicationYear": 2021
}'
```

##### Delete a Book

```bash
curl --location --request DELETE 'http://localhost:8080/deleteBook' \
--header 'Authorization: Bearer <your_token_here>' \
--data-raw '{
    "bookName": "New Book"
}'
```

### Note
- Replace `<your_token_here>` with the JWT token obtained from the login response.
- Ensure that your `.env` file is properly configured as mentioned in the Installation section.

---