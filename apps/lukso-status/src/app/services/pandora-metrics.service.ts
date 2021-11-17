import { Injectable } from '@angular/core';
import { Observable, of, timer } from 'rxjs';
import { HttpClient } from '@angular/common/http';
import { catchError, switchMap } from 'rxjs/operators';

@Injectable({
  providedIn: 'root',
})
export class PandoraMetricsService {
  metrics$: Observable<any>;
  peersOverTime$: Observable<any>;
  constructor(private httpClient: HttpClient) {
    const timer$ = timer(0, 3000);
    this.metrics$ = timer$.pipe(
      switchMap(() => {
        return httpClient.get('/api/pandora/debug/metrics').pipe(
          catchError(() => {
            return of({});
          })
        );
      })
    );
    this.peersOverTime$ = timer$.pipe(
      switchMap(() => {
        return httpClient.get('/api/pandora/peers-over-time').pipe(
          catchError(() => {
            return of({});
          })
        );
      })
    );
  }

  getMetrics$() {
    return this.metrics$;
  }

  getPeersOverTime$() {
    return this.peersOverTime$;
  }
}
