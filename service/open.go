package service

import (
    "os/exec"
)

func OpenURL(url string) error {
    cmd := exec.Command("open", url)
    return cmd.Start()
}