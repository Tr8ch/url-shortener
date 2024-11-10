package tests

import (
	"context"
	"testing"

	svc "url-shortener/internal/shortener"
	"url-shortener/internal/shortener/service"
)

func TestCreateShortURL(t *testing.T) {
	type args struct {
		input service.CreateShortURLInput
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "SUCCESS_TEST_1",
			args: args{
				input: service.CreateShortURLInput{
					OriginalURL: "https://example.com/test1",
				},
			},
			wantErr: false,
		},
		{
			name: "ERROR_TEST_2",
			args: args{
				input: service.CreateShortURLInput{
					OriginalURL: "",
				},
			},
			wantErr: true,
		},
		{
			name: "ERROR_TEST_3",
			args: args{
				input: service.CreateShortURLInput{
					OriginalURL: "errors",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := svc.NewService()
			_, err := s.CreateShortURL(context.Background(), tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.CreateShortURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
