# FreePN

FreePN is a **free, open‑source VPN** that aims to provide privacy‑preserving connectivity
with modern cryptography, graceful fallback options, and transparent
governance.  This repository contains a minimal starter skeleton with
documentation, configuration files and placeholder code to help you
bootstrap a fully‑fledged service.  It is designed to be indexed by
GitHub so that ChatGPT and other tooling can inspect it once you
publish it.

## Project goals

* **Security first** – use strong, modern encryption by default.  The
  primary protocol is WireGuard (UDP) with a MASQUE (HTTP/3/QUIC)
  fallback on port 443.  OpenVPN on TLS 1.3 and an IKEv2/IPsec profile
  are provided as last‑resort fallbacks.  These transports follow
  best‑practice cipher suites and parameter selections.
* **User privacy** – there are no per‑user activity logs.  A system
  kill switch blocks traffic outside the tunnel, and DNS queries are
  always sent over the VPN using encrypted DNS (DoT/DoH).
* **Openness and transparency** – all source code and configuration
  files are licensed under permissive or copyleft licences (see
  [LICENSES](LICENSES/)).  Documents in `docs/` describe the threat
  model, the security baseline and the high‑level architecture.  The
  goal is to make it easy for others to audit, reproduce and
  contribute to the project.

## Repository layout

```
freepn_repo/
├── .gitignore            – ignored files and directories
├── README.md             – this file
├── LICENSES/             – licence stubs for server and client code
│   ├── CLIENT-APACHE-2.0.txt
│   └── SERVER-AGPL-3.0.txt
├── docs/                 – design and security documentation
│   ├── architecture.md
│   ├── security-baseline.md
│   └── threat-model.md
├── control-plane/        – Go server that provisions WireGuard peers
│   ├── cmd/freepn-api/main.go
│   └── README.md
├── clients/              – client implementations or stubs
│   ├── cli/main.go
│   └── README.md
├── gateway/              – scripts and playbooks to bring up nodes
│   ├── ansible/playbook.yml
│   ├── ansible/roles/wireguard/templates/wg0.conf.j2
│   ├── ansible/roles/nftables/templates/freepn.nft.j2
│   └── README.md
└── .gitignore
```

The `control-plane/` directory contains a minimal Go program that
prints a stub response; you can replace it with a full REST API
implementation that provisions WireGuard peers via the `wgctrl` Go
package.  The `gateway/ansible` directory contains a basic Ansible
playbook and template files to bring up a Linux gateway with
WireGuard and nftables.  Client implementations live under
`clients/`; currently there is a command‑line stub that can be
extended or replaced with platform‑specific GUI applications.

## Getting started

1. **Install dependencies.**  For the control plane and CLI you need
   Go 1.20 or newer.  Gateway scripts assume a modern Linux host with
   Ansible installed.
2. **Run the gateway.**  Edit `gateway/ansible/inventory.ini` (not
   provided) to point at your host.  Run:
   ```bash
   cd gateway/ansible
   ansible-playbook -i inventory.ini playbook.yml
   ```
   This will install WireGuard, configure `wg0`, enable IP forwarding
   and deploy nftables rules.
3. **Start the control plane.**  In another terminal:
   ```bash
   cd control-plane/cmd/freepn-api
   go run .
   ```
   The stub program simply prints a message.  Replace it with a real
   API that accepts a public key, allocates an IP address and calls
   the WireGuard interface on the gateway.
4. **Generate a client configuration.**  The `clients/cli` program
   demonstrates how to generate a WireGuard keypair, request a
   session from the control plane and output a `.conf` file.  To run
   it:
   ```bash
   cd clients/cli
   go run . > my-freepn.conf
   ```
   Import the resulting file into the official WireGuard client to
   establish a tunnel.

## Contributing

Contributions are welcome!  Please see the documents in `docs/` to
understand the threat model and security requirements before
submitting changes.  All new features should include tests and
documentation.  Use pull requests to contribute and sign your
commits.
