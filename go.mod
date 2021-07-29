module github.com/gravitational/etcd-backup

go 1.12

require (
	github.com/cenkalti/backoff v2.2.1+incompatible
	github.com/coreos/go-semver v0.3.0
	github.com/gravitational/coordinate/v4 v4.0.0
	github.com/gravitational/trace v0.0.0-20200326194303-2d27a078e25b
	github.com/sirupsen/logrus v1.5.0
	github.com/spf13/cobra v0.0.3
	go.etcd.io/bbolt v1.3.4 // indirect
	go.etcd.io/etcd v3.3.22+incompatible
)

replace (
	github.com/gravitational/coordinate/v4 => github.com/a-palchikov/coordinate/v4 v4.0.0-20210729114333-cca0da0c9f47
	go.etcd.io/bbolt => go.etcd.io/bbolt v1.3.4
	go.etcd.io/etcd => go.etcd.io/etcd v0.5.0-alpha.5.0.20200401174654-e694b7bb0875
)
