import http from 'k6/http';
import { check } from 'k6';

export let options = {
  vus: 20,
  duration: '20s',
};

export default function () {
  const payload = JSON.stringify({
    title: `Test Movie ${__VU}-${__ITER}`,
    year: 2025
  });

  const headers = { 'Content-Type': 'application/json' };

  let res = http.post('http://localhost:8080/movies', payload, { headers });

  check(res, {
    'status is 201': (r) => r.status === 201,
  });
}