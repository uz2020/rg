/*
Copyright © 2022 Zhj Rong <rongzhj2020@163.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// roffnewCmd represents the roffnew command
var roffnewCmd = &cobra.Command{
	Use:   "new",
	Short: "new post",
	Long:  ``,
	Run:   roffnewRun,
}

const (
	HEAD = `.PH "'' %s ''"
.PF "'' -%%- ''"
`
)

func roffnewRun(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		fmt.Println("post name pls")
		return
	}

	baseDir := "content"
	for _, name := range args {
		dirName := baseDir + "/" + filepath.Dir(name)
		pathName := baseDir + "/" + name + ".ms"
		err := os.MkdirAll(dirName, 0755)
		if err != nil {
			fmt.Println(err)
			return
		}
		f, err := os.OpenFile(pathName, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		// 初始化内容
		head := fmt.Sprintf(HEAD, name)
		f.Write([]byte(head))
	}
	fmt.Println("roffnew called")
}

func init() {
	rootCmd.AddCommand(roffnewCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// roffnewCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// roffnewCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
