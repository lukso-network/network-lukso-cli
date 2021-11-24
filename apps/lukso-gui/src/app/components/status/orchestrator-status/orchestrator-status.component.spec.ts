import { ComponentFixture, TestBed } from '@angular/core/testing';

import { OrchestratorStatusComponent } from './orchestrator-status.component';

describe('OrchestratorStatusComponent', () => {
  let component: OrchestratorStatusComponent;
  let fixture: ComponentFixture<OrchestratorStatusComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ OrchestratorStatusComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(OrchestratorStatusComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
