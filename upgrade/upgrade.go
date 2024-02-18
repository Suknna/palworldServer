package upgrade

import (
	"fmt"
	"os/exec"
)

func Upgrade() error {
	out, err := exec.Command("steamcmd", "+login anonymous +app_update 2394010 validate +quit").Output()
	if err != nil {
		return err
	}
	fmt.Println(string(out))
	return nil
}
