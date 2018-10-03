## etcd-backup
A simple utility for backup / restore of etcd data, that works between versions, and can migrate v2 data to v3.

```
Backup / Restore etcd data

Usage:
  etcd-backup [command]

Available Commands:
  backup      backup etcd datastore
  help        Help about any command
  restore     restore etcd datastore

Flags:
      --etcd-cafile string     SSL Certificate Authority file used to secure etcd communication. (default "/var/state/root.cert")
      --etcd-certfile string   SSL certification file used to secure etcd communication. (default "/var/state/etcd.cert")
      --etcd-keyfile string    SSL key file used to secure etcd communication. (default "/var/state/etcd.key")
      --etcd-servers strings   List of etcd servers to connect with (scheme://ip:port), comma separated. (default [https://127.0.0.1:2379])
  -h, --help                   help for etcd-backup

Use "etcd-backup [command] --help" for more information about a command.
```