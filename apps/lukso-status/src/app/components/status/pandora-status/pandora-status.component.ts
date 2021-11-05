import { Component, ChangeDetectionStrategy } from '@angular/core';
import { BehaviorSubject, combineLatest, Observable, of } from 'rxjs';
import {
  catchError,
  debounceTime,
  distinctUntilChanged,
  filter,
} from 'rxjs/operators';
import { PandoraMetricsService } from '../../../services/pandora-metrics.service';
import { map } from 'rxjs/operators';

@Component({
  selector: 'lukso-pandora-status',
  templateUrl: './pandora-status.component.html',
  styleUrls: ['./pandora-status.component.scss'],
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class PandoraStatusComponent {
  metrics$: Observable<any>;
  filteredMetrics$: Observable<any>;
  searchTerm$ = new BehaviorSubject('');
  latestBlock = 0;
  constructor(pandorsMetrics: PandoraMetricsService) {
    const searchTerm$ = this.searchTerm$.pipe(
      filter((text) => text.length > 2),
      debounceTime(10),
      distinctUntilChanged()
    );
    this.metrics$ = pandorsMetrics.getMetrics$();
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
