package app

import (
	"fmt"
	"os"

	"github.com/layer5io/meshery/mesheryctl/pkg/utils"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var (
	availableSubcommands []*cobra.Command
	file                 string
)

// AppCmd represents the root command for app commands
var AppCmd = &cobra.Command{
	Use:   "app",
	Short: "Service Mesh Apps Management",
	Long:  `Manage all apps operations; ;list, view, onboard and offboard`,
	Example: `
// Base command
mesheryctl app [subcommand]
	`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			// 	return errors.New(utils.AppError(fmt.Sprintln("mesheryctl app requires at least 1 arg(s), got only 0")))
			cmd.Help()
			os.Exit(0)
		}
		if ok := utils.IsValidSubcommand(availableSubcommands, args[0]); !ok {
			return errors.New(utils.AppError(fmt.Sprintf("'%s' is a invalid command. Use 'mesheryctl app --help' to display usage guide.\n", args[0])))
		}
		return nil
	},
}

func init() {
	AppCmd.PersistentFlags().StringVarP(&utils.TokenFlag, "token", "t", "", "Path to token file default from current context")

	availableSubcommands = []*cobra.Command{onboardCmd, viewCmd, offboardCmd, listCmd}
	AppCmd.AddCommand(availableSubcommands...)
}
