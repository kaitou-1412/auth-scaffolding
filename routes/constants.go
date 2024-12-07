package routes

var (
	ErrAuthSignup = map[string]string{
		"Request":  "Could not parse request data.",
		"Username": "Invalid input: Username should be 5-50 characters, alphanumeric.",
		"Email":    "Invalid input: Must be valid email format.",
		"Password": "Invalid input: Password should have minimum 8 characters. Password must contain at least one uppercase letter, one lowercase letter, one number and one special character.",
		"Conflict": "Username or email already exists.",
		"Create":   "Could not create user.",
	}
	SuccessAuthSignup = "User registered successfully"
	ErrAuthLogin      = map[string]string{
		"Request":       "Could not parse request data.",
		"Credentials":   "Invalid credentials.",
		"Inactive":      "Account locked or disabled.",
		"GenerateToken": "Could not authenticate user.",
	}
	SuccessAuthLogin = "Successful login"
	ErrAuthRefresh   = map[string]string{
		"Token": "Invalid refresh token.",
	}
	SuccessAuthRefresh = "Successful refresh"
	ErrProtectedRoute  = map[string]string{
		"Authorization": "You are not authorized to access this resource.",
	}
	SuccessProtectedRoute = "Successful access to protected resource."
)
