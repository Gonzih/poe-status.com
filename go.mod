// +heroku install ./cmd/poe-status-server/

module github.com/Gonzih/poe-status.com

require (
	github.com/docker/distribution v2.7.1+incompatible // indirect
	github.com/gogo/protobuf v1.3.1
	github.com/golang-migrate/migrate/v4 v4.9.1
	github.com/golang/protobuf v1.3.3
	github.com/jmoiron/sqlx v1.2.0
	github.com/lib/pq v1.3.0
	github.com/shurcooL/httpfs v0.0.0-20190527155220-6a4d4a70508b // indirect
	github.com/shurcooL/vfsgen v0.0.0-20181202132449-6a9ea43bcacd // indirect
	github.com/spf13/cobra v0.0.5
	github.com/stretchr/testify v1.5.1
	github.com/tomsteele/go-nmap v0.0.0-20181105160706-3b9bafddefee
	github.com/twitchtv/twirp v5.10.1+incompatible
	golang.org/x/sync v0.0.0-20190423024810-112230192c58
	gopkg.in/yaml.v2 v2.2.8
)

go 1.13
