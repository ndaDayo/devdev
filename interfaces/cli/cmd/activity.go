/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/ndaDayo/devdev/di"
	"github.com/ndaDayo/devdev/interfaces/cli/tui"
	usecase "github.com/ndaDayo/devdev/usecase/activity"
	"github.com/spf13/cobra"
)

var activityCmd = &cobra.Command{
	Use:   "activity",
	Short: "get activity",
	Run: func(cmd *cobra.Command, args []string) {
		g := &tui.Github{}
		p := tea.NewProgram(tui.New(g), tea.WithoutCatchPanics())

		model, err := p.Run()
		tuiModel, ok := model.(tui.Model)
		if !ok {
			fmt.Println("Error: Model returned from TUI program is not of type tui.Model")
			return
		}

		githubUsername := tuiModel.Github.Username
		githubRepo := tuiModel.Github.Repo
		var opts []func(*usecase.Input)
		if githubUsername != "" && githubRepo != "" {
			opts = append(opts, usecase.WithGithub(&usecase.CodeInput{
				Owner: githubUsername,
				Repo:  githubRepo,
			}))
		}

		u := di.InitializeActivityUseCase()
		_, err = u.Run(opts...)
		if err != nil {
			fmt.Printf("Error in activity cmd: %v\n", err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(activityCmd)
}
