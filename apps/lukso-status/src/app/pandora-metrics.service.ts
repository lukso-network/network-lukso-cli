import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { HttpClient } from '@angular/common/http';

@Injectable({
  providedIn: 'root',
})
export class PandoraMetricsService {
  metrics$: Observable<any>;
  constructor(private httpClient: HttpClient) {
    this.metrics$ = httpClient.get(
      'http://localhost:4200/pandora/debug/metrics'
    );
  }

  getMetrics$() {
    return this.metrics$;
  }
}
