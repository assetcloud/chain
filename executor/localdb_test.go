package executor_test

import (
	"testing"

	dbm "github.com/assetcloud/chain/common/db"
	"github.com/assetcloud/chain/executor"
	"github.com/assetcloud/chain/types"
	"github.com/assetcloud/chain/util/testnode"
	"github.com/stretchr/testify/assert"
)

func TestLocalDBGet(t *testing.T) {
	mock33 := testnode.New("", nil)
	defer mock33.Close()
	db := executor.NewLocalDB(mock33.GetClient(), mock33.GetAPI(), false)
	defer db.(*executor.LocalDB).Close()
	testDBGet(t, db)
}

func TestLocalDBGetReadOnly(t *testing.T) {
	mock33 := testnode.New("", nil)
	defer mock33.Close()
	db := executor.NewLocalDB(mock33.GetClient(), mock33.GetAPI(), true)
	defer db.(*executor.LocalDB).Close()
	testDBGet(t, db)
}

func TestLocalDBEnable(t *testing.T) {
	mock33 := testnode.New("", nil)
	defer mock33.Close()
	db := executor.NewLocalDB(mock33.GetClient(), mock33.GetAPI(), false)
	ldb := db.(*executor.LocalDB)
	defer ldb.Close()
	_, err := ldb.Get([]byte("hello"))
	assert.Equal(t, err, types.ErrNotFound)
	ldb.DisableRead()
	_, err = ldb.Get([]byte("hello"))
	assert.Equal(t, err, types.ErrDisableRead)
	_, err = ldb.List(nil, nil, 0, 0)
	assert.Equal(t, err, types.ErrDisableRead)
	ldb.EnableRead()
	_, err = ldb.Get([]byte("hello"))
	assert.Equal(t, err, types.ErrNotFound)
	_, err = ldb.List(nil, nil, 0, 0)
	assert.Equal(t, err, nil)
	ldb.DisableWrite()
	err = ldb.Set([]byte("hello"), nil)
	assert.Equal(t, err, types.ErrDisableWrite)
	ldb.EnableWrite()
	err = ldb.Set([]byte("hello"), nil)
	assert.Equal(t, err, nil)
}

func BenchmarkLocalDBGet(b *testing.B) {
	mock33 := testnode.New("", nil)
	defer mock33.Close()
	db := executor.NewLocalDB(mock33.GetClient(), mock33.GetAPI(), false)
	defer db.(*executor.LocalDB).Close()

	err := db.Set([]byte("k1"), []byte("v1"))
	assert.Nil(b, err)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		v, err := db.Get([]byte("k1"))
		assert.Nil(b, err)
		assert.Equal(b, v, []byte("v1"))
	}
}

func TestLocalDBTxGet(t *testing.T) {
	mock33 := testnode.New("", nil)
	defer mock33.Close()
	db := executor.NewLocalDB(mock33.GetClient(), mock33.GetAPI(), false)
	testTxGet(t, db)
}

func testDBGet(t *testing.T, db dbm.KV) {
	err := db.Set([]byte("k1"), []byte("v1"))
	assert.Nil(t, err)
	v, err := db.Get([]byte("k1"))
	assert.Nil(t, err)
	assert.Equal(t, v, []byte("v1"))

	err = db.Set([]byte("k1"), []byte("v11"))
	assert.Nil(t, err)
	v, err = db.Get([]byte("k1"))
	assert.Nil(t, err)
	assert.Equal(t, v, []byte("v11"))
}

func testTxGet(t *testing.T, db dbm.KV) {
	//新版本
	db.Begin()
	err := db.Set([]byte("k1"), []byte("v1"))
	assert.Nil(t, err)
	v, err := db.Get([]byte("k1"))
	assert.Nil(t, err)
	assert.Equal(t, v, []byte("v1"))

	db.Commit()
	v, err = db.Get([]byte("k1"))
	assert.Nil(t, err)
	assert.Equal(t, v, []byte("v1"))

	//在非transaction中set，直接set成功，不能rollback
	err = db.Set([]byte("k1"), []byte("v11"))
	assert.Nil(t, err)

	db.Begin()
	v, err = db.Get([]byte("k1"))
	assert.Nil(t, err)
	assert.Equal(t, v, []byte("v11"))

	err = db.Set([]byte("k1"), []byte("v12"))
	assert.Nil(t, err)
	v, err = db.Get([]byte("k1"))
	assert.Nil(t, err)
	assert.Equal(t, v, []byte("v12"))

	db.Rollback()
	v, err = db.Get([]byte("k1"))
	assert.Nil(t, err)
	assert.Equal(t, v, []byte("v11"))
}

func TestLocalDBDel(t *testing.T) {
	mock33 := testnode.New("", nil)
	defer mock33.Close()
	db := executor.NewLocalDB(mock33.GetClient(), mock33.GetAPI(), false)
	//设置key
	db.Begin()
	err := db.Set([]byte("k1"), []byte("v1"))
	assert.Nil(t, err)
	db.Commit()

	//modify key
	db.Begin()
	err = db.Set([]byte("k1"), []byte("v11"))
	assert.Nil(t, err)
	err = db.Set([]byte("k2"), []byte("v2"))
	assert.Nil(t, err)

	//读取列表
	values, err := db.List([]byte("k"), nil, 0, 0)
	assert.Nil(t, err)
	assert.Equal(t, len(values), 2)
	assert.Equal(t, string(values[0]), "v2")
	assert.Equal(t, string(values[1]), "v11")
	db.Commit()

	//test key
	db.Begin()
	v, err := db.Get([]byte("k1"))
	assert.Nil(t, err)
	assert.Equal(t, v, []byte("v11"))
	db.Commit()

	//删除key
	db.Begin()
	err = db.Set([]byte("k1"), nil)
	assert.Nil(t, err)
	//在transaction 内部读取
	v, err = db.Get([]byte("k1"))
	assert.Equal(t, err, types.ErrNotFound)
	assert.Equal(t, v, []byte(nil))

	values, err = db.List([]byte("k"), nil, 0, 0)
	assert.Nil(t, err)
	assert.Equal(t, len(values), 1)
	assert.Equal(t, string(values[0]), "v2")
	//测试列表
	db.Commit()

	//在transaction 外部读取key
	db.Begin()
	v, err = db.Get([]byte("k1"))
	assert.Equal(t, err, types.ErrNotFound)
	assert.Equal(t, v, []byte(nil))
	values, err = db.List([]byte("k"), nil, 0, 0)
	assert.Nil(t, err)
	assert.Equal(t, len(values), 1)
	assert.Equal(t, string(values[0]), "v2")
	db.Commit()
}

func TestLocalDB(t *testing.T) {
	mock33 := testnode.New("", nil)
	defer mock33.Close()
	db := executor.NewLocalDB(mock33.GetClient(), mock33.GetAPI(), false)
	defer db.(*executor.LocalDB).Close()
	err := db.Set([]byte("k1"), []byte("v1"))
	assert.Nil(t, err)
	v, err := db.Get([]byte("k1"))
	assert.Nil(t, err)
	assert.Equal(t, v, []byte("v1"))

	err = db.Set([]byte("k1"), []byte("v11"))
	assert.Nil(t, err)
	v, err = db.Get([]byte("k1"))
	assert.Nil(t, err)
	assert.Equal(t, v, []byte("v11"))

	//beigin and rollback not imp
	db.Begin()
	err = db.Set([]byte("k2"), []byte("v2"))
	assert.Nil(t, err)
	db.Rollback()
	_, err = db.Get([]byte("k2"))
	assert.Equal(t, err, types.ErrNotFound)
	err = db.Set([]byte("k2"), []byte("v2"))
	assert.Nil(t, err)
	//get
	v, err = db.Get([]byte("k2"))
	assert.Nil(t, err)
	assert.Equal(t, v, []byte("v2"))
	//list
	values, err := db.List([]byte("k"), nil, 0, 0)
	assert.Nil(t, err)
	assert.Equal(t, len(values), 2)
	assert.Equal(t, string(values[0]), "v2")
	assert.Equal(t, string(values[1]), "v11")
	err = db.Commit()
	assert.Nil(t, err)
	//get
	v, err = db.Get([]byte("k2"))
	assert.Nil(t, err)
	assert.Equal(t, v, []byte("v2"))
	//list
	values, err = db.List([]byte("k"), nil, 0, 0)
	assert.Nil(t, err)
	assert.Equal(t, len(values), 2)
	assert.Equal(t, string(values[0]), "v2")
	assert.Equal(t, string(values[1]), "v11")
}
