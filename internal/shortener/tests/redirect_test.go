package tests

import (
	"context"
	"reflect"
	"testing"

	svc "url-shortener/internal/shortener"
	"url-shortener/internal/shortener/domain"
	"url-shortener/internal/shortener/service"
)

func TestRedirect(t *testing.T) {
	type args struct {
		input       service.RedirectInput
		OriginalURL string
	}
	tests := []struct {
		name    string
		args    args
		want    *domain.URLInfo
		wantErr bool
	}{
		{
			name: "SUCCESS_TEST_1",
			args: args{
				OriginalURL: "https://example.com/test1",
			},
			want: &domain.URLInfo{
				OriginalURL: "https://example.com/test1",
			},
			wantErr: false,
		},
		{
			name: "ERROR_TEST_2",
			want: &domain.URLInfo{
				OriginalURL: "",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			s := svc.NewService()
			shortURL, _ := s.CreateShortURL(ctx, service.CreateShortURLInput{
				OriginalURL: tt.args.OriginalURL,
			})
			tt.args.input.ShortURL = shortURL
			got, err := s.Redirect(ctx, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.Redirect() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(got.OriginalURL, tt.want.OriginalURL) {
				t.Errorf("service.Redirect() = %v, want %v", got, tt.want)
			}
		})
	}
}
