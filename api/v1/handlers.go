package v1

import (
	"fmt"
	"goarch/app/domain"
	"goarch/app/presentors/json"
	"io/ioutil"
	"net/http"
)

func itemsGetAll(_ domain.RouteVars, c domain.Context, _ *http.Request, p domain.Presenter) (int, error) {

	items, err := c.Connection().ItemRepository().GetAll()
	if err != nil {
		return InternalServerError(err)
	}

	if err := p.Items(items); err != nil {
		return InternalServerError(err)
	}

	return OK()
}

func itemGet(v domain.RouteVars, c domain.Context, r *http.Request, p domain.Presenter) (int, error) {

	id, ok := v.Get("id")
	if !ok {
		return BadRequest(fmt.Errorf("item id is empty"))
	}

	item, err := c.Connection().ItemRepository().Get(id)
	if err != nil {
		return InternalServerError(err)
	}

	if err := p.Item(item); err != nil {
		return InternalServerError(err)
	}

	return OK()
}

func itemCreate(_ domain.RouteVars, c domain.Context, r *http.Request, p domain.Presenter) (int, error) {

	in, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return BadRequest(err)
	}

	item, err := json.UnmarshalItem(in)
	if err != nil {
		return BadRequest(err)
	}

	item.Id = ""

	item, err = c.Connection().ItemRepository().Create(item)
	if err != nil {
		return InternalServerError(err)
	}

	if err := p.Item(item); err != nil {
		return InternalServerError(err)
	}

	return Created()
}

func itemDelete(v domain.RouteVars, c domain.Context, _ *http.Request, p domain.Presenter) (int, error) {
	id, ok := v.Get("id")
	if !ok {
		return BadRequest(fmt.Errorf("item id is empty"))
	}
	err := c.Connection().ItemRepository().Delete(id)
	if err != nil {
		InternalServerError(err)
	}

	return NoContent()
}
