package util

import (
	"log"
	"os/exec"
)

func Cmd(repo, commitid string) string {
	// 创建 Command 对象
	cmd := exec.Command("kubectl", "-n", "linkflow", "set", "image", "deploy/"+repo, repo+"="+"reg.leadswarp.com/"+repo+"/dev:"+commitid)

	// 执行命令
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error executing command: %v\n", err)
		log.Printf("Command output: %s\n", output)
		return string(output)
	}

	//fmt.Println("Command executed successfully.")
	//fmt.Println(string(output))

	return "Command executed successfully.\n" + string(output)
}
func CmdUpgrade() string {
	// 创建 Command 对象
	cmd := exec.Command("bash", "/opt/upgrade.sh")

	// 执行命令
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error executing command: %v\n", err)
		log.Printf("Command output: %s\n", output)
		return string(output)
	}

	//fmt.Println("Command executed successfully.")
	//fmt.Println(string(output))

	return "Command executed successfully.\n" + string(output)
}
