import { HttpClientTestingModule } from '@angular/common/http/testing';
import { NO_ERRORS_SCHEMA } from '@angular/core';
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { RxState } from '@rx-angular/state';
import { PushModule } from '@rx-angular/template';
import { GLOBAL_RX_STATE } from '../../../../../shared/rx-state';

import { LaunchpadComponent } from './launchpad.component';
describe('LaunchpadComponent', () => {
  let component: LaunchpadComponent;
  let fixture: ComponentFixture<LaunchpadComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      schemas: [NO_ERRORS_SCHEMA],
      declarations: [LaunchpadComponent],
      imports: [HttpClientTestingModule, PushModule],
      providers: [{ provide: GLOBAL_RX_STATE, useClass: RxState }],
    }).compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(LaunchpadComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
