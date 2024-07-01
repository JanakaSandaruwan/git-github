openapi: 3.0.1
info:
  title: /
  version: 0.1.0
servers:
- url: "{server}:{port}/"
  variables:
    server:
      default: http://localhost
    port:
      default: "8090"
paths:
  /:
    post:
      operationId: post
      requestBody:
        content:
          text/plain:
            schema:
              type: string
      responses:
        "201":
          description: Created
          content:
            text/plain:
              schema:
                type: string
components: {}
