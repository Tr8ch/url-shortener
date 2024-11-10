package tests

import (
	"context"
	"testing"

	svc "url-shortener/internal/shortener"
	"url-shortener/internal/shortener/service"
)

func TestDeleteURL(t *testing.T) {
	type args struct {
		input       service.DeleteShortURLInput
		OriginalURL string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "SUCCESS_TEST_1",
			args: args{
				OriginalURL: "https://example.com/test1",
			},
			wantErr: false,
		},
		{
			name:    "ERROR_TEST_2",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			s := svc.NewService()
			shortURL, _ := s.CreateShortURL(ctx, service.CreateShortURLInput{OriginalURL: tt.args.OriginalURL})
			tt.args.input.ShortURL = shortURL
			if err := s.DeleteURL(ctx, tt.args.input); (err != nil) != tt.wantErr {
				t.Errorf("service.DeleteURL() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
