 <!-- <tr>
        <td>
          <div class="control">
            <input
              #searchbox
              class="input"
              type="text"
              placeholder="Search Metrics"
              (keyup)="searchTerm$.next(searchbox.value)"
            />
          </div>
        </td>
        <td>
          <table
            *ngIf="this.filteredMetrics$ | async as filteredMetrics"
            class="
              table
              is-bordered is-striped is-narrow is-hoverable is-fullwidth
            "
          >
            <ng-container *ngIf="filteredMetrics.length > 0; else noMatches">
              <tr *ngFor="let metric of filteredMetrics">
                <td class="property-name">{{ metric[0] }}</td>
                <td>{{ metric[1] }}</td>
              </tr>
            </ng-container>
            <ng-template #noMatches>No Matches</ng-template>
          </table>
        </td>
      </tr> -->

        <!-- <tr>
        <td class="property-name">Current Justified Epoch</td>
        <td>{{ metrics['beacon_current_justified_epoch'] }}</td>
      </tr>
      <tr>
        <td class="property-name">Finalized Epoch</td>
        <td>{{ metrics['beacon_finalized_epoch'] }}</td>
      </tr>
      <tr>
        <td class="property-name">Reorg Total</td>
        <td>{{ metrics['beacon_reorg_total'] }}</td>x
      </tr> -->



      <!-- <tr>
      <td class="property-name">Startup Arguments</td>
      <td class="cmdLine">
        <details>
          <summary>Show command-line arguments</summary>
          <table>
            <tr *ngFor="let item of metrics['cmdline']">
              <td>{{ item.split('=')[0] }}</td>
              <td class="values">{{ item.split('=')[1] }}</td>
            </tr>
          </table>
        </details>
      </td>
    </tr> -->
      <!-- <tr>
      <td>
        <div class="control">
          <input
            #searchbox
            class="input"
            type="text"
            placeholder="Search Metrics"
            (keyup)="searchTerm$.next(searchbox.value)"
          />
        </div>
      </td>
      <td>
        <table
          *ngIf="this.filteredMetrics$ | async as filteredMetrics"
          class="
            table
            is-bordered is-striped is-narrow is-hoverable is-fullwidth
          "
        >
          <ng-container *ngIf="filteredMetrics.length > 0; else noMatches">
            <tr *ngFor="let metric of filteredMetrics">
              <td class="property-name">{{ metric[0] }}</td>
              <td>{{ metric[1] }}</td>
            </tr>
          </ng-container>
          <ng-template #noMatches>No Matches</ng-template>
        </table>
      </td>
    </tr> -->













    // now := time.Now()
    // sec := now.Unix()

    // file, _ := os.Create("/home/rryter/.lukso/downloads/logs/" + client + "/pandora" + fmt.Sprint(sec) + ".log")
    // cmnd.Stdout = file

// filteredMetrics$: Observable<any>;
  // searchTerm$ = new BehaviorSubject('');
// const searchTerm$ = this.searchTerm$.pipe(
    //   filter((text) => text.length > 2),
    //   debounceTime(10),
    //   distinctUntilChanged()
    // );
    // this.filteredMetrics$ = combineLatest([searchTerm$, this.metrics$]).pipe(
// map(([searchTerm, metrics]) => {
// return Object.keys(metrics)
// .filter((key) => key.includes(searchTerm))
// .reduce((cur, key) => {
// return Object.assign(cur, { [key]: metrics[key] });
// }, {});
// }),
// map((metrics) => {
// return Object.entries(metrics);
// }),
// catchError(() => {
// return of({});
// })
// );
