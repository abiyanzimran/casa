package main

import (
	"context"
	"errors"
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
		// O_EXCL: a mail never overwrites another (code-reviewer M3 on #32).
		f, err := os.OpenFile(filepath.Join(dir, name), os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0o600) //nolint:gosec // G304: the operator's -mail-dir and a generated name
		if err != nil {
			return err
		}
		_, err = f.WriteString(body)
		return errors.Join(err, f.Close())
	}
}
