# Control plane

This directory contains the control plane for FreePN.  The control
plane is responsible for authenticating clients, allocating tunnel IPs
and updating gateway nodes with new peers.  The current
implementation in `cmd/freepn-api/main.go` is a minimal stub that
exposes a health check endpoint.  It does **not** perform any
provisioning.

## Running

To run the API locally with Go 1.20 or newer:

```bash
cd cmd/freepn-api
go run .
```

Visit http://localhost:8080/health to verify that it works.

## Extending

In a full implementation you would:

1. Accept a client’s public key and optional authentication token via
   `POST /v1/sessions`.
2. Allocate an unused IPv4 (/32) and optionally IPv6 (/128) address.
3. Use the [wgctrl](https://pkg.go.dev/golang.zx2c4.com/wireguard/wgctrl)
   Go library to add the client as a peer on the chosen gateway’s
   WireGuard interface.  This involves calling `AddPeer()` with the
   client’s public key and allowed IPs.
4. Return a JSON response containing the gateway’s endpoint
   (`host:port`), DNS settings, server public key and allowed IPs.
5. Implement rate limiting and quotas.  Consider using Redis or
   Postgres to store sessions and enforce quotas.

Please consult the design documents in the `docs/` directory for
recommended algorithms, cipher suites and network policies.
