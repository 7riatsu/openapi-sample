/*
Open API Sample

This is a sample code.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// Pet struct for Pet
type Pet struct {
	Name string `json:"name"`
	Tag *string `json:"tag,omitempty"`
	Id int64 `json:"id"`
}

// NewPet instantiates a new Pet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPet(name string, id int64) *Pet {
	this := Pet{}
	this.Name = name
	this.Id = id
	return &this
}

// NewPetWithDefaults instantiates a new Pet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPetWithDefaults() *Pet {
	this := Pet{}
	return &this
}

// GetName returns the Name field value
func (o *Pet) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Pet) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Pet) SetName(v string) {
	o.Name = v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *Pet) GetTag() string {
	if o == nil || o.Tag == nil {
		var ret string
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pet) GetTagOk() (*string, bool) {
	if o == nil || o.Tag == nil {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *Pet) HasTag() bool {
	if o != nil && o.Tag != nil {
		return true
	}

	return false
}

// SetTag gets a reference to the given string and assigns it to the Tag field.
func (o *Pet) SetTag(v string) {
	o.Tag = &v
}

// GetId returns the Id field value
func (o *Pet) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Pet) GetIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Pet) SetId(v int64) {
	o.Id = v
}

func (o Pet) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Tag != nil {
		toSerialize["tag"] = o.Tag
	}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullablePet struct {
	value *Pet
	isSet bool
}

func (v NullablePet) Get() *Pet {
	return v.value
}

func (v *NullablePet) Set(val *Pet) {
	v.value = val
	v.isSet = true
}

func (v NullablePet) IsSet() bool {
	return v.isSet
}

func (v *NullablePet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePet(val *Pet) *NullablePet {
	return &NullablePet{value: val, isSet: true}
}

func (v NullablePet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

