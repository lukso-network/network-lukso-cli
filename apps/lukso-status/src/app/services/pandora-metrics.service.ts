import { Injectable } from '@angular/core';
import { Observable, of, timer } from 'rxjs';
import { HttpClient } from '@angular/common/http';
import { catchError, switchMap } from 'rxjs/operators';

@Injectable({
  providedIn: 'root',
})
export class PandoraMetricsService {
  metrics$: Observable<any>;
  constructor(private httpClient: HttpClient) {
    this.metrics$ = timer(0, 3000).pipe(
      switchMap(() => {
        return httpClient.get('/pandora/debug/metrics').pipe(
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
}
