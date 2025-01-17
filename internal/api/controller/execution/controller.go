// Copyright 2023 Harness, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package execution

import (
	"github.com/harness/gitness/internal/auth/authz"
	"github.com/harness/gitness/internal/pipeline/canceler"
	"github.com/harness/gitness/internal/pipeline/commit"
	"github.com/harness/gitness/internal/pipeline/triggerer"
	"github.com/harness/gitness/internal/store"

	"github.com/jmoiron/sqlx"
)

type Controller struct {
	db             *sqlx.DB
	authorizer     authz.Authorizer
	executionStore store.ExecutionStore
	checkStore     store.CheckStore
	canceler       canceler.Canceler
	commitService  commit.CommitService
	triggerer      triggerer.Triggerer
	repoStore      store.RepoStore
	stageStore     store.StageStore
	pipelineStore  store.PipelineStore
}

func NewController(
	db *sqlx.DB,
	authorizer authz.Authorizer,
	executionStore store.ExecutionStore,
	checkStore store.CheckStore,
	canceler canceler.Canceler,
	commitService commit.CommitService,
	triggerer triggerer.Triggerer,
	repoStore store.RepoStore,
	stageStore store.StageStore,
	pipelineStore store.PipelineStore,
) *Controller {
	return &Controller{
		db:             db,
		authorizer:     authorizer,
		executionStore: executionStore,
		checkStore:     checkStore,
		canceler:       canceler,
		commitService:  commitService,
		triggerer:      triggerer,
		repoStore:      repoStore,
		stageStore:     stageStore,
		pipelineStore:  pipelineStore,
	}
}
