/*
Copyright 2021 The OpenYurt Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package clients

// import "errors"
import "strings"

type NotFoundError struct{}

func (e *NotFoundError) Error() string { return "Item not found" }

func IsNotFoundErr(err error) bool {
	return err.Error() == "Item not found" || strings.HasPrefix(err.Error(), "no item found")
	// return errors.Is(err, &NotFoundError{})
}
