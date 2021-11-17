import { Injectable } from '@angular/core';
import { BehaviorSubject } from 'rxjs';
import { shareReplay } from 'rxjs/operators';
import { NETWORKS } from '../modules/launchpad/launchpad/helpers/create-keys';

@Injectable({
  providedIn: 'root',
})
export class DataService {
  network$: BehaviorSubject<NETWORKS>;
  constructor() {
    this.network$ = new BehaviorSubject<NETWORKS>(NETWORKS.L15_DEV);
  }

  getNetwork$() {
    return this.network$.pipe(shareReplay());
  }

  setNetwork(network: NETWORKS) {
    localStorage.setItem('network', network);
    this.network$.next(network);
  }
}
