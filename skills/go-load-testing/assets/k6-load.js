import http from 'k6/http';
import { check, sleep } from 'k6';

// Load test: validate under expected production traffic.
// Run: k6 run tests/load/notification/load.js

export const options = {
  stages: [
    { duration: '1m', target: 20 },   // Ramp up to 20 VUs
    { duration: '3m', target: 50 },   // Hold at 50 VUs (expected load)
    { duration: '1m', target: 0 },    // Ramp down
  ],
  thresholds: {
    http_req_duration: ['p(95)<200', 'p(99)<500'],
    http_req_failed: ['rate<0.01'],
    http_reqs: ['rate>50'],            // At least 50 req/s
  },
};

const BASE_URL = __ENV.BASE_URL || 'http://localhost:8083';
const CHANNELS = ['email', 'sms', 'push', 'whatsapp'];

export default function () {
  const channel = CHANNELS[Math.floor(Math.random() * CHANNELS.length)];

  const payload = JSON.stringify({
    channel: channel,
    recipient: `user-${__VU}-${__ITER}@test.com`,
    subject: 'Load test notification',
    content: `Testing ${channel} channel under normal load. VU=${__VU} ITER=${__ITER}`,
  });

  const params = { headers: { 'Content-Type': 'application/json' } };
  const res = http.post(`${BASE_URL}/api/v1/notifications/send`, payload, params);

  check(res, {
    'status 200': (r) => r.status === 200,
    'response has data': (r) => r.json().data !== undefined,
    'notification sent': (r) => r.json().data.status === 'sent',
    'correct channel': (r) => r.json().data.channel === channel,
  });

  sleep(0.5);
}
