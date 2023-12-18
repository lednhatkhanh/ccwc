package cmd

import (
	"ccwc/utils"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func NewRootCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "ccwc",
		Short: "word, line, character, and byte count.",
		Long:  "word, line, character, and byte count.",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			var fileName string
			var fileContent []byte

			if len(args) == 1 {
				filename := args[0]
				content, err := utils.ReadFile(filename)
				if err != nil {
					return err
				}

				fileName = filename
				fileContent = content
			} else {
				stdin, err := utils.ReadStdin()
				if err != nil {
					return err
				}

				fileContent = stdin
			}

			output := ""

			bytesFlag, err := cmd.Flags().GetBool("count")
			if err != nil {
				return err
			}

			linesFlag, err := cmd.Flags().GetBool("line")
			if err != nil {
				return err
			}

			wordsFlag, err := cmd.Flags().GetBool("word")
			if err != nil {
				return err
			}

			charsFlag, err := cmd.Flags().GetBool("character")
			if err != nil {
				return err
			}

			switch {
			case bytesFlag:
				output += fmt.Sprintf("%d", utils.NoOfBytes(fileContent))
			case linesFlag:
				output += fmt.Sprintf("%d", utils.NoOfLines(fileContent))
			case wordsFlag:
				output += fmt.Sprintf("%d", utils.NoOfWords(fileContent))
			case charsFlag:
				output += fmt.Sprintf("%d", utils.NoOfChars(fileContent))
			default:
				output += fmt.Sprintf("%d %d %d", utils.NoOfLines(fileContent), utils.NoOfWords(fileContent), utils.NoOfBytes(fileContent))
			}

			if fileName != "" {
				output += fmt.Sprintf(" %s\n", fileName)
			} else {
				output += "\n"
			}

			fmt.Fprintf(cmd.OutOrStdout(), output)
			return nil
		},
	}

	cmd.PersistentFlags().BoolP("count", "c", false, "Outputs the number of bytes in a file.")
	cmd.PersistentFlags().BoolP("line", "l", false, "Outputs the number of lines in a file.")
	cmd.PersistentFlags().BoolP("word", "w", false, "Outputs the number of words in a file.")
	cmd.PersistentFlags().BoolP("character", "m", false, "Outputs the number of characters in a file.")

	return cmd
}

func Execute() {
	err := NewRootCmd().Execute()
	if err != nil {
		os.Exit(1)
	}
}
