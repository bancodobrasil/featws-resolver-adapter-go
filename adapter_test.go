package adapter

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sort"
	"testing"
	"time"

	"github.com/bancodobrasil/featws-resolver-adapter-go/types"
	"github.com/stretchr/testify/assert"
)

// This is a block of constant declarations in Go. It defines several string constants that are used
// throughout the code, such as labels for context and load, error messages, and a URL for the resolver
// API. These constants are used to ensure consistency and avoid hardcoding values throughout the code.
const (
	labelContextTest        string = "adapter_context_test"
	labelLoadTest           string = "adapter_load_test"
	labelLoadSystemError    string = "adapter_load_system_error"
	msgErrorContextMissing  string = "Context missing"
	msgErrorLoadSystemError string = "This resolver doesn't work for this loads %v"
	msgEchoContext          string = "#Echo# %s"
	messageText             string = "Lorem ipsum dolor sit amet"
	urlResolver             string = "http://localhost:7000/api/v1/resolve/"
	contentType             string = "application/text"
)

// TestMain function sets up and shuts down a test environment for Go tests.
func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	shutdown()
	os.Exit(code)
}

// setup sets up a server in a goroutine and waits for it to start running before continuing.
func setup() {
	go func() {
		Run(resolverTest, Config{
			Port: "7000",
		})
	}()
	// have to wait for the goroutine to start and run the server
	// otherwise the main thread will complete
	time.Sleep(5 * time.Millisecond)
}

func shutdown() {}

// resolverTest checks if a specific load is present in the input and returns an error
// message if it is not, or returns a context value if it is present.
func resolverTest(ctx context.Context, resolveInput types.ResolveInput, resolveOutput *types.ResolveOutput) {
	sort.Strings(resolveInput.Load)
	if contains(resolveInput.Load, labelLoadTest) {
		contextValue, ok := resolveInput.Context[labelContextTest]
		if !ok {
			resolveOutput.Errors[labelLoadTest] = msgErrorContextMissing
		} else {
			resolveOutput.Context[labelLoadTest] = fmt.Sprintf(msgEchoContext, contextValue)
		}

	} else {
		resolveOutput.Errors[labelLoadSystemError] = fmt.Sprintf(msgErrorLoadSystemError, resolveInput.Load)
	}
}

// contains checks if a given string is present in a sorted slice of strings.
func contains(s []string, searchterm string) bool {
	i := sort.SearchStrings(s, searchterm)
	return i < len(s) && s[i] == searchterm
}

// TestAdapterSuccess is a test function in Go that checks if the output of a request matches an expected output.
func TestAdapterSuccess(t *testing.T) {
	resolveOutput := testRequest(t, labelContextTest, messageText, labelLoadTest)
	assert.Equal(t, resolveOutput.Context[labelLoadTest], fmt.Sprintf(msgEchoContext, messageText))
}

// TestAdapterLabelInvalid is a test function in Go that checks for an invalid label and returns an error message.
func TestAdapterLabelInvalid(t *testing.T) {
	resolveOutput := testRequest(t, labelContextTest, messageText, "label_invalid")
	assert.Equal(t, resolveOutput.Errors[labelLoadSystemError], fmt.Sprintf(msgErrorLoadSystemError, []string{"label_invalid"}))
}

// TestAdapterContextInvalid is a test function in Go that checks for invalid context and returns an error message.
func TestAdapterContextInvalid(t *testing.T) {
	resolveOutput := testRequest(t, "context_invalid", messageText, labelLoadTest)
	assert.Equal(t, resolveOutput.Errors[labelLoadTest], msgErrorContextMissing)
}

// testRequest sends a POST request to a URL with JSON data and returns the response as a
// ResolveOutput object.
func testRequest(t *testing.T, context string, data string, load string) *types.ResolveOutput {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	client := &http.Client{Transport: tr}

	body, _ := json.Marshal(types.ResolveInput{
		Context: map[string]interface{}{
			context: data,
		},
		Load: []string{load},
	})

	postBody := bytes.NewBuffer((body))

	defer client.CloseIdleConnections()

	resp, err := client.Post(urlResolver, contentType, postBody)

	assert.NoError(t, err)
	assert.Equal(t, "200 OK", resp.Status)

	defer resp.Body.Close()

	resBody, err := ioutil.ReadAll(resp.Body)
	assert.NoError(t, err)

	resolveOutput := types.ResolveOutput{
		Context: make(map[string]interface{}),
		Errors:  make(map[string]interface{}),
	}
	json.Unmarshal(resBody, &resolveOutput)

	return &resolveOutput

}

// go test -bench . -run="none" -v -count=5

// BenchmarkAdapterResolver is a benchmark function in Go that tests the performance of a function called
// "testRequestBench" using parallel testing.
func BenchmarkAdapterResolver(b *testing.B) {
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			testRequestBench(labelContextTest, messageText, labelLoadTest)
		}
	})
}

// testRequestBench sends a POST request to a specified URL with JSON data and reads the response.
func testRequestBench(context string, data string, load string) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	client := &http.Client{Transport: tr}

	body, _ := json.Marshal(types.ResolveInput{
		Context: map[string]interface{}{
			context: data,
		},
		Load: []string{load},
	})

	postBody := bytes.NewBuffer((body))

	defer client.CloseIdleConnections()

	resp, err := client.Post(urlResolver, contentType, postBody)
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode != 200 {
		log.Fatal("Http response status not equals 200")
	}

	defer resp.Body.Close()

	resBody, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	resolveOutput := types.ResolveOutput{
		Context: make(map[string]interface{}),
		Errors:  make(map[string]interface{}),
	}

	err = json.Unmarshal(resBody, &resolveOutput)

	if err != nil {
		log.Fatal(err)
	}

}
