/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/ndaDayo/devdev/di"
	activity "github.com/ndaDayo/devdev/interfaces/cli/tui"
	usecase "github.com/ndaDayo/devdev/usecase/activity"
	"github.com/spf13/cobra"
)

var githubUsername, githubRepo string

var activityCmd = &cobra.Command{
	Use:   "activity",
	Short: "get activity",
	Run: func(cmd *cobra.Command, args []string) {
		if _, err := tea.NewProgram(activity.InitialModel()).Run(); err != nil {
			fmt.Printf("could not start program: %s\n", err)
			os.Exit(1)
		}

		var opts []func(*usecase.Input)
		if githubUsername != "" && githubRepo != "" {
			opts = append(opts, usecase.WithGithub(&usecase.CodeInput{
				Owner: githubUsername,
				Repo:  githubRepo,
			}))
		}

		u := di.InitializeActivityUseCase()
		acs, err := u.Run(opts...)
		if err != nil {
			fmt.Printf("Error activity cmd: %v\n", err)
			return
		}
		// TODO delete because for debug
		fmt.Println("type %T", acs)
	},
}

func init() {
	rootCmd.AddCommand(activityCmd)

	activityCmd.Flags().StringVarP(&githubUsername, "github-username", "u", "", "GitHub username")
	activityCmd.Flags().StringVarP(&githubRepo, "github-repo", "r", "", "GitHub repository name")
}
