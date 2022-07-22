service "http" "web_proxy" {
  listen_addr = "127.0.0.1:8080"
  foo = "bar"
  process "main" {
    command = ["/usr/local/bin/awesome-app", "server"]
  }

  process "mgmt" {
    command = ["/usr/local/bin/awesome-app", "mgmt"]
  }
}