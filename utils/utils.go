package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	"os/exec"
)

func GetPrettyJSON(v interface{}) string {
	j, _ := json.Marshal(v)
	c1 := exec.Command("echo", string(j))
	c2 := exec.Command("jq", ".", "-C")

	r, w := io.Pipe()
	c1.Stdout = w
	c2.Stdin = r

	var b2 bytes.Buffer
	c2.Stdout = &b2

	c1.Start()
	c2.Start()
	c1.Wait()
	w.Close()
	c2.Wait()
	io.Copy(os.Stdout, &b2)

	return b2.String()
}
