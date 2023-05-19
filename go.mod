// +heroku goVersion go1.17
// +heroku install ./cmd/poe-status-server/

module github.com/Gonzih/poe-status.com

require (
	github.com/gogo/protobuf v1.3.2
	github.com/golang-migrate/migrate/v4 v4.15.2
	github.com/golang/protobuf v1.5.2
	github.com/jmoiron/sqlx v1.3.5
	github.com/lib/pq v1.10.8
	github.com/shurcooL/httpfs v0.0.0-20190707220628-8d4bc4ba7749 // indirect
	github.com/shurcooL/vfsgen v0.0.0-20200824052919-0d455de96546
	github.com/spf13/cobra v1.7.0
	github.com/stretchr/testify v1.8.3
	github.com/tomsteele/go-nmap v0.0.0-20191202052157-3507e0b03523
	github.com/twitchtv/twirp v8.1.3+incompatible
	go.uber.org/atomic v1.9.0 // indirect
	golang.org/x/sync v0.2.0
	google.golang.org/protobuf v1.30.0
	gopkg.in/yaml.v2 v2.4.0
)

go 1.16
