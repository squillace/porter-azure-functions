package main

import (
	"github.com/squillace/porter-azure-functions/pkg/functions"
	"github.com/spf13/cobra"
)

func buildBuildCommand(m *functions.Mixin) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "build",
		Short: "Generate Dockerfile lines for the bundle invocation image",
		RunE: func(cmd *cobra.Command, args []string) error {
			return m.Build()
		},
	}
	return cmd
}
