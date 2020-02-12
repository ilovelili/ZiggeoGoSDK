package ziggeo

import "testing"

import "fmt"

func Test_Get(t *testing.T) {
	zeggio := NewZiggeo("<<application token>>", "<<private key>>", "<<encrypt key>>")
	resp, err := zeggio.Videos().Index(map[string]string{})
	fmt.Println(string(resp))
	if err != nil {
		t.Error(err.Error())
	}
}
