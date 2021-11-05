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
