package tests

import (
	"context"
	"reflect"
	"testing"

	svc "url-shortener/internal/shortener"
	"url-shortener/internal/shortener/domain"
	"url-shortener/internal/shortener/service"
)

func TestGetURLs(t *testing.T) {
	type args struct {
		OriginalURL string
	}
	tests := []struct {
		name    string
		args    []args
		want    *service.GetURLsResponse
		wantErr bool
	}{
		{
			name: "SUCCESS_TEST_1",
			args: []args{
				{
					OriginalURL: "https://example/com/test1",
				},
				{
					OriginalURL: "https://example/com/test2",
				},
				{
					OriginalURL: "https://example/com/test3",
				},
			},
			want: &service.GetURLsResponse{
				Total: 3,
				URLs: []domain.URLs{
					{
						OriginalURL: "https://example/com/test1",
					},
					{
						OriginalURL: "https://example/com/test2",
					},
					{
						OriginalURL: "https://example/com/test3",
					},
				},
			},
			wantErr: false,
		},
		{
			name:    "ERROR_TEST_2",
			args:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			s := svc.NewService()
			for i, arg := range tt.args {
				shortURL, _ := s.CreateShortURL(ctx, service.CreateShortURLInput{
					OriginalURL: arg.OriginalURL,
				})
				tt.want.URLs[i].ShortURL = shortURL
			}
			got, err := s.GetURLs(ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.GetURLs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.GetURLs() = %v, want %v", got, tt.want)
			}
		})
	}
}
