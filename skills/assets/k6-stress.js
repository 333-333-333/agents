import http from 'k6/http';
import { check, sleep } from 'k6';

// Stress test: find the breaking point by gradually increasing load.
// Run: k6 run tests/load/notification/stress.js

export const options = {
  stages: [
    { duration: '2m', target: 50 },    // Normal load
    { duration: '2m', target: 100 },   // Above normal
    { duration: '2m', target: 200 },   // Heavy load
    { duration: '2m', target: 350 },   // Stress
    { duration: '2m', target: 500 },   // Breaking point?
    { duration: '2m', target: 0 },     // Recovery
  ],
  thresholds: {
    http_req_duration: ['p(95)<500'],   // Relaxed for stress
    http_req_failed: ['rate<0.10'],     // Allow up to 10% errors under extreme load
  },
};

const BASE_URL = __ENV.BASE_URL || 'http://localhost:8083';

export default function () {
  const payload = JSON.stringify({
    channel: 'email',
    recipient: `stress-${__VU}-${__ITER}@test.com`,
    subject: 'Stress test',
    content: 'Finding the breaking point.',
  });

  const params = { headers: { 'Content-Type': 'application/json' } };
  const res = http.post(`${BASE_URL}/api/v1/notifications/send`, payload, params);

  check(res, {
    'status 200 or 503': (r) => r.status === 200 || r.status === 503,
    'valid response': (r) => {
      try { r.json(); return true; } catch (e) { return false; }
    },
  });

  sleep(0.1);
}
