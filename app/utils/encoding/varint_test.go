package encoding_test

import (
	"github/com/codecrafters-io/sqlite-starter-go/app/utils/encoding"
	"os"
	"strconv"
	"testing"
)

func TestVarInt(t *testing.T) {
	tests := []struct {
		value uint
		bytes []byte
	}{
		{0, []byte{0}},
		{1, []byte{1}},
		{127, []byte{127}},

		{128, []byte{0b1_0000001, 0}},            // 0000001 0000000
		{129, []byte{0b1_0000001, 1}},            // 0000001 0000001
		{200, []byte{0b1_0000001, 0b0_1001000}},  // 0000001 1001000
		{16383, []byte{0b1_1111111, 0b01111111}}, // 0111111 1111111

		{16384, []byte{0b1_0000001, 0b10000000, 0}}, // 1000000 0000001 0000000
		{16385, []byte{0b1_0000001, 0b10000000, 1}}, // 1000000 0000001 0000001
	}

	test_dir := os.TempDir()

	for ind, test := range tests {
		t.Logf("Testing OS File based Implementation [Test %d with value %d]", ind, test.value)

		// Create a file with the test bytes
		file, err := os.Create(test_dir + "/varint_test_" + strconv.Itoa(ind))
		if err != nil {
			t.Errorf("Failed to create file: %s", err)
		}
		_, err = file.Write(test.bytes)
		if err != nil {
			t.Errorf("Failed to write to file: %s", err)
		}

		// Read the file and parse the varint
		value, nextIndex := encoding.ReadVarInt(file, 0)
		if value != test.value {
			t.Errorf("Expected %d, but got %d", test.value, value)
		}
		if int(nextIndex) != len(test.bytes) {
			t.Errorf("Expected next index to be %d, but got %d", len(test.bytes), nextIndex)
		}

		// Close the file
		err = file.Close()
		if err != nil {
			t.Errorf("Failed to close file: %s", err)
		}

		t.Logf("Testing Byte Array based Implementation [Test %d with value %d]", ind, test.value)

		res, index := encoding.ReadVarIntFromBytes(test.bytes)
		if res != test.value {
			t.Errorf("Expected %d, but got %d", test.value, res)
		}
		if int(index) != len(test.bytes) {
			t.Errorf("Expected next index to be %d, but got %d", len(test.bytes), index)
		}
	}
}
