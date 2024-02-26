/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

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

		a, err := activity_uc.Get(opts...)
		if err != nil {
			fmt.Printf("Error fetching activity: %v\n", err)
			return
		}

		fmt.Printf("Total GitHub activity changes: %d\n", a.Github.TotalLen)
	},
}

func init() {
	rootCmd.AddCommand(activityCmd)

	activityCmd.Flags().StringVarP(&githubUsername, "github-username", "u", "", "GitHub username")
	activityCmd.Flags().StringVarP(&githubRepo, "github-repo", "r", "", "GitHub repository name")
}