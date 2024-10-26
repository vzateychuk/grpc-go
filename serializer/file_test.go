package serializer_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"

	"vez/grpc/pb"
	"vez/grpc/sample"
	"vez/grpc/serializer"
)

func TestJSONFileSerializer(t *testing.T) {
	// run the test in parallel
	t.Parallel()

	jsonFile := "../tmp/laptop.json"

	// Create a new laptop
	laptop1 := sample.NewLaptop()
	err := serializer.WriteToJSONFile(laptop1, jsonFile)
	require.NoError(t, err)
}

func TestBinaryFileSerializer(t *testing.T) {
	// run the test in parallel
	t.Parallel()

	binaryFile := "../tmp/laptop.bin"

	// Create a new laptop
	laptop1 := sample.NewLaptop()
	err := serializer.WriteToBinaryFile(laptop1, binaryFile)
	require.NoError(t, err)

	// Read the laptop from the file
	laptop2 := &pb.Laptop{}
	err = serializer.ReadFromBinaryFile(binaryFile, laptop2)
	require.NoError(t, err)

	// Compare the two laptops
	require.True(t, proto.Equal(laptop1, laptop2))
}
