package gopager

import "reflect"

// Pageable is the interface that wraps Len method.
//
// Len returns the length of slice.
type Pageable interface {
	Len() int
}

// Paginater is the interface that supports retrieving iterator items a page at a time.
type Paginater interface {
	Next(slicePtr interface{})
	Current(slicePtr interface{})
	Previous(slicePtr interface{})
	HasNext() bool
	HasPrevious() bool
	Page(num int) Paginater
	CurrentPage() int
	MaxPage() int
}

// Paginate is the struct that has page information necessary for paging.
type Paginate struct {
	items       Pageable
	pageSize    int
	current     int
	maxSize     int
	currentPage int
	maxPage     int
}

// NewPaginater returns a paginater.
func NewPaginater(items Pageable, pageSize int) Paginater {
	p := &Paginate{
		items:       items,
		pageSize:    pageSize,
		current:     0,
		maxSize:     items.Len(),
		currentPage: 0,
	}
	if p.maxSize == 0 {
		p.maxPage = 0
	} else {
		p.maxPage = ((p.maxSize - 1) / pageSize) + 1
	}
	return p
}

// Page changes the current page to the number specified in the argument `num`.
func (p *Paginate) Page(num int) Paginater {
	if num < 0 {
		num = 0
	}
	p.currentPage = num
	p.current = num * p.pageSize
	return p
}

// CurrentPage returns the current page number.
func (p *Paginate) CurrentPage() int {
	return p.currentPage
}

// MaxPage returns the maxinum value of pages.
func (p *Paginate) MaxPage() int {
	return p.maxPage
}

// Next retrieves a sequence of items in next page from the iterator
// and appends them to slicePtr, which must be a pointer to a slice
// of the iterator's item type.
func (p *Paginate) Next(slicePtr interface{}) {
	if !p.HasNext() {
		return
	}

	e := reflect.ValueOf(slicePtr).Elem()

	from := p.current
	to := p.current + p.pageSize
	if to > p.maxSize {
		to = p.maxSize
	}

	e.Set(reflect.AppendSlice(e, reflect.ValueOf(p.items).Slice(from, to)))

	p.current = p.current + p.pageSize

	p.currentPage++
}

// Current retrieves a sequence of items in current page from the iterator
// and appends them to slicePtr, which must be a pointer to a slice
// of the iterator's item type.
func (p *Paginate) Current(slicePtr interface{}) {
	if p.CurrentPage() == 0 {
		return
	}

	e := reflect.ValueOf(slicePtr).Elem()

	from := p.current - p.pageSize
	if from < 0 {
		from = 0
	}
	to := p.current
	if to > p.maxSize {
		to = p.maxSize
	}

	e.Set(reflect.AppendSlice(e, reflect.ValueOf(p.items).Slice(from, to)))
}

// Previous retrieves a sequence of items in previous page from the iterator
// and appends them to slicePtr, which must be a pointer to a slice
// of the iterator's item type.
func (p *Paginate) Previous(slicePtr interface{}) {
	if !p.HasPrevious() {
		return
	}

	e := reflect.ValueOf(slicePtr).Elem()

	from := p.current - p.pageSize*2
	if from < 0 {
		from = 0
	}
	to := p.current - p.pageSize
	if to > p.maxSize {
		to = p.maxSize
	}

	e.Set(reflect.AppendSlice(e, reflect.ValueOf(p.items).Slice(from, to)))

	p.currentPage--

	p.current = to
}

// HasNext returns whether the next page exists.
func (p *Paginate) HasNext() bool {
	if p.current >= p.maxSize {
		return false
	}
	return true
}

// HasPrevious returns whether the previous page exists.
func (p *Paginate) HasPrevious() bool {
	if p.current-p.pageSize <= 0 {
		return false
	}
	return true
}
