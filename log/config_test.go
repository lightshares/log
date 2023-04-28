package log

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	config, err := initConfig()
	fmt.Println(config)
	fmt.Println(err)
}
