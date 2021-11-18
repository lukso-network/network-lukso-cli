import { Injectable } from '@angular/core';
import { merge, Observable, of, timer } from 'rxjs';
import { HttpClient } from '@angular/common/http';
import { catchError, map, switchMap, tap } from 'rxjs/operators';
import { DEFAULT_UPDATE_INTERVAL } from '../shared/config';

import { webSocket, WebSocketSubject } from 'rxjs/webSocket';

@Injectable({
  providedIn: 'root',
})
export class PandoraMetricsService {
  metrics$: Observable<any>;
  peersOverTime$: Observable<any>;
  myWSData$: Observable<any>;
  myWebSocket: WebSocketSubject<any>;

  constructor(private httpClient: HttpClient) {
    const timer$ = timer(0, DEFAULT_UPDATE_INTERVAL);

    this.metrics$ = this.setMetrics$(timer$);
    this.peersOverTime$ = this.setPeersOverTime$(timer$);
    this.myWebSocket = webSocket('ws://localhost:8546');
    const a = this.myWebSocket.pipe(
      map((data) => {
        return {
          blockNumber: data?.params?.result.number
            ? parseInt(data?.params?.result.number, 16)
            : undefined,
          timeStamp: data?.params?.result.timestamp
            ? parseInt(data?.params?.result.timestamp, 16) * 1000 + 6000
            : undefined,
        };
      })
    );
    this.myWSData$ = merge(this.setLastBlock$(), a);
    this.myWebSocket.next({
      jsonrpc: '2.0',
      id: 1,
      method: 'eth_subscribe',
      params: ['newHeads'],
    });
  }

  getMetrics$() {
    return this.metrics$;
  }

  getPeersOverTime$() {
    return this.peersOverTime$;
  }

  private setLastBlock$() {
    return this.httpClient
      .post('https://staging.rpc.l15.lukso.network', {
        jsonrpc: '2.0',
        method: 'eth_blockNumber',
        params: [],
        id: 83,
      })
      .pipe(
        switchMap((blockNumberResponse: any) => {
          return this.httpClient
            .post('https://staging.rpc.l15.lukso.network', {
              jsonrpc: '2.0',
              method: 'eth_getBlockByNumber',
              params: [blockNumberResponse.result, true],
              id: 83,
            })
            .pipe(
              map((result: any) => {
                console.log(Date.now());
                console.log(parseInt(result.result.timestamp, 16) * 1000);
                return {
                  blockNumber: blockNumberResponse.result,
                  timeStamp:
                    parseInt(result.result.timestamp, 16) * 1000 + 5000,
                };
              })
            );
        })
      );
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
