package migrate

import (
	"common-inv/internal/conf"
	"common-inv/internal/data"
	"common-inv/internal/models"
	"flag"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/spf13/cobra"
	"os"
)

var flagconf string

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "migrate",
		Short: "Create or migrate the database tables",
		RunE: func(cmd *cobra.Command, args []string) error {
			logger := log.With(log.NewStdLogger(os.Stdout),
				"ts", log.DefaultTimestamp,
				"caller", log.DefaultCaller,
				"trace.id", tracing.TraceID(),
				"span.id", tracing.SpanID(),
			)
			flag.Parse()
			c := config.New(
				config.WithSource(
					file.NewSource(flagconf),
				),
			)
			defer c.Close()

			if err := c.Load(); err != nil {
				panic(err)
			}

			var bc conf.Bootstrap
			if err := c.Scan(&bc); err != nil {
				panic(err)
			}
			db, _, _ := data.NewData(bc.Data, logger)
			return models.Migrate(db.DbFactory.DB)
		},
	}
	cmd.PersistentFlags().StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
	return cmd
}
