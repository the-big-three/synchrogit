package main

import (
	"fmt"
	"github.com/the-big-three/synchrogit/helper"
	"sync"
)

func main() {
	helper.CheckEnv("CONFIG");
	fmt.Println("Synchrogit is warming up. Sit down, watch the world burn")
	defer fmt.Println("Synchrogit is calling it a day, see you next time - copycat")
	synchroGitSettings, _ := helper.ParseConfig()
	var wg sync.WaitGroup
	wg.Add(len(synchroGitSettings.Syncs))
	for index, sync := range synchroGitSettings.Syncs {
		go helper.SyncRepo(index,sync, &wg)
	}
	wg.Wait()
}
