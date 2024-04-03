package main

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
)

func executeQuery(query string) {
	// 这里编写你的查询逻辑
	fmt.Printf("Query executed with input: %s\n", query)
}

func init() {
	// 在这里使用 Cobra 和 Viper 配置你的命令
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Hugo",
	Long:  `All software has versions. This is Hugo's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hugo Static Site Generator v0.9 -- HEAD")
	},
}

func main() {
	var rootCmd = &cobra.Command{
		Use:   "q",
		Short: "A brief description of your application",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				fmt.Printf("No enough args")
				return errors.New("no enough args")
			}
			return nil
		},
		Long: `A longer description that spans in lines and is ideal for adding repositories, bug trackers, and other useful information that users may need.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(args)
			executeQuery(args[0]) // 这里调用你的查询函数
		},
	}
	rootCmd.AddCommand(versionCmd)

	// 程序入口
	rootCmd.Execute()
}
