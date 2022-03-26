package adapter

import (
	"bytes"
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

const (
	LABEL_CONTEXT_TEST          string = "adapter_context_test"
	LABEL_LOAD_TEST             string = "adapter_load_test"
	LABEL_LOAD_SYSTEM_ERROR     string = "adapter_load_system_error"
	MSG_ERROR_CONTEXT_MISSING   string = "Context missing"
	MSG_ERROR_LOAD_SYSTEM_ERROR string = "This resolver doesn't work for this loads %v"
	MSG_ECHO_CONTEXT            string = "#Echo# %s"
	MESSAGE_TEXT                string = "Lorem ipsum dolor sit amet"
	URL_RESOLVER                string = "http://localhost:7000/api/v1/resolve/"
	CONTENT_TYPE                string = "application/text"
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	shutdown()
	os.Exit(code)
}

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

func resolverTest(resolveInput types.ResolveInput, resolveOutput *types.ResolveOutput) {
	sort.Strings(resolveInput.Load)
	if contains(resolveInput.Load, LABEL_LOAD_TEST) {
		contextValue, ok := resolveInput.Context[LABEL_CONTEXT_TEST]
		if !ok {
			resolveOutput.Errors[LABEL_LOAD_TEST] = MSG_ERROR_CONTEXT_MISSING
		} else {
			resolveOutput.Context[LABEL_LOAD_TEST] = fmt.Sprintf(MSG_ECHO_CONTEXT, contextValue)
		}

	} else {
		resolveOutput.Errors[LABEL_LOAD_SYSTEM_ERROR] = fmt.Sprintf(MSG_ERROR_LOAD_SYSTEM_ERROR, resolveInput.Load)
	}
}

func contains(s []string, searchterm string) bool {
	i := sort.SearchStrings(s, searchterm)
	return i < len(s) && s[i] == searchterm
}

func TestAdapterSuccess(t *testing.T) {
	resolveOutput := testRequest(t, LABEL_CONTEXT_TEST, MESSAGE_TEXT, LABEL_LOAD_TEST)
	assert.Equal(t, resolveOutput.Context[LABEL_LOAD_TEST], fmt.Sprintf(MSG_ECHO_CONTEXT, MESSAGE_TEXT))
}

func TestAdapterLabelInvalid(t *testing.T) {
	resolveOutput := testRequest(t, LABEL_CONTEXT_TEST, MESSAGE_TEXT, "label_invalid")
	assert.Equal(t, resolveOutput.Errors[LABEL_LOAD_SYSTEM_ERROR], fmt.Sprintf(MSG_ERROR_LOAD_SYSTEM_ERROR, []string{"label_invalid"}))
}

func TestAdapterContextInvalid(t *testing.T) {
	resolveOutput := testRequest(t, "context_invalid", MESSAGE_TEXT, LABEL_LOAD_TEST)
	assert.Equal(t, resolveOutput.Errors[LABEL_LOAD_TEST], MSG_ERROR_CONTEXT_MISSING)
}

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

	resp, err := client.Post(URL_RESOLVER, CONTENT_TYPE, postBody)

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

func BenchmarkAdapterResolver(b *testing.B) {
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			testRequestBench(LABEL_CONTEXT_TEST, MESSAGE_TEXT, LABEL_LOAD_TEST)
		}
	})
}

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

	resp, err := client.Post(URL_RESOLVER, CONTENT_TYPE, postBody)
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
