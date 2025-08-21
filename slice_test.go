package maps_test

import (
	"sync"
	"testing"

	"github.com/go-dev-pkg/maps"
)

func TestSlice(t *testing.T) {
	s := maps.NewSlice()
	wg := &sync.WaitGroup{}
	// 并发写入
	wg.Go(func() {
		_wg := &sync.WaitGroup{}
		for i := 0; i < 1000; i++ {
			_wg.Go(func() {
				for j := 0; j < 1000; j++ {
					s.Store(i, j, j+1, j+2)
				}
			})
		}
		_wg.Wait()
	})
	// 并发读取
	wg.Go(func() {
		_wg := &sync.WaitGroup{}
		for i := 0; i < 100; i++ {
			_wg.Go(func() {
				s.Load(i)
			})
		}
		_wg.Wait()
	})
	// 并发循环
	wg.Go(func() {
		_wg := &sync.WaitGroup{}
		for i := 0; i < 100; i++ {
			_wg.Go(func() {
				s.Range(func(key interface{}, value []interface{}) bool {
					return true
				})
			})
		}
		_wg.Wait()
	})
	// 并发获取长度
	wg.Go(func() {
		_wg := &sync.WaitGroup{}
		for i := 0; i < 100; i++ {
			_wg.Go(func() {
				s.Len()
			})
		}
		_wg.Wait()
	})
	// 并发删除
	wg.Go(func() {
		_wg := &sync.WaitGroup{}
		for i := 0; i < 100; i++ {
			_wg.Go(func() {
				s.Delete(i)
			})
		}
		_wg.Wait()
	})
	// 并发清空
	wg.Go(func() {
		_wg := &sync.WaitGroup{}
		for i := 0; i < 100; i++ {
			_wg.Go(func() {
				s.Clear()
			})
		}
		_wg.Wait()
	})
	wg.Wait()
}
