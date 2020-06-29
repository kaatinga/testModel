package repository

import (
	"reflect"
	"testing"

	"shop/models"
)

func TestNewMapDB(t *testing.T) {
	tests := []struct {
		name string
		want Repository
	}{
		{"ok", &mapDB{db: make(map[int32]*models.Item)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMapDB(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMapDB() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mapDB_CreateItem(t *testing.T) {

	var inputItem models.Item
	outputItem := models.Item{ID: 1}

	tests := []struct {
		name    string
		object  mapDB
		arg     *models.Item
		want    *models.Item
		wantErr bool
	}{
		{"ok", mapDB{db: make(map[int32]*models.Item), maxID: 0}, &inputItem, &outputItem, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			m := &tt.object

			got, err := m.CreateItem(tt.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateItem() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateItem() got = %v, want %v", got, tt.want)
			}
		})
	}
}

//func Test_mapDB_DeleteItem(t *testing.T) {
//	type fields struct {
//		db    map[int32]*models.Item
//		maxID int32
//	}
//	type args struct {
//		ID int32
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			m := &mapDB{
//				db:    tt.fields.db,
//				maxID: tt.fields.maxID,
//			}
//			if err := m.DeleteItem(tt.args.ID); (err != nil) != tt.wantErr {
//				t.Errorf("DeleteItem() error = %v, wantErr %v", err, tt.wantErr)
//			}
//		})
//	}
//}
//
//func Test_mapDB_GetItem(t *testing.T) {
//	type fields struct {
//		db    map[int32]*models.Item
//		maxID int32
//	}
//	type args struct {
//		ID int32
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		want    *models.Item
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			m := &mapDB{
//				db:    tt.fields.db,
//				maxID: tt.fields.maxID,
//			}
//			got, err := m.GetItem(tt.args.ID)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("GetItem() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("GetItem() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
func Test_mapDB_UpdateItem(t *testing.T) {

	var testItem models.Item
	testItem.ID = 1
	testItem.Name = "test"
	testItem.Price = 100

	var incorrectItem models.Item
	incorrectItem.ID = 2

	var fields mapDB
	fields.db = make(map[int32]*models.Item)
	fields.maxID = 1

	fields.db[1] = &testItem

	tests := []struct {
		name    string
		fields  mapDB
		args    *models.Item
		want    *models.Item
		wantErr bool
	}{
		{"ok", fields, &testItem, &testItem, false},
		{"wrong", fields, &incorrectItem, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &mapDB{
				db:    tt.fields.db,
				maxID: tt.fields.maxID,
			}
			got, err := m.UpdateItem(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateItem() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateItem() got = %v, want %v", got, tt.want)
			}
		})
	}
}
