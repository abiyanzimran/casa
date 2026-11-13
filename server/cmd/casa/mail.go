package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/abiyanzimran/casa/server/internal/identity"
)

// mailKinds names each message kind in the file name and body.
var mailKinds = map[identity.MessageKind]string{
	identity.MessageVerification: "verification",
	identity.MessageRecovery:     "recovery",
}

// writeMail returns a mail sender that writes each message to its own file
// in dir, readable by the owner only.
//
// ponytail: a fake sender for the developer machine and the temporary test
// server, which holds test data only (14 §7.8 item 2); the SES sender
// replaces it at step 7 (ADR-026).
func writeMail(dir string) func(context.Context, identity.Message) error {
	return func(_ context.Context, m identity.Message) error {
		kind, ok := mailKinds[m.Kind]
		if !ok {
			return fmt.Errorf("unknown message kind %d", m.Kind)
		}
		name := time.Now().UTC().Format("20060102T150405.000000000Z") + "-" + kind + ".txt"
		body := fmt.Sprintf("To: %s\nKind: %s\nLink: %s\n", m.To, kind, m.Link)
		return os.WriteFile(filepath.Join(dir, name), []byte(body), 0o600)
	}
}
