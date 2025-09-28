package main

import "fmt"

// This is a placeholder command‑line tool for FreePN.  It should
// generate a WireGuard key pair, call the control plane to request a
// session and output a WireGuard configuration file (.conf).  For now
// it only prints a message.
func main() {
    fmt.Println("# FreePN CLI placeholder")
    fmt.Println("# TODO: generate keypair, call freepn-api, output wg config")
}
