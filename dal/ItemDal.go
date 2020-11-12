package dal

import (
	"connect-to-mysql/helper"
	"connect-to-mysql/model"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//GetAllItems ...
func GetAllItems(writer http.ResponseWriter, request *http.Request) {
	db := helper.GetConnection()
	var item []model.Item
	result := db.Find(&item)
	if result.Error != nil {
		responseWithJSON(writer, http.StatusBadRequest, map[string]string{"message": "Can't get all items1"})
		return
	}
	responseWithJSON(writer, http.StatusOK, item)
}

// GetItemByID ...
func GetItemByID(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		responseWithJSON(writer, http.StatusBadRequest, map[string]string{"message": "Invalid item id"})
		return
	}
	var item model.Item
	db := helper.GetConnection()
	result := db.First(&item, id)
	if result.Error != nil {
		responseWithJSON(writer, http.StatusBadRequest, map[string]string{"message": "Can't get item!"})
		return
	}
	responseWithJSON(writer, http.StatusOK, item)
}

//CreateItem ...
func CreateItem(writer http.ResponseWriter, request *http.Request) {
	db := helper.GetConnection()
	var newItem model.Item
	if err := json.NewDecoder(request.Body).Decode(&newItem); err != nil {
		responseWithJSON(writer, http.StatusBadRequest, map[string]string{"message": "Invalid body"})
		return
	}
	result := db.Create(newItem)
	if result.Error != nil {
		responseWithJSON(writer, http.StatusBadRequest, map[string]string{"message": "Error create item"})
		return
	}
	responseWithJSON(writer, http.StatusOK, newItem)
}

//UpdateItem ...
func UpdateItem(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	var updateItem model.Item

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		responseWithJSON(writer, http.StatusBadRequest, map[string]string{"message": "Invalid item id"})
		return
	}
	if err2 := json.NewDecoder(request.Body).Decode(&updateItem); err2 != nil {
		responseWithJSON(writer, http.StatusBadRequest, map[string]string{"message": "Invalid body"})
		return
	}
	db := helper.GetConnection()

	result := db.Where("item_id = ?", id).Updates(&updateItem)
	if result.Error != nil {
		responseWithJSON(writer, http.StatusNotFound, map[string]string{"message": "Update not success!"})
		return
	}
	responseWithJSON(writer, http.StatusNotFound, map[string]string{"message": "Update successfully!"})
}

//DeleteItem ...
func DeleteItem(writer http.ResponseWriter, request *http.Request) {
	db := helper.GetConnection()
	params := mux.Vars(request)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		responseWithJSON(writer, http.StatusBadRequest, map[string]string{"message": "Invalid item id"})
		return
	}

	var item model.Item
	result := db.Where("item_id = ?", id).Delete(&item)
	if result.Error != nil {
		responseWithJSON(writer, http.StatusNotFound, map[string]string{"message": "Can't delete item"})
		return
	}
	responseWithJSON(writer, http.StatusOK, map[string]string{"message": "Item was deleted"})
}

func responseWithJSON(writer http.ResponseWriter, status int, object interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)
	json.NewEncoder(writer).Encode(object)
}
