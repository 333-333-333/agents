---
name: go-load-testing
description: >
  Load testing with k6 for Go microservices: smoke, load, stress, spike, and soak tests with pass/fail thresholds.
  Trigger: When writing load tests, running stress tests, validating performance thresholds, or benchmarking endpoints.
metadata:
  author: 333-333-333
  version: "1.0"
  type: generic
  scope: [api]
  auto_invoke:
    - "Writing load tests with k6"
    - "Running stress or spike tests"
    - "Validating service performance under load"
    - "Benchmarking API endpoint throughput"
---

## When to Use

- Validating service performance under expected and peak load
- Finding the breaking point of an endpoint
- Testing sudden traffic spikes (e.g., promotional push notifications)
- Detecting memory leaks or connection pool exhaustion (soak tests)
- Running performance smoke tests before releases

---

## Critical Patterns

### Test Types

| Test Type | Purpose | VUs | Duration |
|-----------|---------|-----|----------|
| **Smoke** | Verify service works under minimal load | 1-5 | 30s |
| **Load** | Validate under expected production traffic | 50-100 | 5m |
| **Stress** | Find breaking point | 100→500 ramp | 10m |
| **Spike** | Test sudden traffic burst | 1→500→1 | 3m |
| **Soak** | Find memory leaks, connection pool exhaustion | 50 | 30m+ |

### Directory Structure

k6 scripts live under `tests/load/{service}/`:

```
tests/load/
  notification/
    smoke.js           # Quick smoke test (1-5 VUs)
    load.js            # Normal load test (50-100 VUs)
    stress.js          # Stress test (ramp to 500 VUs)
    spike.js           # Spike test (sudden burst)
  auth/
    smoke.js
    load.js
```

### Thresholds

Every k6 test MUST define thresholds — the pass/fail criteria:

| Metric | Target | Meaning |
|--------|--------|---------|
| `http_req_duration p(95)` | < 200ms | 95% of requests under 200ms |
| `http_req_duration p(99)` | < 500ms | 99% of requests under 500ms |
| `http_req_failed` | < 1% | Less than 1% error rate |
| `http_reqs` | > 100/s | Minimum throughput |

---

## Decision Tree

```
Quick sanity check before deploy?
  → Smoke test (1-5 VUs, 30s)

Validating expected production traffic?
  → Load test (50-100 VUs, 5 min)

Finding the breaking point?
  → Stress test (ramp to 500 VUs)

Testing sudden traffic spike (e.g., push campaign)?
  → Spike test (1 → 500 → 1)

Detecting memory leaks or pool exhaustion?
  → Soak test (50 VUs, 30+ min)
```

---

## Assets

| File | Description |
|------|-------------|
| `assets/k6-smoke.js` | Smoke test: minimal load sanity check |
| `assets/k6-load.js` | Load test: expected production traffic with ramp stages |
| `assets/k6-stress.js` | Stress test: find breaking point with increasing load |
| `assets/k6-spike.js` | Spike test: sudden burst of traffic |

> See each asset for the full script template with thresholds.

---

## Commands

```bash
brew install k6                                             # Install k6
k6 run tests/load/notification/smoke.js                     # Run smoke test
k6 run tests/load/notification/load.js                      # Run load test
k6 run --out json=results.json tests/load/notification/stress.js  # With JSON output
```

---

## Anti-Patterns

| Don't | Do |
|----------|-------|
| Load test without thresholds | Always define p95, p99, error rate |
| Skip smoke test before deploy | Smoke is 30 seconds — always run it |
| Run stress tests against production | Run against staging or a dedicated environment |
| Ignore soak test results | Memory growth means a leak — fix it |
| Hardcode URLs in scripts | Use k6 environment variables for base URL |
