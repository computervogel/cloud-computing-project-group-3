import http from 'k6/http';
import { check, sleep } from 'k6';

export const options = {
  vus: 500, // Don't forget to change this value for each round of testing,
  duration: '360s',
};

export default function () {
  const res = http.get('http://localhost:9090');
  check(res, { 'status was 200': (r) => r.status == 200 });
  sleep(10);
}