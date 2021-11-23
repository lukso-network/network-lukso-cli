import { Inject, Injectable } from '@angular/core';
import { merge, Observable, of, timer } from 'rxjs';
import { HttpClient } from '@angular/common/http';
import {
  catchError,
  map,
  retry,
  switchMap,
  withLatestFrom,
} from 'rxjs/operators';
import { DEFAULT_UPDATE_INTERVAL, getNamespacePrefix } from '../shared/config';

import { webSocket, WebSocketSubject } from 'rxjs/webSocket';
import { GlobalState, GLOBAL_RX_STATE } from '../shared/rx-state';
import { RxState } from '@rx-angular/state';
import { NETWORKS } from '../modules/launchpad/launchpad/helpers/create-keys';

interface MetricsState {
  network: NETWORKS;
}

@Injectable({
  providedIn: 'root',
})
export class PandoraMetricsService extends RxState<MetricsState> {
  readonly network$ = this.select('network');
  metrics$: Observable<any>;
  peersOverTime$: Observable<any>;

  constructor(
    @Inject(GLOBAL_RX_STATE) private globalState: RxState<GlobalState>,
    private httpClient: HttpClient
  ) {
    super();

    const timer$ = timer(0, DEFAULT_UPDATE_INTERVAL);

    this.metrics$ = this.setMetrics$(timer$);
    this.peersOverTime$ = this.setPeersOverTime$(timer$);
    // this.myWebSocket = this.network$.pipe(
    //   switchMap((network: NETWORKS) => {
    //     return webSocket(
    //       `wss://${getNamespacePrefix(network)}rpc.l15.lukso.network:8546`
    //     );
    //   })
    // );

    // const newHeads$ = this.myWebSocket.pipe(
    //   retry(),
    //   map((data: any) => {
    //     const lastBlock = new Date();
    //     lastBlock.setSeconds(lastBlock.getSeconds() - 1);
    //     return {
    //       blockNumber: data?.params?.result.number
    //         ? parseInt(data?.params?.result.number, 16)
    //         : undefined,
    //       timeStamp: lastBlock,
    //     };
    //   })
    // );
    // this.networkData$ = this.network$.pipe(
    //   switchMap((network: string) => {
    //     return merge([this.getLastBlock$(network), newHeads$]);
    //   })
    // );
    // this.myWebSocket.next({
    //   jsonrpc: '2.0',
    //   id: 1,
    //   method: 'eth_subscribe',
    //   params: ['newHeads'],
    // });
  }

  getMetrics$() {
    return this.metrics$;
  }

  getPeersOverTime$() {
    return this.peersOverTime$;
  }

  private getLastBlock$(network: string) {
    return this.httpClient
      .post(`https://${getNamespacePrefix(network)}rpc.l15.lukso.network`, {
        jsonrpc: '2.0',
        method: 'eth_blockNumber',
        params: [],
        id: 83,
      })
      .pipe(
        switchMap((blockNumberResponse: any) => {
          return this.httpClient
            .post(
              `https://${getNamespacePrefix(network)}rpc.l15.lukso.network`,
              {
                jsonrpc: '2.0',
                method: 'eth_getBlockByNumber',
                params: [blockNumberResponse.result, true],
                id: 83,
              }
            )
            .pipe(
              map((result: any) => {
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
