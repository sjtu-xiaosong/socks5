package socks5

import (
	"bytes"
	"reflect"
	"testing"
)

func TestNewClientAuthMessage(t *testing.T) {
	t.Run("should generate a message", func(t *testing.T) {
		b := []byte{socks5Version, 2, 0x00, 0x01}
		reader := bytes.NewReader(b)
		message, err := NewClientAuthMessage(reader)
		if err != nil {
			t.Fatalf("unexpected error: %s", err) //测试用例标记为不通过
		}

		if message.Version != socks5Version {
			t.Fatalf("unexpected version: %d", message.Version)
		}

		if message.NMethods != 2 {
			t.Fatalf("unexpected number of methods: %d", message.NMethods)
		}

		if !reflect.DeepEqual(message.Methods, []byte{0x00, 0x01}) {
			t.Fatalf("unexpected methods: %v", message.Methods)
		}

	})
}
