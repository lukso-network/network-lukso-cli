import { Component, ChangeDetectionStrategy } from '@angular/core';
import { BehaviorSubject, combineLatest, fromEvent, Observable } from 'rxjs';
import { VanguardServiceService } from '../../services/vanguard-service.service';
import {
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
      })
    );
  }

  calculatePeersStatus(metrics: any) {
    const numberOfPeers = metrics['p2p_peer_count{state="Connected"}'];
    switch (true) {
      case numberOfPeers >= 10:
        return {
          'has-background-success': true,
        };
      case numberOfPeers < 10 && numberOfPeers > 5:
        return {
          'has-background-warning': true,
        };
      case numberOfPeers <= 5:
        return {
          'has-background-danger': true,
        };

      default:
        return {};
    }
  }
}
