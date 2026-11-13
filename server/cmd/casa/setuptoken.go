package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"time"

	"buf.build/go/protovalidate"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/protobuf/reflect/protoreflect"

	casav1 "github.com/abiyanzimran/casa/server/internal/gen/casa/v1"
	"github.com/abiyanzimran/casa/server/internal/identity"
)

// appDatabaseURL names the LoadCredential= file holding the connection
// string of role app_rw.
const appDatabaseURL = "app-database-url"

const setupTokenUsage = "usage: casa setup-token -email address\n" //nolint:gosec // G101: a usage line, no credential

// setupToken issues one setup token for one named person, bound to their
// email (SEC-158 (b); 09.4 §2). The operator runs it over SSH and records the
// issuance in the operations log (SEC-33, SEC-37). It prints the token once;
// the database keeps only its hash.
func setupToken(ctx context.Context, args []string, stdout, stderr io.Writer) int {
	flags := flag.NewFlagSet("setup-token", flag.ContinueOnError)
	flags.SetOutput(io.Discard) // setupTokenUsage is printed instead
	email := flags.String("email", "", "")
	if err := flags.Parse(args); err != nil || flags.NArg() != 0 || *email == "" {
		_, _ = fmt.Fprint(stderr, setupTokenUsage)
		return 2
	}
	if !validEmail(*email) {
		_, _ = fmt.Fprint(stderr, "casa setup-token: -email is not an email address\n")
		return 2
	}
	pool, err := openAppPool(ctx)
	if err != nil {
		_, _ = fmt.Fprintf(stderr, "casa setup-token: %v\n", err)
		return 1
	}
	defer pool.Close()
	token, expiresAt, err := identity.IssueSetupToken(ctx, pool, *email, time.Now())
	if err != nil {
		// Errors from pgx name the host and user at most, never the password.
		_, _ = fmt.Fprintf(stderr, "casa setup-token: %v\n", err)
		return 1
	}
	_, _ = fmt.Fprintf(stdout, "casa setup-token: %s\nexpires: %s\nrecord this issuance in the operations log (SEC-37)\n",
		token, expiresAt.UTC().Format(time.RFC3339))
	return 0
}

// validEmail applies the contract's own rule for StartSignUpRequest.email,
// so the operator cannot issue a token the API would refuse (ES §2).
func validEmail(email string) bool {
	request := &casav1.StartSignUpRequest{Email: email}
	onlyEmail := protovalidate.FilterFunc(func(_ protoreflect.Message, d protoreflect.Descriptor) bool {
		return d.FullName() == request.ProtoReflect().Descriptor().Fields().ByName("email").FullName()
	})
	return protovalidate.Validate(request, protovalidate.WithFilter(onlyEmail)) == nil
}

// openAppPool connects as app_rw with the connection string from its
// credential (09.4 §1; SEC-42).
func openAppPool(ctx context.Context) (*pgxpool.Pool, error) {
	s, err := connString(appDatabaseURL, true)
	if err != nil {
		return nil, err
	}
	pool, err := pgxpool.New(ctx, s)
	if err != nil {
		return nil, fmt.Errorf("open the database: %w", err)
	}
	return pool, nil
}
