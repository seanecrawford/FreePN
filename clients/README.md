# Clients

This directory contains client implementations for FreePN.  Clients
establish a tunnel to a gateway, implement a kill switch and provide
configuration options for DNS and split‑tunnelling.  At the moment it
only includes a very simple command‑line program in `cli/` that
demonstrates the intended workflow.

Future clients may include:

* **Desktop GUI apps** for Windows, macOS and Linux using native
  platform frameworks and incorporating kill‑switch logic.
* **Mobile apps** for Android and iOS built on `VpnService` and
  NetworkExtension, respectively.
* **SDKs** for embedding FreePN connectivity into other applications.

When adding a new client, please create a new subdirectory (e.g.
`windows/`, `macos/`) and include a README explaining how to build
and run it.  Consult the security baseline for platform‑specific kill
switch requirements.
