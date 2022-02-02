// +heroku install ./cmd/poe-status-server/

module github.com/Gonzih/poe-status.com

require (
	github.com/gogo/protobuf v1.3.2
	github.com/golang-migrate/migrate/v4 v4.15.1
	github.com/golang/protobuf v1.5.2
	github.com/jmoiron/sqlx v1.3.1
	github.com/lib/pq v1.10.0
	github.com/spf13/cobra v1.0.0
	github.com/stretchr/testify v1.7.0
	github.com/tomsteele/go-nmap v0.0.0-20181105160706-3b9bafddefee
	github.com/twitchtv/twirp v5.10.1+incompatible
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	google.golang.org/protobuf v1.27.1
	gopkg.in/yaml.v2 v2.4.0
)

go 1.16
