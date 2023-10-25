package secretclient

import (
	"context"
	"errors"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"github.com/tyeryan/l-common-util/config"
	"time"
)

type FileSecretClient struct {
	ConfigType string
}

func (c *FileSecretClient) Write(ctx context.Context, path string, value interface{}, rsp interface{}) (time.Duration, error) {
	return 0, errors.New("unsupported operation")
}
func (c *FileSecretClient) List(ctx context.Context, path string, rsp interface{}) (time.Duration, error) {
	return 0, errors.New("unsupported operation")
}
func (c *FileSecretClient) Delete(ctx context.Context, path string, rsp interface{}) error {
	return errors.New("unsupported operation")
}

// Duration is always 1 hour
func (c *FileSecretClient) Read(ctx context.Context, path string, rsp interface{}) (time.Duration, error) {
	log.Debugw(ctx, "reading", "path", path)

	v := viper.New()
	v.SetConfigFile(path)
	v.SetConfigType(c.ConfigType)

	if err := v.ReadInConfig(); err != nil {
		return 0, err
	}

	log.Debugw(ctx, "secret file read", "keys", v.AllKeys())

	if err := v.Unmarshal(rsp, func(option *mapstructure.DecoderConfig) {
		option.TagName = config.DefStructTagName
		option.WeaklyTypedInput = true
	}); err != nil {
		return 0, err
	}

	return 1 * time.Hour, nil
}

func ProvideFileSecretClient(configType string) (SecretClient, error) {
	return &FileSecretClient{
		ConfigType: configType,
	}, nil
}
