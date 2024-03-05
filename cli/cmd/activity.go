/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	activity_uc "github.com/ndaDayo/devdev/usecase/activity"
	"github.com/spf13/cobra"
)

var owner, repo, user string

// activityCmd represents the activity command
var activityCmd = &cobra.Command{
	Use:   "activity",
	Short: "get activity",
	Run: func(cmd *cobra.Command, args []string) {
		var opts []func(*activity_uc.ActivityOptions)
		if owner != "" && repo != "" {
			opts = append(opts, activity_uc.WithGithub(&activity_uc.CodeParams{
				Owner: owner,
				Repo:  repo,
			}))
		}

		acs, err := activity_uc.Get(opts...)
		if err != nil {
			fmt.Printf("Error fetching activity: %v\n", err)
			return
		}

		for _, ac := range acs.CodeActivity.PullRequests {
			fmt.Println("CreatedAt:", ac.CreatedAt)
		}
	},
}

func init() {
	rootCmd.AddCommand(activityCmd)

	activityCmd.Flags().StringVarP(&owner, "repo-owner", "u", "", "GitHub repository owner")
	activityCmd.Flags().StringVarP(&repo, "repo", "r", "", "GitHub repository name")
}
