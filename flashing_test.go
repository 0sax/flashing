package flashing

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestFlashMessage_Set(t *testing.T) {
	type fields struct {
		Message string
		Type    string
	}
	type args struct {
		w          http.ResponseWriter
		cookieName string
	}
	//var xx http.ResponseWriter
	xx := httptest.NewRecorder()
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "test 1",
			fields: fields{
				Message: "Cookie message",
				Type:    "good",
			},
			args: args{
				w:          xx,
				cookieName: "nikolai",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fm := &FlashMessage{
				Message: tt.fields.Message,
				Type:    tt.fields.Type,
			}
			if err := fm.Set(tt.args.w, tt.args.cookieName); (err != nil) != tt.wantErr {
				t.Errorf("Set() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetFlash(t *testing.T) {
	//var rr http.Request
	flasho := &FlashMessage{
		Message: "totorimi once",
		Type:    "alrighty",
	}
	type args struct {
		w          http.ResponseWriter
		r          *http.Request
		cookieName string
		flash      *FlashMessage
	}
	tests := []struct {
		name    string
		args    args
		want    *FlashMessage
		wantErr bool
	}{
		{
			name: "test 1",
			args: args{
				w:          httptest.NewRecorder(),
				r:          httptest.NewRequest(http.MethodGet, "", nil),
				cookieName: "nikolai",
				flash:      flasho,
			},
			want:    flasho,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			tt.args.flash.Set(tt.args.w, tt.args.cookieName)

			got, err := GetFlash(tt.args.w, tt.args.r, tt.args.cookieName)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFlash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetFlash() got = %v, want %v", got, tt.want)
			}
		})
	}
}
