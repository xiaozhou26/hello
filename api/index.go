//go:build vercel
// +build vercel

package main

import (
	"net/http"
	"os"
	"os/exec"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("./hello")
	cmd.Stdout = w
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
