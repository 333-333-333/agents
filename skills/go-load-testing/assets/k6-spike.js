import http from 'k6/http';
import { check, sleep } from 'k6';

// Spike test: simulate sudden burst of traffic (e.g., marketing campaign push).
// Run: k6 run tests/load/notification/spike.js

export const options = {
  stages: [
    { duration: '30s', target: 5 },     // Normal traffic
    { duration: '10s', target: 500 },   // SPIKE!
    { duration: '1m', target: 500 },    // Hold spike
    { duration: '10s', target: 5 },     // Return to normal
    { duration: '1m', target: 5 },      // Recovery period
    { duration: '10s', target: 0 },     // Ramp down
  ],
  thresholds: {
    http_req_duration: ['p(95)<1000'],  // Allow up to 1s during spike
    http_req_failed: ['rate<0.15'],     // Allow 15% during spike
  },
};

const BASE_URL = __ENV.BASE_URL || 'http://localhost:8083';

export default function () {
  const payload = JSON.stringify({
    channel: 'push',
    recipient: `device-${__VU}-${__ITER}`,
    subject: 'Spike test',
    content: 'Simulating traffic burst from marketing push.',
  });

  const params = { headers: { 'Content-Type': 'application/json' } };
  const res = http.post(`${BASE_URL}/api/v1/notifications/send`, payload, params);

  check(res, {
    'responded (any status)': (r) => r.status > 0,
    'valid JSON': (r) => {
      try { r.json(); return true; } catch (e) { return false; }
    },
  });

  sleep(0.05);
}
