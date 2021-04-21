package v1

import (
	"fmt"
	"goarch/app/domain"
	"goarch/app/presentors/jsonapi"
	"io/ioutil"
	"net/http"
)

func itemsGetAll(_ map[string]string, conn domain.Connection, _ http.ResponseWriter, _ *http.Request) (int, []byte, error) {

	items, err := conn.Item().GetAll()
	if err != nil {
		return InternalServerError(err)
	}

	out, err := jsonapi.MarshalItems(items)

	if err != nil {
		return InternalServerError(err)
	}

	return OK(out)
}

func itemGet(v map[string]string, conn domain.Connection, _ http.ResponseWriter, _ *http.Request) (int, []byte, error) {

	id, ok := v["id"]
	if !ok {
		return BadRequest(fmt.Errorf("item id is empty"))
	}

	item, err := conn.Item().Get(id)

	if err != nil {
		return InternalServerError(err)
	}

	out, err := jsonapi.MarshalItem(item)

	if err != nil {
		return InternalServerError(err)
	}

	return OK(out)
}

func itemCreate(_ map[string]string, conn domain.Connection, _ http.ResponseWriter, r *http.Request) (int, []byte, error) {

	in, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return BadRequest(err)
	}

	item, err := jsonapi.UnmarshalItem(in)
	if err != nil {
		return BadRequest(err)
	}

	item.Id = ""

	item, err = conn.Item().Create(item)
	if err != nil {
		return InternalServerError(err)
	}

	out, err := jsonapi.MarshalItem(item)
	if err != nil {
		return InternalServerError(err)
	}

	return Created(out)
}

func itemDelete(v map[string]string, conn domain.Connection, _ http.ResponseWriter, _ *http.Request) (int, []byte, error) {
	id, ok := v["id"]
	if !ok {
		return BadRequest(fmt.Errorf("item id is empty"))
	}
	err := conn.Item().Delete(id)
	if err != nil {
		InternalServerError(err)
	}

	return NoContent()
}
