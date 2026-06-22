package repository

import "github.com/lib/pq"

func pqInt64Array(ids []int64) interface{} {
	return pq.Array(ids)
}
