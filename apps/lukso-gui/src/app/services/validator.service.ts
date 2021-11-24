import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { environment } from '../../environments/environment';

@Injectable({
  providedIn: 'root',
})
export class ValidatorService {
  constructor(private httpClient: HttpClient) {}

  resetValidator(network: string) {
    return this.httpClient.post(
      environment.API + '/launchpad/reset-validator',
      {
        network,
      }
    );
  }
}
