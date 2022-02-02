// +heroku goVersion go1.17
// +heroku install ./cmd/poe-status-server/

module github.com/Gonzih/poe-status.com

require (
	github.com/apache/thrift v0.12.0 // indirect
	github.com/cockroachdb/cockroach-go v0.0.0-20181001143604-e0a95dfd547c // indirect
	github.com/gogo/protobuf v1.3.2
	github.com/golang-migrate/migrate/v4 v4.15.1
	github.com/golang/protobuf v1.5.2
	github.com/google/go-github v17.0.0+incompatible // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/jackc/fake v0.0.0-20150926172116-812a484cc733 // indirect
	github.com/jackc/pgx v3.2.0+incompatible // indirect
	github.com/jmoiron/sqlx v1.3.4
	github.com/lib/pq v1.10.4
	github.com/neo4j-drivers/gobolt v1.7.4 // indirect
	github.com/openzipkin/zipkin-go v0.1.6 // indirect
	github.com/shurcooL/httpfs v0.0.0-20190707220628-8d4bc4ba7749 // indirect
	github.com/shurcooL/vfsgen v0.0.0-20200824052919-0d455de96546
	github.com/spf13/cobra v1.3.0
	github.com/stretchr/testify v1.7.0
	github.com/tomsteele/go-nmap v0.0.0-20191202052157-3507e0b03523
	github.com/twitchtv/twirp v8.1.1+incompatible
	github.com/xdg/scram v0.0.0-20180814205039-7eeb5667e42c // indirect
	github.com/xdg/stringprep v1.0.0 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	google.golang.org/protobuf v1.27.1
	gopkg.in/yaml.v2 v2.4.0
)

go 1.16
