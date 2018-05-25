package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_getData(t *testing.T) {
	type args struct {
		w      *httptest.ResponseRecorder
		req    *http.Request
		status int
	}
	rec := httptest.NewRecorder()
	rec1 := httptest.NewRecorder()
	data1 := strings.NewReader(`[{"data":"hello", "time":1000}]`)
	data2 := strings.NewReader(`[{"data":"hello", "time":"1000"}]`)
	req1 := httptest.NewRequest("POST", "/", data1)
	req1.Header.Add("Content-Type", "application/json")
	req2 := httptest.NewRequest("POST", "/", data2)
	req2.Header.Add("Content-Type", "application/json")
	tests := []struct {
		name string
		args args
	}{
		{name: "sample", args: args{rec, req1, 200}},
		{name: "invalid", args: args{rec1, req2, 400}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			getData(tt.args.w, tt.args.req)
			res := tt.args.w.Result()
			res.Body.Close()
			defer tt.args.w.Flush()
			if res.StatusCode != tt.args.status {
				t.Errorf("expected %v; got %v", tt.args.status, res.Status)
			}

			res.Body.Close()
			tt.args.w.Flush()
		})
	}
}
