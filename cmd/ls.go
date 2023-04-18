package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/pipikai/yun/common/leveldb"
	"github.com/pipikai/yun/core/tracker/models"
	"github.com/spf13/cobra"
)

var LsCmd = &cobra.Command{
	Use:   "ls",
	Short: "show all file info ",
	Run: func(cmd *cobra.Command, args []string) {
		ls()
	},
}

func ls() {
	res, err := leveldb.GetAll[models.FileInfo]()
	if err != nil {
		fmt.Println("err :", err)
		return
	}
	for _, v := range res {
		show, _ := json.Marshal(v)
		fmt.Printf("%s : %s \n", v.ID, show)
	}
}

func init() {
	RootCmd.AddCommand(LsCmd)
}
