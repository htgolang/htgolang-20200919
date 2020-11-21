package main

import "sync"

func main() {
	const N = 3

	names := listDir()
	// 工头 摊牌任务
	// 管道 taskPool

	wg := &sync.WaitGroup{}
	// worker 工人
	for i := 0; i < N; i++ {
		wg.Add(1)
		go func() {
			for {
				// 复制文件
				task, ok <- taskPool
				if !ok {
					break
				}
				copyFile(task[src], task[dst])
			}
			wg.Done()
		}()
	}

	for name := range {
		taskPool <- []src, dst
	}
	close(taskPool)
	wg.Wait()


}
