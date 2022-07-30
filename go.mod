module github.com/glugox/uno

go 1.18

//replace github.com/glugox/uno => ./pkg/

require (
	github.com/google/uuid v1.3.0
	github.com/joho/godotenv v1.4.0
	gorm.io/gorm v1.23.8
)

require (
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.4 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
)

require (
	github.com/mattn/go-sqlite3 v1.14.14
	github.com/spf13/cobra v1.5.0
	golang.org/x/exp v0.0.0-20220713135740-79cabaa25d75
)
