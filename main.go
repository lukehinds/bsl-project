package main

import (
    "log"
    "os"

    "github.com/hashicorp/terraform-exec/tfexec"
)

func main() {
    workingDir := "./terraform-config"

    // Ensure the directory exists
    if err := os.MkdirAll(workingDir, 0755); err != nil {
        log.Fatalf("error creating directory: %s", err)
    }

    // Initialize a new Terraform instance
    tf, err := tfexec.NewTerraform(workingDir, "terraform")
    if err != nil {
        log.Fatalf("error running NewTerraform: %s", err)
    }

    // Initialize Terraform configuration
    err = tf.Init(context.Background(), tfexec.Upgrade(true))
    if err != nil {
        log.Fatalf("error running Init: %s", err)
    }

    log.Printf("Terraform has been initialized in %s", workingDir)
}

