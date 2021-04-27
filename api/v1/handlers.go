package v1

import (
	"goarch/app/domain"
	"goarch/app/domain/cases/item_create"
	"goarch/app/domain/cases/item_delete"
	"goarch/app/domain/cases/item_get"
	"goarch/app/domain/cases/items_get"
	"goarch/app/presentors/jsonapi"
	"io/ioutil"
	"net/http"
)

func itemsGetAll(_ map[string]string, conn domain.Connection, _ http.ResponseWriter, _ *http.Request) (int, []byte, error) {

	items, err := items_get.Run(conn)
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

	item, err := item_get.Run(conn, v["id"])
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

	item, err = item_create.Run(conn, item)
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

	err := item_delete.Run(conn, v["id"])
	if err != nil {
		InternalServerError(err)
	}

	return OK(nil)
}
