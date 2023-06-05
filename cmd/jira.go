/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/brizer/mycli/service"
)

// jiraCmd represents the jira command
var jiraCmd = &cobra.Command{
	Use:   "jira",
	Short: "打开对应jira链接",
	Long: `基于你所在工程当前的分支名，打开对应的jira链接`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("jira called")
		branch, err := service.GetGitBranch()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		jiraID := service.GetJiraID(branch)
		fmt.Println("Current branch:", branch)
		fmt.Println("Current jiraID:", jiraID)
		if jiraID == "" {
			fmt.Println("当前分支名称不对，请符合feature-****-****格式")
			return
		} 
		url := "http://jira.virtueit.net/browse/" + jiraID
		service.OpenURL(url)
	},
}

func init() {
	rootCmd.AddCommand(jiraCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// jiraCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// jiraCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
