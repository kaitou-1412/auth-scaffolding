# Auth Scaffolding

This backend application serves as a foundational skeleton for building servers with authentication and authorization, featuring Role-Based Access Control (RBAC). It is designed with extensibility in mind, enabling seamless integration of custom features and modifications.

## Running the Application Using Docker

Follow the steps below to set up and run the application using Docker:

**Step 1: Clone the Repository**  
Clone this repository to your local machine and ensure that a `.env` file is present at the root of the repository. A sample configuration file (`sample.env`) is provided as a reference to help you create your `.env` file.

**Step 2: Build and Start the Application**  
Open a terminal, navigate to the project directory, and execute the following command:

```Shell
docker compose up --build
```

**Step 3: Test the Application**  
Once the application is up and running, you can test it by interacting with the available endpoints to ensure it operates as expected.

## API Documentation

### 1. **User Sign-Up**

**URL**: `/api/auth/signup`  
**Method**: `POST`  
**Description**: Register a new user.

#### Request Body:

```json
{
  "name": "string",
  "username": "string",
  "email": "string",
  "password": "string"
}
```

#### Responses:

- 201 Created: User registered successfully.
- 400 Bad Request: Invalid input (e.g., username, email, or password validation failure).
- 409 Conflict: Username or email already exists.

### 2. **User Login**

**URL**: `/api/auth/login`  
**Method**: `POST`  
**Description**: Authenticate a user and return access tokens.

#### Request Body:

```json
{
  "username": "string",
  "password": "string"
}
```

#### Responses:

- 200 OK: Login successful. Returns access and refresh tokens.
- 400 Bad Request: Missing or invalid request body.
- 401 Unauthorized: Invalid credentials.
- 403 Forbidden: Account is locked or disabled.

### 3. **Token Refresh**

**URL**: `/api/auth/refresh`  
**Method**: `POST`  
**Description**: Refresh the access token using a valid refresh token.

#### Headers:

`Authorization`: `Bearer <refresh_token>`

#### Responses:

- 200 OK: Token refreshed successfully.
- 401 Unauthorized: Invalid or expired refresh token.

### 4. **Protected Resource Access**

**URL**: `/api/protected`  
**Method**: `GET`  
**Description**: Access a resource that requires authentication.

#### Headers:

`Authorization`: `Bearer <access_token>`

#### Responses:

- 200 OK: Successful access to the protected resource.
- 401 Unauthorized: Access token is missing or invalid.

---

Let me know if any further improvements are required, or feel free to submit a pull request!
