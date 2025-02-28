/*
Copyright 2023 Huawei Cloud Computing Technologies Co., Ltd.

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

package sparseindex

import (
	"github.com/openGemini/openGemini/engine/hybridqp"
	"github.com/openGemini/openGemini/lib/index"
	"github.com/openGemini/openGemini/lib/record"
	"github.com/openGemini/openGemini/lib/rpn"
)

var _ = RegistrySKFileReaderCreator(index.SetIndex, &SetReaderCreator{})

type SetReaderCreator struct {
}

func (index *SetReaderCreator) CreateSKFileReader(rpnExpr *rpn.RPNExpr, schema record.Schemas, option hybridqp.Options, isCache bool) (SKFileReader, error) {
	return NewSetIndexReader(rpnExpr, schema, option, isCache)
}

type SetIndexReader struct {
	init    bool
	isCache bool
	schema  record.Schemas
	option  hybridqp.Options
	sk      SKCondition
}

func NewSetIndexReader(rpnExpr *rpn.RPNExpr, schema record.Schemas, option hybridqp.Options, isCache bool) (*SetIndexReader, error) {
	sk, err := NewSKCondition(rpnExpr, schema)
	if err != nil {
		return nil, err
	}
	return &SetIndexReader{schema: schema, option: option, isCache: isCache, sk: sk}, nil
}

func (r *SetIndexReader) MayBeInFragment(fragId uint32) (bool, error) {
	return false, nil
}

func (r *SetIndexReader) ReInit(file interface{}) (err error) {
	if !r.init {
		r.init = true
		return
	}
	return
}

func (r *SetIndexReader) Close() error {
	return nil
}
