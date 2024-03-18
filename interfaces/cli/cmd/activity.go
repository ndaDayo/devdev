/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"time"

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

		owner := tuiModel.Github.Owner
		repo := tuiModel.Github.Repo
		sinceStr := tuiModel.Github.Period.Since

		const layout = "20060102"
		since, err := time.Parse(layout, sinceStr)
		if err != nil {
			fmt.Println("Since date parsing error:", err)
			return
		}
		untilStr := tuiModel.Github.Period.Until
		until, err := time.Parse(layout, untilStr)
		if err != nil {
			fmt.Println("Until date parsing error:", err)
			return
		}

		var opts []func(*usecase.Input)
		if owner != "" && repo != "" {
			opts = append(opts, usecase.WithGithub(&usecase.CodeInput{
				Owner: owner,
				Repo:  repo,
				Since: since.Format(time.RFC3339),
				Until: until.Format(time.RFC3339),
			}))
		}

		usecase := di.InitializeActivityUseCase()
		result, err := usecase.Run(opts...)
		if err != nil {
			fmt.Printf("Error in activity cmd: %v\n", err)
			return
		}

		m, _ := json.MarshalIndent(result, "", "    ")
		fmt.Println(string(m))
	},
}

func init() {
	rootCmd.AddCommand(activityCmd)
}
