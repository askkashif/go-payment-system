module payment-system-four // Module name

go 1.22.0 // Go version required

// Direct dependencies required by the application
require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // JWT library for handling JSON Web Tokens
	github.com/gin-contrib/cors v1.7.2 // CORS middleware for Gin framework
	github.com/gin-gonic/gin v1.10.0 // Gin web framework for building APIs in Go
	github.com/joho/godotenv v1.5.1 // GoDotEnv library for loading environment variables from a .env file
	gorm.io/driver/postgres v1.5.7 // PostgreSQL driver for GORM ORM
	gorm.io/gorm v1.25.10 // GORM ORM for database interactions
)

// Indirect dependencies required by the direct dependencies
require (
	github.com/bytedance/sonic v1.11.6 // Indirect dependency
	github.com/bytedance/sonic/loader v0.1.1 // Indirect dependency
	github.com/cloudwego/base64x v0.1.4 // Indirect dependency
	github.com/cloudwego/iasm v0.2.0 // Indirect dependency
	github.com/gabriel-vasile/mimetype v1.4.3 // Indirect dependency
	github.com/gin-contrib/sse v0.1.0 // Indirect dependency
	github.com/go-playground/locales v0.14.1 // Indirect dependency
	github.com/go-playground/universal-translator v0.18.1 // Indirect dependency
	github.com/go-playground/validator/v10 v10.20.0 // Indirect dependency
	github.com/goccy/go-json v0.10.2 // Indirect dependency
	github.com/jackc/pgpassfile v1.0.0 // Indirect dependency
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // Indirect dependency
	github.com/jackc/pgx/v5 v5.4.3 // Indirect dependency
	github.com/jinzhu/inflection v1.0.0 // Indirect dependency
	github.com/jinzhu/now v1.1.5 // Indirect dependency
	github.com/json-iterator/go v1.1.12 // Indirect dependency
	github.com/klauspost/cpuid/v2 v2.2.7 // Indirect dependency
	github.com/kr/text v0.2.0 // Indirect dependency
	github.com/leodido/go-urn v1.4.0 // Indirect dependency
	github.com/mattn/go-isatty v0.0.20 // Indirect dependency
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // Indirect dependency
	github.com/modern-go/reflect2 v1.0.2 // Indirect dependency
	github.com/pelletier/go-toml/v2 v2.2.2 // Indirect dependency
	github.com/twitchyliquid64/golang-asm v0.15.1 // Indirect dependency
	github.com/ugorji/go/codec v1.2.12 // Indirect dependency
	golang.org/x/arch v0.8.0 // Indirect dependency
	golang.org/x/crypto v0.23.0 // Indirect dependency
	golang.org/x/net v0.25.0 // Indirect dependency
	golang.org/x/sys v0.20.0 // Indirect dependency
	golang.org/x/text v0.15.0 // Indirect dependency
	google.golang.org/protobuf v1.34.1 // Indirect dependency
	gopkg.in/yaml.v3 v3.0.1 // Indirect dependency
)
