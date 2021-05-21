package handlers

import (
	"net/http"
	"os/exec"
	"strings"

	"github.com/seggewiss/e2eCleanUp/internal/output"
)

type BuildExclusionTree struct{}

func (c *BuildExclusionTree) Supports(url string) bool {
	return strings.Index(url, "/build-exclusions?templateId=") == 0
}

func (c *BuildExclusionTree) Handle(dir string, res http.ResponseWriter, req *http.Request) {
	cmd := exec.Command("bin/console", "cupro-template:generate-tree", req.URL.Query().Get("templateId"))
	cmd.Dir = dir

	output.HandleCommandOutput(cmd.Output())

	CreateSuccessResponse(res)
}
