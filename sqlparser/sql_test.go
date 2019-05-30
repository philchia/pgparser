// Copyright 2016 The kingshard Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package sqlparser

import (
	"testing"
)

func testParse(t *testing.T, sql string) {
	stmt, err := Parse(sql)
	if err != nil {
		t.Fatalf("parse sql:%s,err:%v", sql, err)
	}

	_ = stmt
}

func TestSet(t *testing.T) {
	sql := "set names gbk"
	testParse(t, sql)
}

func TestSimpleSelect(t *testing.T) {
	sql := "select last_insert_id() as a"
	testParse(t, sql)
}

func TestMixer(t *testing.T) {

	sql := `select * from tb where id = $1`
	testParse(t, sql)

	sql = `select a || b from tb where id = $1`
	testParse(t, sql)

	sql = `insert into table1
			(col1,col2,col3)values($1,$2,$3)`
	testParse(t, sql)

	sql = `SHOW TRANSACTION ISOLATION LEVEL`
	testParse(t, sql)

	sql = `SHOW TIME ZONE`
	testParse(t, sql)

	sql = `SHOW SESSION AUTHORIZATION`
	testParse(t, sql)

	sql = `SHOW YOURNAME`
	testParse(t, sql)
}
