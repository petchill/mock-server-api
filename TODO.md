# Mock Server

## Phase 1

- save response
  - input
    - path (exactly match)
    - status code
    - response body json
- auth
  - from db mongo
- project
  - global project (everyone use same user)
- error handler
  - custom error

## Phase 2

- save reponse
  - input
    - path (exactly match)
    - status code
    - response body json
    - response time
    - header
    - have root and leave node
    - param of path
- auth
  - google login
- project
  - individual per user
  