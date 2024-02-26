/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	activity_uc "github.com/ndaDayo/devdev/usecase/activity"
	"github.com/spf13/cobra"
)

var githubUsername, githubRepo string

// activityCmd represents the activity command
var activityCmd = &cobra.Command{
	Use:   "activity",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		var opts []func(*activity_uc.ActivityOptions)
		if githubUsername != "" && githubRepo != "" {
			opts = append(opts, activity_uc.WithGithub(&activity_uc.GithubParams{
				Username: githubUsername,
				Repo:     githubRepo,
			}))
		}

	},
}

func init() {
	rootCmd.AddCommand(activityCmd)

	activityCmd.Flags().StringVarP(&githubUsername, "github-username", "gu", "", "GitHub username")
	activityCmd.Flags().StringVarP(&githubRepo, "github-repo", "gr", "", "GitHub repository name")
}
