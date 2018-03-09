# Redirector

Tiny server to aid in redirecting to canonical domains on Kubernetes.

For a request to path `/{scheme}/{domain}/{path*}`, a redirect to `{scheme}://{domain}/{path}` will be produced.
