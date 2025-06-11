// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pagination

import (
	"net/http"

	"github.com/m3ter-com/m3ter-sdk-go/internal/apijson"
	"github.com/m3ter-com/m3ter-sdk-go/internal/requestconfig"
	"github.com/m3ter-com/m3ter-sdk-go/option"
)

type Cursor[T any] struct {
	Data      []T        `json:"data"`
	NextToken string     `json:"nextToken,nullable"`
	JSON      cursorJSON `json:"-"`
	cfg       *requestconfig.RequestConfig
	res       *http.Response
}

// cursorJSON contains the JSON metadata for the struct [Cursor[T]]
type cursorJSON struct {
	Data        apijson.Field
	NextToken   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Cursor[T]) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cursorJSON) RawJSON() string {
	return r.raw
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *Cursor[T]) GetNextPage() (res *Cursor[T], err error) {
	next := r.NextToken
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	err = cfg.Apply(option.WithQuery("nextToken", next))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *Cursor[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &Cursor[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type CursorAutoPager[T any] struct {
	page *Cursor[T]
	cur  T
	idx  int
	run  int
	err  error
}

func NewCursorAutoPager[T any](page *Cursor[T], err error) *CursorAutoPager[T] {
	return &CursorAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *CursorAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Data) == 0 {
		return false
	}
	if r.idx >= len(r.page.Data) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Data) == 0 {
			return false
		}
	}
	r.cur = r.page.Data[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *CursorAutoPager[T]) Current() T {
	return r.cur
}

func (r *CursorAutoPager[T]) Err() error {
	return r.err
}

func (r *CursorAutoPager[T]) Index() int {
	return r.run
}
