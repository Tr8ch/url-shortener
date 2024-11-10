package tests

import (
	"context"
	"reflect"
	"testing"

	svc "url-shortener/internal/shortener"
	"url-shortener/internal/shortener/domain"
	"url-shortener/internal/shortener/service"
)

func TestGetStats(t *testing.T) {
	type args struct {
		input       service.GetStatsInput
		OriginalURL string
		ClickCounts int
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
				ClickCounts: 3,
			},
			want: &domain.URLInfo{
				OriginalURL: "https://example.com/test1",
				ClickCounts: 3,
			},
			wantErr: false,
		},
		{
			name:    "ERROR_TEST_2",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		ctx := context.Background()
		t.Run(tt.name, func(t *testing.T) {
			s := svc.NewService()
			shortURL, _ := s.CreateShortURL(ctx, service.CreateShortURLInput{
				OriginalURL: tt.args.OriginalURL,
			})
			tt.args.input.ShortURL = shortURL
			for range tt.args.ClickCounts {
				urlInfo, _ := s.Redirect(ctx, service.RedirectInput{
					ShortURL: shortURL,
				})
				tt.want.CreatedAt = urlInfo.CreatedAt
				tt.want.LastEnteredAt = urlInfo.LastEnteredAt
			}

			got, err := s.GetStats(ctx, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.GetStats() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.GetStats() = %v, want %v", got, tt.want)
			}
		})
	}
}
