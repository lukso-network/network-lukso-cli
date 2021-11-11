import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable, timer, of } from 'rxjs';
import { switchMap, catchError, map } from 'rxjs/operators';

@Injectable({
  providedIn: 'root',
})
export class ValidatorMetricsService {
  metrics$: Observable<any>;
  constructor(private httpClient: HttpClient) {
    this.metrics$ = timer(0, 3000).pipe(
      switchMap(() => {
        return httpClient
          .get('/validator/metrics', {
            responseType: 'text',
          })
          .pipe(
            map((result: any) => {
              return result.split('\n');
            }),
            map((lines) => {
              return lines.filter((line: string) => {
                return line !== '' && !line.startsWith('#');
              });
            }),
            map((lines) => {
              return lines.reduce((acc: any, curr: any) => {
                const [key, value] = curr.split(' ');
                acc[key] = value;
                return acc;
              }, {});
            }),
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
