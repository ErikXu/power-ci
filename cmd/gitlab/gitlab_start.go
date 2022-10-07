package gitlab

import (
	"bufio"
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var gitlabStartCmd = &cobra.Command{
	Use:   "install",
	Short: "Install gitlab",
	Run: func(cmd *cobra.Command, args []string) {
		command := exec.Command("gitlab-ctl", "reconfigure")
		cmdReader, _ := command.StdoutPipe()
		scanner := bufio.NewScanner(cmdReader)
		done := make(chan bool)
		go func() {
			for scanner.Scan() {
				fmt.Println(scanner.Text())
			}
			done <- true
		}()
		command.Start()
		<-done
		err := command.Wait()
		if err != nil {
			fmt.Println(err.Error())
		}
	},
}
