package fourth_pos_test

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	fourth_pos "github.com/omniboost/go-fourth-pos"
)

func TestSendFilesRequest(t *testing.T) {
	file, err := os.Open("test.csv")
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}
	defer file.Close()

	req := client.NewSendFilesRequest()
	req.FormParams().ExtfFile = fourth_pos.FormFile{
		Filename: "test.csv",
		Content:  file,
	}

	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
