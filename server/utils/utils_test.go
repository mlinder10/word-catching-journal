package utils

import (
	"net/http/httptest"
	"strings"
	"testing"
)

const url = "https://test.com"

func TestParsePageQuery(t *testing.T) {
	t.Run("no page", func(t *testing.T) {
		r := httptest.NewRequest("GET", url, nil)
		page := ParsePageQuery(r)
		if page != 0 {
			t.Errorf("expected 0, got %d", page)
		}
	})

	t.Run("page 1", func(t *testing.T) {
		r := httptest.NewRequest("GET", url+"?page=1", nil)
		page := ParsePageQuery(r)
		if page != 0 {
			t.Errorf("expected 0, got %d", page)
		}
	})

	t.Run("page 10", func(t *testing.T) {
		r := httptest.NewRequest("GET", url+"?page=10", nil)
		page := ParsePageQuery(r)
		if page != 9 {
			t.Errorf("expected 9, got %d", page)
		}
	})

	t.Run("page -1", func(t *testing.T) {
		r := httptest.NewRequest("GET", url+"?page=-1", nil)
		page := ParsePageQuery(r)
		if page != 0 {
			t.Errorf("expected 0, got %d", page)
		}
	})

	t.Run("page test", func(t *testing.T) {
		r := httptest.NewRequest("GET", url+"?page=test", nil)
		page := ParsePageQuery(r)
		if page != 0 {
			t.Errorf("expected 0, got %d", page)
		}
	})
}

func TestRandomColor(t *testing.T) {
	color := RandomColor()

	if len(color) != 7 {
		t.Errorf("expected 7, got %d", len(color))
	}

	if color[0] != '#' {
		t.Errorf("expected #, got %s", string(color[0]))
	}

	for i := range 6 {
		if !strings.Contains("0123456789abcdef", string(color[i+1])) {
			t.Errorf("expected 0123456789abcdef, got %s", string(color[i+1]))
		}
	}
}
