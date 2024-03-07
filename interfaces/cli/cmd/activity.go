/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/ndaDayo/devdev/di"
	usecase "github.com/ndaDayo/devdev/usecase/activity"
	"github.com/spf13/cobra"
)

var githubUsername, githubRepo string

var activityCmd = &cobra.Command{
	Use:   "activity",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		var opts []func(*usecase.Input)
		if githubUsername != "" && githubRepo != "" {
			opts = append(opts, usecase.WithGithub(&usecase.CodeInput{
				Username: githubUsername,
				Repo:     githubRepo,
			}))
		}

		u := di.InitializeActivityUseCase()
		acs, err := u.Run(opts...)
		if err != nil {
			fmt.Printf("Error fetching activity: %v\n", err)
			return
		}

		fmt.Println("type %T", acs)
	},
}

func init() {
	rootCmd.AddCommand(activityCmd)

	activityCmd.Flags().StringVarP(&githubUsername, "github-username", "u", "", "GitHub username")
	activityCmd.Flags().StringVarP(&githubRepo, "github-repo", "r", "", "GitHub repository name")
}
