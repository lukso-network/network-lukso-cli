import { Injectable } from '@angular/core';
import { Observable, of, timer } from 'rxjs';
import { HttpClient } from '@angular/common/http';
import { catchError, switchMap } from 'rxjs/operators';
import { DEFAULT_UPDATE_INTERVAL } from '../shared/config';

@Injectable({
  providedIn: 'root',
})
export class PandoraMetricsService {
  metrics$: Observable<any>;
  peersOverTime$: Observable<any>;

  constructor(private httpClient: HttpClient) {
    const timer$ = timer(0, DEFAULT_UPDATE_INTERVAL);

    this.metrics$ = this.setMetrics$(timer$);
    this.peersOverTime$ = this.setPeersOverTime$(timer$);
  }

  getMetrics$() {
    return this.metrics$;
  }

  getPeersOverTime$() {
    return this.peersOverTime$;
  }

  private setMetrics$(timer$: Observable<number>) {
    return timer$.pipe(
      switchMap(() => {
        return this.httpClient.get('/api/pandora/debug/metrics').pipe(
          catchError(() => {
            return of({});
          })
        );
      })
    );
  }

  private setPeersOverTime$(timer$: Observable<number>) {
    return timer$.pipe(
      switchMap(() => {
        return this.httpClient.get('/api/pandora/peers-over-time').pipe(
          catchError(() => {
            return of({});
          })
        );
      })
    );
  }
}
