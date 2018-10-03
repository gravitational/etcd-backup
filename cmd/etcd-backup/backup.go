/*
Copyright 2018 Gravitational, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package etcdexport

import (
	"context"
	"time"

	etcdconf "github.com/gravitational/coordinate/config"
	"github.com/gravitational/etcd-backup/lib/etcd"
	"github.com/gravitational/trace"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var backupCmd = &cobra.Command{
	Use:   "backup [file]",
	Short: "backup etcd datastore",
	Long:  ``,
	Args:  cobra.ExactArgs(1),
	RunE:  backup,
}

var (
	backupTimeout time.Duration
	backupPrefix  []string
)

func init() {
	backupCmd.Flags().DurationVarP(&backupTimeout, "timeout", "", 2*time.Minute, "Cancel the backup if it takes too long")
	backupCmd.Flags().StringSliceVarP(&backupPrefix, "prefix", "", []string{"/"}, "The Etcd path to backup")
	rootCmd.AddCommand(backupCmd)
}

func backup(cmd *cobra.Command, args []string) error {
	ctx, cancel := context.WithTimeout(context.Background(), backupTimeout)
	defer cancel()

	err := etcd.Backup(ctx, etcd.BackupConfig{
		EtcdConfig: etcdconf.Config{
			Endpoints: endpoints,
			CAFile:    caFile,
			CertFile:  certFile,
			KeyFile:   keyFile,
		},
		Prefix: backupPrefix,
		File:   args[0],
		Log:    log.New(),
	})
	if err != nil {
		return trace.Wrap(err)
	}
	return nil
}
