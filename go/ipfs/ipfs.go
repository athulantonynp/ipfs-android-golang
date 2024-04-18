package ipfs

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/ipfs/kubo/config"
	core "github.com/ipfs/kubo/core"
	"github.com/ipfs/kubo/plugin/loader"
	"github.com/ipfs/kubo/repo/fsrepo"
)

func Hello() string {
	return "Hello, World!"
}

func InitIPFS() {
	// Create a context with cancellation capability
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    // Set up the IPFS repository
    repoPath := "/data/user/0/com.dreamcatcher.android/files/ipfs" // Set your desired repo path
    err := os.MkdirAll(repoPath, 0755)
    if err != nil {
        fmt.Printf("Failed to create repo directory: %s\n", err)
        return
    }

	repoConfig, err := config.Init(io.Discard, 2048)
	if err != nil {
		fmt.Printf("Failed to initialize repo config: %s\n", err)
		return
	}
	// Add plugins to the config
	plugins, err := loader.NewPluginLoader(repoPath)

	if err != nil {
		panic(fmt.Errorf("error loading plugins: %s", err))
	}

	if err := plugins.Initialize(); err != nil {
		panic(fmt.Errorf("error initializing plugins: %s", err))
	}

	if err := plugins.Inject(); err != nil {
		panic(fmt.Errorf("error initializing plugins: %s", err))
	}
	
	fsrepo.Init(repoPath, repoConfig)

    // Initialize the IPFS repo
    r, err := fsrepo.Open(repoPath)
    if err != nil {
        fmt.Printf("Failed to open repo: %s\n", err)
        return
    }
    defer r.Close()

    // Construct the IPFS node
    cfg := &core.BuildCfg{
        Repo:   r,
        Online: true,
    }

    node, err := core.NewNode(ctx, cfg)
    if err != nil {
        fmt.Printf("Failed to create IPFS node: %s\n", err)
        return
    }
	fmt.Println(node)

    // Start the IPFS node
    // if err := node.Start(ctx); err != nil {
    //     fmt.Printf("Failed to start IPFS node: %s\n", err)
    //     return
    // }
    // defer node.Close()

    fmt.Println("IPFS node is created...")

}

func main() {
	fmt.Println(Hello())
}