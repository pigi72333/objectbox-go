// Code generated by ObjectBox; DO NOT EDIT.

package perf

import (
	"github.com/google/flatbuffers/go"
	"github.com/objectbox/objectbox-go/objectbox"
	"github.com/objectbox/objectbox-go/objectbox/fbutils"
)

type entity_EntityInfo struct {
	Id  objectbox.TypeId
	Uid uint64
}

var EntityBinding = entity_EntityInfo{
	Id:  1,
	Uid: 1737161401460991620,
}

// Entity_ contains type-based Property helpers to facilitate some common operations such as Queries.
var Entity_ = struct {
	Id      *objectbox.PropertyUint64
	Int32   *objectbox.PropertyInt32
	Int64   *objectbox.PropertyInt64
	String  *objectbox.PropertyString
	Float64 *objectbox.PropertyFloat64
}{
	Id: &objectbox.PropertyUint64{
		Property: &objectbox.Property{
			Id: 1,
		},
	},
	Int32: &objectbox.PropertyInt32{
		Property: &objectbox.Property{
			Id: 2,
		},
	},
	Int64: &objectbox.PropertyInt64{
		Property: &objectbox.Property{
			Id: 3,
		},
	},
	String: &objectbox.PropertyString{
		Property: &objectbox.Property{
			Id: 4,
		},
	},
	Float64: &objectbox.PropertyFloat64{
		Property: &objectbox.Property{
			Id: 5,
		},
	},
}

// GeneratorVersion is called by the ObjectBox to verify the compatibility of the generator used to generate this code
func (entity_EntityInfo) GeneratorVersion() int {
	return 1
}

// AddToModel is called by the ObjectBox during model build
func (entity_EntityInfo) AddToModel(model *objectbox.Model) {
	model.Entity("Entity", 1, 1737161401460991620)
	model.Property("Id", objectbox.PropertyType_Long, 1, 7373286741377356014)
	model.PropertyFlags(objectbox.PropertyFlags_ID)
	model.Property("Int32", objectbox.PropertyType_Int, 2, 4837914178321008766)
	model.Property("Int64", objectbox.PropertyType_Long, 3, 3841825182616422591)
	model.Property("String", objectbox.PropertyType_String, 4, 6473251296493454829)
	model.Property("Float64", objectbox.PropertyType_Double, 5, 8933082277725371577)
	model.EntityLastPropertyId(5, 8933082277725371577)
}

// GetId is called by the ObjectBox during Put operations to check for existing ID on an object
func (entity_EntityInfo) GetId(object interface{}) (uint64, error) {
	return object.(*Entity).Id, nil
}

// SetId is called by the ObjectBox during Put to update an ID on an object that has just been inserted
func (entity_EntityInfo) SetId(object interface{}, id uint64) error {
	object.(*Entity).Id = id
	return nil
}

// Flatten is called by the ObjectBox to transform an object to a FlatBuffer
func (entity_EntityInfo) Flatten(object interface{}, fbb *flatbuffers.Builder, id uint64) {
	obj := object.(*Entity)
	var offsetString = fbutils.CreateStringOffset(fbb, obj.String)

	// build the FlatBuffers object
	fbb.StartObject(5)
	fbutils.SetUint64Slot(fbb, 0, id)
	fbutils.SetInt32Slot(fbb, 1, obj.Int32)
	fbutils.SetInt64Slot(fbb, 2, obj.Int64)
	fbutils.SetUOffsetTSlot(fbb, 3, offsetString)
	fbutils.SetFloat64Slot(fbb, 4, obj.Float64)
}

// ToObject is called by the ObjectBox to load an object from a FlatBuffer
func (entity_EntityInfo) ToObject(bytes []byte) interface{} {
	table := &flatbuffers.Table{
		Bytes: bytes,
		Pos:   flatbuffers.GetUOffsetT(bytes),
	}

	return &Entity{
		Id:      table.GetUint64Slot(4, 0),
		Int32:   table.GetInt32Slot(6, 0),
		Int64:   table.GetInt64Slot(8, 0),
		String:  fbutils.GetStringSlot(table, 10),
		Float64: table.GetFloat64Slot(12, 0),
	}
}

// MakeSlice is called by the ObjectBox to construct a new slice to hold the read objects
func (entity_EntityInfo) MakeSlice(capacity int) interface{} {
	return make([]*Entity, 0, capacity)
}

// AppendToSlice is called by the ObjectBox to fill the slice of the read objects
func (entity_EntityInfo) AppendToSlice(slice interface{}, object interface{}) interface{} {
	return append(slice.([]*Entity), object.(*Entity))
}

// Box provides CRUD access to Entity objects
type EntityBox struct {
	*objectbox.Box
}

// BoxForEntity opens a box of Entity objects
func BoxForEntity(ob *objectbox.ObjectBox) *EntityBox {
	return &EntityBox{
		Box: ob.InternalBox(1),
	}
}

// Put synchronously inserts/updates a single object.
// In case the Id is not specified, it would be assigned automatically (auto-increment).
// When inserting, the Entity.Id property on the passed object will be assigned the new ID as well.
func (box *EntityBox) Put(object *Entity) (uint64, error) {
	return box.Box.Put(object)
}

// PutAsync asynchronously inserts/updates a single object.
// When inserting, the Entity.Id property on the passed object will be assigned the new ID as well.
//
// It's executed on a separate internal thread for better performance.
//
// There are two main use cases:
//
// 1) "Put & Forget:" you gain faster puts as you don't have to wait for the transaction to finish.
//
// 2) Many small transactions: if your write load is typically a lot of individual puts that happen in parallel,
// this will merge small transactions into bigger ones. This results in a significant gain in overall throughput.
//
//
// In situations with (extremely) high async load, this method may be throttled (~1ms) or delayed (<1s).
// In the unlikely event that the object could not be enqueued after delaying, an error will be returned.
//
// Note that this method does not give you hard durability guarantees like the synchronous Put provides.
// There is a small time window (typically 3 ms) in which the data may not have been committed durably yet.
func (box *EntityBox) PutAsync(object *Entity) (uint64, error) {
	return box.Box.PutAsync(object)
}

// PutAll inserts multiple objects in single transaction.
// In case Ids are not set on the objects, they would be assigned automatically (auto-increment).
//
// Returns: IDs of the put objects (in the same order).
// When inserting, the Entity.Id property on the objects in the slice will be assigned the new IDs as well.
//
// Note: In case an error occurs during the transaction, some of the objects may already have the Entity.Id assigned
// even though the transaction has been rolled back and the objects are not stored under those IDs.
//
// Note: The slice may be empty or even nil; in both cases, an empty IDs slice and no error is returned.
func (box *EntityBox) PutAll(objects []*Entity) ([]uint64, error) {
	return box.Box.PutAll(objects)
}

// Get reads a single object.
//
// Returns nil (and no error) in case the object with the given ID doesn't exist.
func (box *EntityBox) Get(id uint64) (*Entity, error) {
	object, err := box.Box.Get(id)
	if err != nil {
		return nil, err
	} else if object == nil {
		return nil, nil
	}
	return object.(*Entity), nil
}

// Get reads all stored objects
func (box *EntityBox) GetAll() ([]*Entity, error) {
	objects, err := box.Box.GetAll()
	if err != nil {
		return nil, err
	}
	return objects.([]*Entity), nil
}

// Remove deletes a single object
func (box *EntityBox) Remove(object *Entity) (err error) {
	return box.Box.Remove(object.Id)
}

// Creates a query with the given conditions. Use the fields of the Entity_ struct to create conditions.
// Keep the *EntityQuery if you intend to execute the query multiple times.
// Note: this function panics if you try to create illegal queries; e.g. use properties of an alien type.
// This is typically a programming error. Use QueryOrError instead if you want the explicit error check.
func (box *EntityBox) Query(conditions ...objectbox.Condition) *EntityQuery {
	return &EntityQuery{
		box.Box.Query(conditions...),
	}
}

// Creates a query with the given conditions. Use the fields of the Entity_ struct to create conditions.
// Keep the *EntityQuery if you intend to execute the query multiple times.
func (box *EntityBox) QueryOrError(conditions ...objectbox.Condition) (*EntityQuery, error) {
	if query, err := box.Box.QueryOrError(conditions...); err != nil {
		return nil, err
	} else {
		return &EntityQuery{query}, nil
	}
}

// Query provides a way to search stored objects
//
// For example, you can find all Entity which Id is either 42 or 47:
// 		box.Query(Entity_.Id.In(42, 47)).Find()
type EntityQuery struct {
	*objectbox.Query
}

// Find returns all objects matching the query
func (query *EntityQuery) Find() ([]*Entity, error) {
	objects, err := query.Query.Find()
	if err != nil {
		return nil, err
	}
	return objects.([]*Entity), nil
}
