package get

import (
	"net/http"
	"reflect"
	"testing"
)

func TestNewClient(t *testing.T) {
	t.Run("when call NewClient must return a new client instance", func(t *testing.T) {
		result := NewClient()
		inter := reflect.TypeOf((*Request)(nil)).Elem()
		if result == nil && !reflect.TypeOf(result).Implements(inter) {
			t.Errorf("fail NewGetClient not instance a new Client")
		}
	})
}

func Test_get_Get(t *testing.T) {
	type fields struct {
		request Request
	}
	type args struct {
		url         string
		pathParams  map[string]string
		queryParams map[string]string
		headers     map[string]string
		cookies     []*http.Cookie
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *http.Response
		wantErr bool
	}{
		{
			name:    "when call Get with empty url then return nil and error",
			fields:  fields{},
			args:    args{"", nil, nil, nil, nil},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "when call Get with not empty url but request fail then return nil and error",
			fields:  fields{request: newRequestMock(nil, "something was wrong")},
			args:    args{"/test", nil, nil, nil, nil},
			want:    nil,
			wantErr: true,
		},
		{
			name: "when call Get with not empty url and not fail then return http response and not fail",
			fields: fields{request: newRequestMock(&http.Response{
				Status: "",
			}, "")},
			args: args{"/test", nil, nil, nil, nil},
			want: &http.Response{
				Status: "",
			},
			wantErr: false,
		},
		{
			name: "when call Get with not empty url and all params and not fail then return http response with status code and not fail",
			fields: fields{request: newRequestMock(&http.Response{
				Status: "",
			}, "")},
			args: args{"/test", map[string]string{"test": "test"}, map[string]string{"test": "test"}, map[string]string{"test": "test"}, []*http.Cookie{&http.Cookie{
				Name:  "Test",
				Value: "Test",
			}}},
			want: &http.Response{
				Status: "query,path,headers,cookies,",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &get{
				request: tt.fields.request,
			}
			got, err := g.Get(tt.args.url, tt.args.pathParams, tt.args.queryParams, tt.args.headers, tt.args.cookies)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}
