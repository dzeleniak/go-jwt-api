echo "Get: GORM"

go get -u gorm.io/gorm

echo "Get: GORM - postgres"

go get gorm.io/driver/postgres

echo "Get: gin"

go get -u github.com/gin-gonic/gin

echo "Get: bcrypt"

go get -u golang.org/x/crypto/bcrypt

echo "Get: JWT"

go get -u github.com/golang-jwt/jwt/v4

echo "Get: GoDotEnv"

go get github.com/joho/godotenv

echo "Get: CompileDaemon"

go get github.com/githubnemo/CompileDaemon

echo "Install: CompileDaemon"

go install github.com/githubnemo/CompileDaemon
