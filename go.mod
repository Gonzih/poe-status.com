// +heroku install ./cmd/poe-status-server/

module github.com/Gonzih/poe-status.com

require (
	github.com/gogo/protobuf v1.3.1
	github.com/golang-migrate/migrate/v4 v4.9.1
	github.com/golang/protobuf v1.4.1
	github.com/jmoiron/sqlx v1.2.0
	github.com/lib/pq v1.5.1
	github.com/spf13/cobra v0.0.7
	github.com/stretchr/testify v1.5.1
	github.com/tomsteele/go-nmap v0.0.0-20181105160706-3b9bafddefee
	github.com/twitchtv/twirp v5.10.1+incompatible
	golang.org/x/sync v0.0.0-20190911185100-cd5d95a43a6e
	google.golang.org/protobuf v1.22.0
	gopkg.in/yaml.v2 v2.4.0
)

go 1.16
