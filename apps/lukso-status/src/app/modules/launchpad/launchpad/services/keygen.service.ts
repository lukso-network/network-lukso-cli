import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root',
})
export class KeygenService {
  httpClient: HttpClient;
  constructor(httpClient: HttpClient) {
    this.httpClient = httpClient;
  }

  genereateKeys(password: string, network: string, amountOfValidators: string) {
    return this.httpClient.post('/api/launchpad/generate-keys', {
      password,
      network,
      amountOfValidators,
    });
  }
}
