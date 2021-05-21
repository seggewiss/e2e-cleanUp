package handlers

import (
	"net/http"
	"os/exec"

	"github.com/seggewiss/e2eCleanUp/internal/output"
)

type CleanUp struct{}

func (c *CleanUp) Supports(url string) bool {
	return url == "/cleanup"
}

func (c *CleanUp) Handle(dir string, res http.ResponseWriter, req *http.Request) {
	cmd := exec.Command("psh", "e2e:cleanup")
	cmd.Dir = dir

	output.HandleCommandOutput(cmd.Output())

	CreateSuccessResponse(res)
}
