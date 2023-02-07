package git

import (
	"context"
	"fmt"

	"github.com/taubyte/go-auth-http/git/common"
	githubClient "github.com/taubyte/go-auth-http/git/github"
)

func New(ctx context.Context, provider, token string) common.Client {
	switch provider {
	case "github":
		return githubClient.New(ctx, token)
	default:
		panic(fmt.Sprintf("provider `%s` is not supported", provider))
	}
}
