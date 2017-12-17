// Copyright 2017 Nicholas Ng <nickng@nickng.io>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sesstype

import (
	"fmt"
)

// Role represents an endpoint or participant in a session interaction.
type Role struct {
	name string
}

// Name returns the identifying name of the Role.
func (r Role) Name() string {
	return r.name
}

// NewRole returns a Role value with the given name.
func NewRole(name string) Role {
	return Role{name: name}
}

// Message represents a message passed between Roles in a session interaction.
type Message struct {
	Label   string  // Message label
	Payload Payload // Freeform string
}

func (m Message) String() string {
	if m.Payload == nil {
		return fmt.Sprintf("%s()", m.Label)
	}
	return fmt.Sprintf("%s(%s)", m.Label, m.Payload.String())
}

type Payload interface {
	String() string
}

type BaseType struct {
	Type string
}

func (t BaseType) String() string {
	return t.Type
}
