package service

import (
    "os/exec"
    "strings"
)
// 获取git分支名称
func GetGitBranch() (string, error) {
    cmd := exec.Command("git", "status")
    output, err := cmd.Output()
    if err != nil {
        return "", err
    }

    for _, line := range strings.Split(string(output), "\n") {
        if strings.Contains(line, "On branch") {
            parts := strings.Split(line, " ")
            if len(parts) >= 3 {
                return parts[2], nil
            }
            break
        }
    }

    return "", nil
}