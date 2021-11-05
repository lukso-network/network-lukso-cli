import { Component, ChangeDetectionStrategy } from '@angular/core';
import {
  BehaviorSubject,
  combineLatest,
  fromEvent,
  Observable,
  of,
} from 'rxjs';
import { VanguardServiceService } from '../../../services/vanguard-metrics.service';
import {
  catchError,
  debounceTime,
  distinctUntilChanged,
  filter,
  map,
} from 'rxjs/operators';

@Component({
  selector: 'lukso-vanguard-status',
  templateUrl: './vanguard-status.component.html',
  styleUrls: ['./vanguard-status.component.scss'],
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class VanguardStatusComponent {
  metrics$: Observable<any>;
  filteredMetrics$: Observable<any>;
  searchTerm$ = new BehaviorSubject('');
  peersSelector = 'p2p_peer_count{state="Connected"}';
  constructor(vanguardMetrics: VanguardServiceService) {
    const searchTerm$ = this.searchTerm$.pipe(
      filter((text) => text.length > 2),
      debounceTime(10),
      distinctUntilChanged()
    );
    this.metrics$ = vanguardMetrics.getMetrics$();
    this.filteredMetrics$ = combineLatest([searchTerm$, this.metrics$]).pipe(
      map(([searchTerm, metrics]) => {
        return Object.keys(metrics)
          .filter((key) => key.includes(searchTerm))
          .reduce((cur, key) => {
            return Object.assign(cur, { [key]: metrics[key] });
          }, {});
      }),
      map((metrics) => {
        return Object.entries(metrics);
      }),
      catchError(() => {
        return of({});
      })
    );
  }
}
