package repository

import (
	"net/http"
	"reflect"
	"testing"
)

type RequestFake struct {
}

func (r RequestFake) GetUrl() string {
	return ""
}

func (r RequestFake) GetPathParams() map[string]string {
	return nil
}

func (r RequestFake) GetQueryParams() map[string]string {
	return nil
}

func (r RequestFake) GetHeaders() map[string]string {
	return nil
}

func (r RequestFake) GetCookies() []*http.Cookie {
	return nil
}

func TestNewRepository(t *testing.T) {
	t.Run("when call NewRepository must return a new repository instance", func(t *testing.T) {
		result := NewRepository(newClientMock(nil, ""))
		inter := reflect.TypeOf((*Repository)(nil)).Elem()
		if result == nil && !reflect.TypeOf(result).Implements(inter) {
			t.Errorf("fail NewRepository not instance a new Repository")
		}
	})
}

func TestRepository_GetResponse(t *testing.T) {
	type fields struct {
		client GetClient
	}
	type args struct {
		request Request
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *http.Response
		wantErr bool
	}{
		{
			name:    "when call GetResponse with nil request, then must return error",
			fields:  fields{client: newClientMock(nil, "")},
			args:    args{request: nil},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "when call GetResponse with not nil request, but fail get, then must return error",
			fields:  fields{client: newClientMock(nil, "some error")},
			args:    args{request: RequestFake{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "when call GetResponse with not nil request, but fail get, then must return error",
			fields:  fields{client: newClientMock(&http.Response{}, "")},
			args:    args{request: RequestFake{}},
			want:    &http.Response{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repository{
				client: tt.fields.client,
			}
			got, err := r.GetResponse(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetResponse() got = %v, want %v", got, tt.want)
			}
		})
	}
}
