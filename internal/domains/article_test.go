package domains

import (
	"reflect"
	"testing"
	"time"
)

func TestArticle_Change(t *testing.T) {
	createdAt, _ := time.Parse(time.RFC3339, "2000-01-01T12:34:56+00:00")
	updatedAt, _ := time.Parse(time.RFC3339, "2000-01-01T12:34:56+00:00")

	type fields struct {
		ID        int64
		AuthorID  int64
		Title     string
		Content   string
		CreatedAt time.Time
		UpdatedAt time.Time
	}
	type args struct {
		item Article
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Article
	}{
		{
			name: "change title and content",
			fields: fields{
				ID:        1,
				AuthorID:  1000000001,
				Title:     "test_title_1",
				Content:   "test_content_1",
				CreatedAt: createdAt,
				UpdatedAt: updatedAt,
			},
			args: args{
				item: Article{
					ID:        1,
					AuthorID:  1000000001,
					Title:     "test_title_1_update",
					Content:   "test_content_1_update",
					CreatedAt: createdAt,
					UpdatedAt: updatedAt,
				},
			},
			want: &Article{
				ID:        1,
				AuthorID:  1000000001,
				Title:     "test_title_1_update",
				Content:   "test_content_1_update",
				CreatedAt: createdAt,
				UpdatedAt: updatedAt,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Article{
				ID:        tt.fields.ID,
				AuthorID:  tt.fields.AuthorID,
				Title:     tt.fields.Title,
				Content:   tt.fields.Content,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
			}
			if got := a.Change(tt.args.item); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Article.Change() = %v, want %v", got, tt.want)
			}
		})
	}
}
