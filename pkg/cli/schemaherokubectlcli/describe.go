package schemaherokubectlcli

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func DescribeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:           "describe",
		Short:         "",
		Long:          `...`,
		Args:          cobra.MinimumNArgs(1),
		SilenceErrors: true,
		PreRun: func(cmd *cobra.Command, args []string) {
			err := viper.BindPFlags(cmd.Flags())
			if err != nil {
				panic(err)
			}
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	cmd.AddCommand(DescribeMigrationCmd())

	return cmd
}
