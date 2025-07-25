package expressionquery

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/modelcontextprotocol/go-sdk/jsonschema"
)

type Foo struct {
	Contents []string
}

func TestDemo(t *testing.T) {
	foo := Foo{}
	ret, err := json.Marshal(foo)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(ret))

	schema, err := jsonschema.For[InstantQueryArguments]()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Schema: %#v\n", schema)
	t.Log(time.Now().Format(time.RFC3339))
	time.Parse(time.RFC3339, "2025-07-28T14:50:38+08:00")
}
