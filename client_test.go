package sendinblue

import "testing"
import "net/http"
import "io/ioutil"
import "fmt"

func TestApiAuthentication(t *testing.T) {
	key := "your key"
	client := NewSIBClient(key)
	req, err := http.NewRequest("GET", apiEndPoint+"account", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header["api-key"] = []string{key}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%s", body)
}
