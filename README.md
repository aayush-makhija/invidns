# invidns

`invidns` is a Caddy DNS provider module that sends requests to a specified URL with the provider's details. This module allows integration with external DNS services by sending necessary credentials and request data.

## Features

- **Simple Configuration**: Configure the module using the Caddyfile.
- **Secure Transmission**: Passwords are base64 encoded before being sent.
- **Timestamping**: Requests include a timestamp in Indian Standard Time (IST).

## Installation

1. **Clone the Repository**:
    ```sh
    git clone https://github.com/aayush-makhija/invidns.git
    cd invidns
    ```

2. **Build the Caddy Module**:
    Follow the [Caddy documentation](https://caddyserver.com/docs/extending-caddy) on how to build and use custom Caddy modules.

## Configuration

Configure the `invidns` module in your Caddyfile. Below is an example configuration:

```caddyfile
:port {
    tls {
        dns invidns {
            url your_url
            username your_username
            password your_password
        }
    }
}
