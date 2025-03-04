package confflags_test

import (
	"context"
	"fmt"
	"testing"

	flag "github.com/spf13/pflag"
	"github.com/stretchr/testify/require"
	"github.com/sv-tools/conf"

	confflags "github.com/sv-tools/conf-reader-flags"
)

func TestNew(t *testing.T) {
	fs := flag.NewFlagSet("test", flag.ExitOnError)
	fs.Int("foo", 1, "test foo")
	fs.Int("bar", 2, "test bar")
	require.NoError(t, fs.Parse([]string{"--foo=42"}))

	c := conf.New().WithReaders(confflags.New(map[string]string{"foo": "foo", "bar": "bar"}, "", fs))
	require.NoError(t, c.Load(t.Context()))

	require.Equal(t, 42, c.GetInt("foo"))
	require.Equal(t, 2, c.GetInt("bar"))
}

func ExampleNew() {
	flag.IntP("foo", "f", 1, "testing flag")
	if err := flag.CommandLine.Parse([]string{"--foo=42"}); err != nil {
		panic(err)
	}

	c := conf.New().WithReaders(confflags.New(map[string]string{"foo": "foo"}, "", nil))
	if err := c.Load(context.Background()); err != nil {
		panic(err)
	}

	fmt.Println(c.GetInt("foo"))
	// Output: 42
}
