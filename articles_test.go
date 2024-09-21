package goshopify

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestArticleList(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("https://fooshop.myshopify.com/%s/blogs/241253187/articles.json", client.pathPrefix),
		httpmock.NewStringResponder(
			200,
			`{"articles": [{"id":1},{"id":2}]}`,
		),
	)

	articles, err := client.Article.List(context.Background(), 241253187, nil)
	if err != nil {
		t.Errorf("Article.List returned error: %v", err)
	}

	expected := []Article{
		{
			Id: 1,
		},
		{
			Id: 2,
		},
	}
	if !reflect.DeepEqual(articles, expected) {
		t.Errorf("Articles.List returned %+v, expected %+v", articles, expected)
	}
}

func TestArticleCreate(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder(
		"POST",
		fmt.Sprintf("https://fooshop.myshopify.com/%s/blogs/241253187/articles.json", client.pathPrefix),
		httpmock.NewStringResponder(
			201,
			`{"article": {"id": 1}}`,
		),
	)

	article := Article{Title: "Test Article"}
	createdArticle, err := client.Article.Create(context.Background(), 241253187, article)
	if err != nil {
		t.Errorf("Article.Create returned error: %v", err)
	}

	expected := &Article{Id: 1}
	if !reflect.DeepEqual(createdArticle, expected) {
		t.Errorf("Article.Create returned %+v, expected %+v", createdArticle, expected)
	}
}

func TestArticleGet(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("https://fooshop.myshopify.com/%s/blogs/241253187/articles/1.json", client.pathPrefix),
		httpmock.NewStringResponder(
			200,
			`{"article": {"id": 1, "title": "Test Article"}}`,
		),
	)

	article, err := client.Article.Get(context.Background(), 241253187, 1)
	if err != nil {
		t.Errorf("Article.Get returned error: %v", err)
	}

	expected := &Article{Id: 1, Title: "Test Article"}
	if !reflect.DeepEqual(article, expected) {
		t.Errorf("Article.Get returned %+v, expected %+v", article, expected)
	}
}

func TestArticleUpdate(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder(
		"PUT",
		fmt.Sprintf("https://fooshop.myshopify.com/%s/blogs/241253187/articles/1.json", client.pathPrefix),
		httpmock.NewStringResponder(
			200,
			`{"article": {"id": 1, "title": "Updated Article"}}`,
		),
	)

	article := Article{Title: "Updated Article"}
	updatedArticle, err := client.Article.Update(context.Background(), 241253187, 1, article)
	if err != nil {
		t.Errorf("Article.Update returned error: %v", err)
	}

	expected := &Article{Id: 1, Title: "Updated Article"}
	if !reflect.DeepEqual(updatedArticle, expected) {
		t.Errorf("Article.Update returned %+v, expected %+v", updatedArticle, expected)
	}
}

func TestArticleDelete(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder(
		"DELETE",
		fmt.Sprintf("https://fooshop.myshopify.com/%s/blogs/241253187/articles/1.json", client.pathPrefix),
		httpmock.NewStringResponder(
			204, // No content response
			``,
		),
	)

	err := client.Article.Delete(context.Background(), 241253187, 1)
	if err != nil {
		t.Errorf("Article.Delete returned error: %v", err)
	}
}

func TestArticleListTags(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("https://fooshop.myshopify.com/%s/articles/tags.json", client.pathPrefix),
		httpmock.NewStringResponder(
			200,
			`{"tags": ["tag1", "tag2"]}`,
		),
	)

	tags, err := client.Article.ListTags(context.Background(), nil)
	if err != nil {
		t.Errorf("Article.ListTags returned error: %v", err)
	}

	expected := []string{"tag1", "tag2"}
	if !reflect.DeepEqual(tags, expected) {
		t.Errorf("Article.ListTags returned %+v, expected %+v", tags, expected)
	}
}

func TestArticleCount(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("https://fooshop.myshopify.com/%s/blogs/241253187/articles/count.json", client.pathPrefix),
		httpmock.NewStringResponder(
			200,
			`{"count": 2}`,
		),
	)

	count, err := client.Article.Count(context.Background(), 241253187, nil)
	if err != nil {
		t.Errorf("Article.Count returned error: %v", err)
	}

	expected := 2
	if count != expected {
		t.Errorf("Article.Count returned %d, expected %d", count, expected)
	}
}

func TestArticleListBlogTags(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder(
		"GET",
		fmt.Sprintf("https://fooshop.myshopify.com/%s/blogs/241253187/articles/tags.json", client.pathPrefix),
		httpmock.NewStringResponder(
			200,
			`{"tags": ["blogTag1", "blogTag2"]}`,
		),
	)

	tags, err := client.Article.ListBlogTags(context.Background(), 241253187, nil)
	if err != nil {
		t.Errorf("Article.ListBlogTags returned error: %v", err)
	}

	expected := []string{"blogTag1", "blogTag2"}
	if !reflect.DeepEqual(tags, expected) {
		t.Errorf("Article.ListBlogTags returned %+v, expected %+v", tags, expected)
	}
}
