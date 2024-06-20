package ssh

import (
	"log"
	"os"
	"strconv"
	"time"
)

var (
	Debug, _ = strconv.Atoi(os.Getenv("SSH_DEBUG"))
	DebugLog = func() func(sent int, typ string, v interface{}, err error) {
		if Debug <= 0 {
			return func(sent int, typ string, v interface{}, err error) {}
		}

		sshLogFile, err := os.Create("ssh.log")
		if err != nil {
			log.Fatalf("failed to open ssh.log: %v", err)
		}

		sshLog := log.New(sshLogFile, "", 0)

		return func(sent int, typ string, v interface{}, err error) {
			data := map[string]interface{}{
				"timestamp": time.Now().Format(time.RFC3339),
				"type":      typ,
				"sent":      sent,
				"value":     v,
			}
			if err != nil {
				data["error"] = err.Error()
			}

			sshLog.Printf("%s", JSON(data))
			_ = sshLogFile.Sync()
		}
	}()
)
