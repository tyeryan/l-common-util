package secretclient

import (
	"context"
	"github.com/google/wire"
	logutil "github.com/tyeryan/l-protocol/log"
	"time"
)

var (
	log     = logutil.GetLogger("secret-store")
	WireSet = wire.NewSet(
		ProvideSecretClient,
	)
)

// SecretClient secret client
type SecretClient interface {
	//Write write value to secret store
	Write(ctx context.Context, path string, value interface{}, rsp interface{}) (time.Duration, error)
	//Read read secret from secret store
	Read(ctx context.Context, path string, rsp interface{}) (time.Duration, error)
	//List list secret from secret store
	List(ctx context.Context, path string, rsp interface{}) (time.Duration, error)
	//Delete list secret from secret store
	Delete(ctx context.Context, path string, rsp interface{}) error
}

func ProvideSecretClient(ctx context.Context) (SecretClient, error) {
	return ProvideFileSecretClient("yaml")
}

func ProvideJsonSecretClient(ctx context.Context) (SecretClient, error) {
	return ProvideFileSecretClient("json")
}
