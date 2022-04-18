package cli

import (
	"bufio"
	"io"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var (
	file              string
	sep               string
	index             uint
	uncategorizedName string
)

func Run(version string, args ...string) error {
	cmd := &cobra.Command{
		Use:     "ctxt",
		Short:   "Categorize text",
		Version: version,
		RunE: func(cmd *cobra.Command, args []string) error {
			st, err := os.Stdin.Stat()
			if err != nil {
				return err
			}
			if st.Mode()&os.ModeNamedPipe != 0 {
				return run(cmd, os.Stdin)
			}

			f, err := os.Open(file)
			if err != nil {
				return err
			}
			defer f.Close()

			return run(cmd, f)
		},
	}

	cmd.PersistentFlags().StringVarP(&file, "file", "f", "", "set file path")
	cmd.PersistentFlags().StringVarP(&sep, "sep", "s", "", "set line separator")
	cmd.PersistentFlags().UintVarP(&index, "index", "i", 0, "set which element of the separation are to be used")
	cmd.PersistentFlags().StringVar(&uncategorizedName, "uncategorized-name", "others", "set uncategorized name")

	return cmd.Execute()
}

func run(cmd *cobra.Command, r io.Reader) error {
	sc := bufio.NewScanner(r)

	result := make(map[string][]string) // key: category name, value: texts.

	index := int(index)
	for sc.Scan() {
		text := sc.Text()

		sp := strings.Split(text, sep)
		if index > len(sp) {
			result[uncategorizedName] = append(result[uncategorizedName], text)
			continue
		}

		name := sp[index]
		if len(name) == 0 {
			result[uncategorizedName] = append(result[uncategorizedName], text)
		} else {
			result[name] = append(result[name], text)
		}
	}

	for n, texts := range result {
		if n == uncategorizedName {
			continue
		}
		if len(texts) <= 1 {
			delete(result, n)
			result[uncategorizedName] = append(result[uncategorizedName], texts...)
			continue
		}
		output(cmd, n, texts)
	}
	output(cmd, uncategorizedName, result[uncategorizedName])
	return nil
}

func output(cmd *cobra.Command, name string, texts []string) {
	if len(texts) == 0 {
		return
	}
	cmd.Printf("%s\n", name)
	for _, t := range texts {
		cmd.Printf("%s\n", t)
	}
}
