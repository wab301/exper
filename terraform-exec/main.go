package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/hashicorp/go-version"
	"github.com/hashicorp/hc-install/product"
	"github.com/hashicorp/hc-install/releases"
	"github.com/hashicorp/terraform-exec/tfexec"
)

var (
	workDir = flag.String("work_dir", "", "terrafrom working dir")
)

func main() {
	flag.Parse()
	installer := &releases.ExactVersion{
		Product: product.Terraform,
		Version: version.Must(version.NewVersion("1.0.6")),
	}
	execPath := "/usr/local/bin/terraform"
	_, err := os.Stat(execPath)
	if os.IsNotExist(err) {
		log.Println("====install====")
		execPath, err = installer.Install(context.Background())
		if err != nil {
			log.Fatalf("error installing Terraform: %s", err)
		}
	}

	log.Println("====new====")
	tf, err := tfexec.NewTerraform(*workDir, execPath)
	if err != nil {
		log.Fatalf("error running NewTerraform: %s", err)
	}
	tf.SetStdout(os.Stdout)
	tf.SetStderr(os.Stderr)

	log.Println("====init====")
	err = tf.Init(context.Background(), tfexec.Upgrade(true))
	if err != nil {
		log.Fatalf("error running Init: %s", err)
	}
	log.Println("====show====")
	state, err := tf.Show(context.Background())
	if err != nil {
		log.Fatalf("error running Show: %s", err)
	}

	fmt.Println(state.FormatVersion) // "0.2"
}
