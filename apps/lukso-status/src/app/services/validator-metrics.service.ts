import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable, timer, of } from 'rxjs';
import { switchMap, catchError } from 'rxjs/operators';

@Injectable({
  providedIn: 'root',
})
export class ValidatorMetricsService {
  metrics$: Observable<any>;
  constructor(private httpClient: HttpClient) {
    this.metrics$ = timer(0, 3000).pipe(
      switchMap(() => {
        return httpClient.get('/validator/metrics');
      }),
      catchError(() => {
        return of({});
      })
    );
  }

  getMetrics$() {
    return this.metrics$;
  }
}
