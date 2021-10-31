import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable, timer } from 'rxjs';
import { map, switchMap, tap } from 'rxjs/operators';

@Injectable({
  providedIn: 'root',
})
export class VanguardServiceService {
  metrics$: Observable<any>;
  constructor(private httpClient: HttpClient) {
    this.metrics$ = timer(0, 3000).pipe(
      switchMap(() => {
        return httpClient
          .get('/vanguard/metrics', {
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
            })
          );
      })
    );
  }

  getMetrics$() {
    return this.metrics$;
  }
}
