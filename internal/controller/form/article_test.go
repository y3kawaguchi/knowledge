package form

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/go-cmp/cmp"
	"github.com/y3kawaguchi/knowledge/internal/domains"
	"github.com/y3kawaguchi/knowledge/test/testutils"
)

func TestArticle(t *testing.T) {

	cases := map[string]struct {
		body       []byte
		wantErr    string
		wantDomain *domains.Article
	}{
		"validation check ok": {
			body:       testutils.GetBytesFromFile("testdata/article_req.json"),
			wantErr:    "",
			wantDomain: func() *domains.Article { return buildDefaultArticle() }(),
		},
		// "without author_id": {
		// 	body: []byte,
		// 	wantErr: string,
		// },

	}

	gin.SetMode(gin.ReleaseMode)
	for k, tc := range cases {
		t.Run(k, func(t *testing.T) {
			gc, _ := gin.CreateTestContext(httptest.NewRecorder())
			gc.Request, _ = http.NewRequest("POST", "/articles", bytes.NewBuffer(tc.body))
			article := Article{}
			err := gc.ShouldBindJSON(&article)
			if err == nil {
				if diff := cmp.Diff(tc.wantErr, ""); diff != "" {
					t.Errorf("%s: failed (-want +got):\n%s", k, diff)
				}
				if tc.wantDomain != nil {
					if diff := cmp.Diff(tc.wantDomain, article.BuildDomain()); diff != "" {
						t.Errorf("%s: failed (-want +got):\n%s", k, diff)
					}
				}
			} else {
				if diff := cmp.Diff(tc.wantErr, fmt.Sprintf("%v", err.Error())); diff != "" {
					t.Errorf("%s: failed (-want +got):\n%s", k, diff)
				}
			}
		})
	}
}

func buildDefaultArticle() *domains.Article {
	return &domains.Article{
		AuthorID: 1,
		Title:    "TestTitle",
		Content:  "TestContent",
	}
}
