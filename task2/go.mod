module chinese_chess

go 1.21

require (
	chinese_chess/src v0.0.0-00010101000000-000000000000
	github.com/cucumber/godog v0.14.1
)

replace chinese_chess/src => ./src

require (
	github.com/cucumber/gherkin/go/v26 v26.2.0 // indirect
	github.com/cucumber/messages/go/v21 v21.0.1 // indirect
	github.com/gofrs/uuid v4.3.1+incompatible // indirect
	github.com/hashicorp/go-immutable-radix v1.3.1 // indirect
	github.com/hashicorp/go-memdb v1.3.4 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
)
