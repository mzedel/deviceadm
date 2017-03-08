// Copyright 2016 Mender Software AS
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.
package main

import "github.com/stretchr/testify/mock"
import "github.com/mendersoftware/go-lib-micro/mongo/migrate"

// MockInventoryApp is an autogenerated mock type for the InventoryApp type
type MockDataStore struct {
	mock.Mock
}

func (_m *MockDataStore) PutDevice(d *Device) error {
	ret := _m.Called(d)

	return ret.Error(0)
}

func (_m *MockDataStore) GetDevice(id AuthID) (*Device, error) {
	ret := _m.Called(id)

	dif := ret.Get(0)
	if dif != nil {
		return dif.(*Device), ret.Error(1)
	}
	return nil, ret.Error(1)
}

func (_m *MockDataStore) GetDevices(skip, limit int, status string) ([]Device, error) {
	ret := _m.Called(skip, limit, status)

	return ret.Get(0).([]Device), ret.Error(1)
}

func (_m *MockDataStore) Migrate(version string, migrations []migrate.Migration) error {
	ret := _m.Called(version, migrations)

	return ret.Error(0)
}

func (_m *MockDataStore) EnsureIndexes() error {
	ret := _m.Called()

	return ret.Error(0)
}

func (_m *MockDataStore) DeleteDevice(id AuthID) error {
	ret := _m.Called(id)

	return ret.Error(0)
}
