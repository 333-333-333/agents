---
name: e2e-bruno
description: >
  End-to-end API testing with Bruno HTTP client: collection organization, environment variables, test scripts, and CI integration.
  Trigger: When creating E2E API tests, organizing Bruno collections, or setting up API test automation.
license: Apache-2.0
metadata:
  author: 333-333-333
  version: "1.0"
  type: project
  scope: [api]
  auto_invoke:
    - "Creating E2E API tests with Bruno"
    - "Organizing Bruno collections for services"
    - "Setting up API test automation"
    - "Writing Bruno test scripts for endpoints"
    - "Configuring Bruno environments for local/staging"
---

## When to Use

- Creating end-to-end tests for API endpoints
- Organizing HTTP request collections per service
- Setting up pre-request and post-response test scripts
- Configuring environment variables for local, dev, staging
- Automating API tests in CI/CD pipelines with Bruno CLI

---

## Critical Patterns

### Why Bruno over Postman

Bruno is open-source, Git-friendly (plain text files), and runs collections from CLI. No cloud sync needed — collections live in the repo under version control.

### Collection Organization

Collections follow the service structure. One directory per service, one file per endpoint:

```
tests/e2e/
  bruno.json                     # Bruno workspace config
  environments/
    local.bru                    # localhost URLs, test tokens
    development.bru              # dev environment URLs
    staging.bru                  # staging environment URLs
  notification/
    health.bru                   # GET /health
    send-email.bru               # POST /api/v1/notifications/send (email)
    send-sms.bru                 # POST /api/v1/notifications/send (sms)
    send-push.bru                # POST /api/v1/notifications/send (push)
    send-invalid-channel.bru     # POST with invalid channel (expect 400)
    send-missing-content.bru     # POST with missing content (expect 400)
  auth/
    register.bru
    login.bru
    logout.bru
  gateway/
    health.bru
    proxy-auth.bru
    proxy-notifications.bru
```

### Bru File Format

Every `.bru` file follows this structure:

> See [assets/example-success.bru](assets/example-success.bru) for a success path test.

> See [assets/example-error.bru](assets/example-error.bru) for an error path test.

> See [assets/example-chained.bru](assets/example-chained.bru) for a chained request that uses a previous response.

### Environment Files

Environments define variables per deployment target. Each `.bru` uses `{{variable}}` placeholders.

> See [assets/environment-local.bru](assets/environment-local.bru) for local environment config.

### Test Scripts

Bruno supports JavaScript test scripts in `tests` blocks. Key patterns:

| Pattern | What it verifies |
|---------|-----------------|
| Status code | `res.getStatus() === 200` |
| Response body field | `res.getBody().data.id !== undefined` |
| Response body type | `typeof res.getBody().data.status === 'string'` |
| Error envelope | `res.getBody().error.code === 'VALIDATION_ERROR'` |
| Set variable | `bru.setVar('notificationId', res.getBody().data.id)` |
| Chain requests | Use `{{notificationId}}` in subsequent requests |

### CI Integration

Bruno CLI runs collections headlessly. The pattern for CI:

> See [assets/ci-bruno.yml](assets/ci-bruno.yml) for GitHub Actions integration.

---

## Decision Tree

```
Testing a single endpoint happy path?
  → One .bru file with status + body assertions

Testing an error path (400, 422, 401)?
  → Separate .bru file named with error scenario

Testing a flow (register → login → use token)?
  → Chained .bru files with bru.setVar/getVar

Need to run in CI?
  → Use bru run with --env flag
```

---

## Assets

| File | Description |
|------|-------------|
| `assets/example-success.bru` | Success path test with body assertions |
| `assets/example-error.bru` | Error path test with error envelope validation |
| `assets/example-chained.bru` | Chained request using variables from previous response |
| `assets/environment-local.bru` | Local environment with localhost URLs |
| `assets/ci-bruno.yml` | GitHub Actions step for running Bruno in CI |
| `assets/bruno.json` | Bruno workspace config file |

---

## Commands

```bash
npm install -g @usebruno/cli   # Install Bruno CLI
bru run tests/e2e/notification --env local  # Run notification E2E tests
bru run tests/e2e --env local               # Run ALL E2E tests
```

---

## Resources

- **Templates**: See [assets/](assets/) for .bru file templates and CI config
