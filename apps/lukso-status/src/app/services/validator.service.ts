import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root',
})
export class ValidatorService {
  constructor(private httpClient: HttpClient) {}

  resetValidator(network: string) {
    return this.httpClient.post('/api/launchpad/reset-validator', {
      network,
    });
  }
}
