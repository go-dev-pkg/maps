package maps_test

import (
	"sync"
	"testing"

	"github.com/go-dev-pkg/maps"
)

func TestMap(t *testing.T) {
	m := maps.NewMap()
	wg := &sync.WaitGroup{}
	// 并发写入
	wg.Go(func() {
		_wg := &sync.WaitGroup{}
		for i := 0; i < 100; i++ {
			_wg.Go(func() {
				for j := 0; j < 100; j++ {
					__wg := &sync.WaitGroup{}
					for k := 0; k < 100; k++ {
						__wg.Go(func() {
							m.Store(i, map[interface{}]interface{}{k: j})
						})
					}
					__wg.Wait()
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
				m.Load(i)
			})
		}
		_wg.Wait()
	})
	// 并发循环
	wg.Go(func() {
		_wg := &sync.WaitGroup{}
		for i := 0; i < 100; i++ {
			_wg.Go(func() {
				m.Range(func(key interface{}, value map[interface{}]interface{}) bool {
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
				m.Len()
			})
		}
		_wg.Wait()
	})
	// 并发删除
	wg.Go(func() {
		_wg := &sync.WaitGroup{}
		for i := 0; i < 100; i++ {
			_wg.Go(func() {
				m.Delete(i)
			})
		}
		_wg.Wait()
	})
	// 并发清空
	wg.Go(func() {
		_wg := &sync.WaitGroup{}
		for i := 0; i < 100; i++ {
			_wg.Go(func() {
				m.Clear()
			})
		}
		_wg.Wait()
	})
	wg.Wait()
}
