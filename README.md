# JWT API Example Project
This project illustrates how to create a REST API to handle JWT token authentication.

## Technologies
- Gin => HTTP Framework
- Postgres => Database
- GORM => Object Relational Manager

## MAIN
### Init()
The init function handles the connection and creation of the database, as well as loading of environment variables.
### Main() 
The main function creates a GIN instance with the required routes from our User Controller and runs the engine.

## User Controller
The user controller is comprised of three main functions.
### Signup
- Pull email and password off the request body
    - create a struct with required fields and bind to the gin context
- Hash the password received in the body using bCrypt.GenerateFromPassword()
- Create a user in the database 
- Send response

### Login
- Get email and password from the request body by binding a struct to the context
- Check the database for the requested user
- Compare the hash with the password sent using bcrypt.CompareHashAndPassword()
- Generate a JWT Token.
    - Set sub to the User ID 
    - Set exp to 30 days after now
        - ```time.Now().Add(time.Hour * 24 * 30).Unix()```
- Sign the token and encode using a secret
    - ```token.SignedString([]byte(os.Getenv("SECRET")))```
- Set the Authorization cookie equal to the token string
- Send the token back as a response

### Validate
The validation function has a requireAuth middleware attached and will not successfully run unless a logged in user is found

- Get the user from gin's context
- return a message with the user object

## Middleware
### Require Auth
- Get the cookie from gin context (```c.Cookie("Authorization")```)
- Decode the token using jwt.Parse(tokenString, keyFunc)
    - Key func should return the SECRET environment variable as a byte array
- Validate the token by checking the claims are ok and the token is valid
    - check expiration
    - Check that the user found under token sub is valid
    - attach the user to the request
    - Continue the pipeline


## Libraries Used
https://gorm.io/
https://gin-gonic.com/docs/quickstart/
https://github.com/joho/godotenv
https://pkg.go.dev/github.com/golang-jwt/jwt
https://github.com/githubnemo/CompileDaemon
https://pkg.go.dev/golang.org/x/crypto/bcrypt