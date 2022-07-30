module github.com/glugox/uno

go 1.18

//replace github.com/glugox/uno => ./pkg/

require (
	github.com/go-sql-driver/mysql v1.6.0
	github.com/google/uuid v1.3.0
	github.com/joho/godotenv v1.4.0
	github.com/stretchr/testify v1.8.0
)

require (
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/mattn/go-sqlite3 v1.14.14
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/spf13/cobra v1.5.0
	golang.org/x/exp v0.0.0-20220713135740-79cabaa25d75
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
