import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { filter, map } from 'rxjs/operators';
import { DepositData } from '../helpers/create-keys';

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

  importKeys(walletPassword: string, network: string) {
    return this.httpClient.post(
      '/api/launchpad/import-keys',
      {
        walletPassword,
        network,
      },
      {
        responseType: 'blob',
      }
    );
  }

  getDepositData(network: string) {
    return this.httpClient.get<DepositData[]>('/api/deposit-data', {
      params: {
        network,
      },
    });
  }
}
