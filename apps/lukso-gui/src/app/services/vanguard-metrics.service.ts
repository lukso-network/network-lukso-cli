import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable, of, timer } from 'rxjs';
import { catchError, map, switchMap, tap } from 'rxjs/operators';
import { environment } from '../../environments/environment';

@Injectable({
  providedIn: 'root',
})
export class VanguardService {
  metrics$: Observable<any>;
  peersOverTime$: Observable<any>;
  constructor(private httpClient: HttpClient) {
    const timer$ = timer(0, 3000);
    this.metrics$ = timer$.pipe(
      switchMap(() => {
        return httpClient.get(environment.API + '/vanguard/metrics').pipe(
          catchError(() => {
            return of({});
          })
        );
      })
    );
    this.peersOverTime$ = this.setPeersOverTime$(timer$);
  }

  getMetrics$() {
    return this.metrics$;
  }

  getPeersOverTime$() {
    return this.peersOverTime$;
  }

  private setPeersOverTime$(timer$: Observable<number>) {
    return timer$.pipe(
      switchMap(() => {
        return this.httpClient
          .get(environment.API + '/vanguard/peers-over-time')
          .pipe(
            catchError(() => {
              return of({});
            })
          );
      })
    );
  }
}
