import { Component, OnInit, ChangeDetectionStrategy } from '@angular/core';

@Component({
  selector: 'lukso-orchestrator-status',
  templateUrl: './orchestrator-status.component.html',
  styleUrls: ['./orchestrator-status.component.scss'],
  changeDetection: ChangeDetectionStrategy.OnPush
})
export class OrchestratorStatusComponent implements OnInit {

  constructor() { }

  ngOnInit(): void {
  }

}
