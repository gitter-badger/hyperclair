package cmd

import (
	"fmt"

	"github.com/wemanity-belgium/hyperclair/docker/image"
	"github.com/spf13/cobra"
	//"strings"
	"errors"
)

var reportCmd = &cobra.Command{
	Use:   "report IMAGE",
	Short: "Generate Docker Image vulnerabilities report",
	Long:  `Generate Docker Image vulnerabilities report as HTML or JSON`,
	RunE: func(cmd *cobra.Command, args []string) error {

		if len(args) != 1 {
			return errors.New("hyperclair: \"report\" requires a minimum of 1 argument")
		}

		image, err := image.Parse(args[0])
		if err != nil {
			return err
		}

		if err := image.Pull(); err != nil {
			return err
		}

		fmt.Println("Report Analysis as HTML")
		if err := image.Report(); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	RootCmd.AddCommand(reportCmd)
}
