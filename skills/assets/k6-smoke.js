import http from 'k6/http';
import { check, sleep } from 'k6';

// Smoke test: verify service is alive and responds correctly under minimal load.
// Run: k6 run tests/load/notification/smoke.js

export const options = {
  vus: 1,
  duration: '30s',
  thresholds: {
    http_req_duration: ['p(95)<200'],  // 95% under 200ms
    http_req_failed: ['rate<0.01'],    // <1% errors
  },
};

const BASE_URL = __ENV.BASE_URL || 'http://localhost:8083';

export default function () {
  // Health check
  const healthRes = http.get(`${BASE_URL}/health`);
  check(healthRes, {
    'health: status 200': (r) => r.status === 200,
    'health: body ok': (r) => r.json().status === 'ok',
  });

  // Send notification
  const payload = JSON.stringify({
    channel: 'email',
    recipient: `user-${__VU}-${__ITER}@test.com`,
    subject: 'Smoke test',
    content: 'Testing notification service under minimal load.',
  });

  const params = { headers: { 'Content-Type': 'application/json' } };
  const sendRes = http.post(`${BASE_URL}/api/v1/notifications/send`, payload, params);

  check(sendRes, {
    'send: status 200': (r) => r.status === 200,
    'send: has data': (r) => r.json().data !== undefined,
    'send: status sent': (r) => r.json().data.status === 'sent',
    'send: has ID': (r) => r.json().data.id.length > 0,
  });

  sleep(1);
}
