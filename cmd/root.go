package cmd

import (
	"common-inv/cmd/common"
	"common-inv/cmd/migrate"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/spf13/cobra"
	"os"
)

var (
	version     = "0.1.0"
	programName = "common-inventory"

	rootCmd = &cobra.Command{
		Use:     programName,
		Version: version,
		Short:   "A simple common inventory system",
	}
)

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
}

func init() {
	serveCMD := common.NewCommand()
	migrateCMD := migrate.NewCommand()
	rootCmd.AddCommand(serveCMD, migrateCMD)
}
