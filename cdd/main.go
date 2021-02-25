package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"

	"github.com/gosuri/uiprogress"
	"github.com/gosuri/uiprogress/util/strutil"
	"github.com/herryg91/cdd/cdd/commands/gen"
	"github.com/herryg91/cdd/cdd/commands/install"

	"github.com/spf13/cobra"
)

var steps = []string{
	"downloading source",
	"installing deps",
	"compiling",
	"packaging",
	"seeding database",
	"deploying",
	"staring servers",
}

func test() {
	fmt.Println("apps: deployment started: app1, app2")
	uiprogress.Start()

	var wg sync.WaitGroup
	wg.Add(1)
	go deploy("app1", &wg)
	wg.Add(1)
	go deploy("app2", &wg)
	wg.Wait()

	fmt.Println("apps: successfully deployed: app1, app2")
}

func deploy(app string, wg *sync.WaitGroup) {
	defer wg.Done()
	bar := uiprogress.AddBar(len(steps)).AppendCompleted().PrependElapsed()
	bar.Width = 50

	// prepend the deploy step to the bar
	bar.PrependFunc(func(b *uiprogress.Bar) string {
		return strutil.Resize(app+": "+steps[b.Current()-1], 22)
	})

	rand.Seed(500)
	for bar.Incr() {
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
	}
}

func main() {
	rootCmd := &cobra.Command{Use: "cdd", Short: "cdd", Long: "cdd"}
	rootCmd.AddCommand(gen.NewGenCmd().Command)
	rootCmd.AddCommand(install.NewInstallCmd().Command)
	// execute
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
