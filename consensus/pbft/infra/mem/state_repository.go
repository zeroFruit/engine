/*
 * Copyright 2018 It-chain
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package mem

import (
	"sync"

	"github.com/it-chain/engine/consensus/pbft"
)

type StateRepository struct {
	state pbft.State
	sync.RWMutex
}

func NewStateRepository() *StateRepository {
	return &StateRepository{
		state:   pbft.State{},
		RWMutex: sync.RWMutex{},
	}
}

func (repo *StateRepository) Save(state pbft.State) error {

	repo.Lock()
	defer repo.Unlock()
	id := repo.state.StateID.ID
	if id == state.StateID.ID || id == "" {
		repo.state = state
		return nil
	}
	return pbft.ErrInvalidSave
}

func (repo *StateRepository) Load() (pbft.State, error) {
	repo.Lock()
	defer repo.Unlock()

	if repo.state.StateID.ID == "" {
		return repo.state, pbft.ErrEmptyRepo
	}

	return repo.state, nil
}

func (repo *StateRepository) Remove() {
	repo.state = pbft.State{}
}
