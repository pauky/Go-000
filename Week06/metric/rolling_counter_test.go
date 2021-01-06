package metric

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRollingCounterAdd(t *testing.T) {
	size := 3
	bucketDuration := time.Second
	opts := RollingCounterOpts{
		Size:           size,
		BucketDuration: bucketDuration,
	}
	r := NewRollingCounter(opts)
	getCounts := func() []int64 {
		buckets := []int64{}
		r.Reduce(func(i Iterator) float64 {
			for i.Next() {
				bucket := i.Bucket()
				buckets = append(buckets, bucket.Count)
			}
			return 0.0
		})
		return buckets
	}
	assert.Equal(t, []int64{0, 0, 0}, getCounts())
	r.Add()
	assert.Equal(t, []int64{0, 0, 1}, getCounts())
	r.Add()
	r.Add()
	assert.Equal(t, []int64{0, 0, 3}, getCounts())
	time.Sleep(time.Second)
	r.Add()
	r.Add()
	assert.Equal(t, []int64{0, 3, 2}, getCounts())
}
