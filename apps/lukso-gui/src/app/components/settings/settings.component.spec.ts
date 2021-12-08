import { HttpClientTestingModule } from '@angular/common/http/testing';
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { RxState } from '@rx-angular/state';
import { PushModule } from '@rx-angular/template';
import { GLOBAL_RX_STATE } from '../../shared/rx-state';

import { SettingsComponent } from './settings.component';

describe('InitialSetupComponent', () => {
  let component: SettingsComponent;
  let fixture: ComponentFixture<SettingsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [SettingsComponent],
      imports: [ReactiveFormsModule, HttpClientTestingModule, PushModule],
      providers: [{ provide: GLOBAL_RX_STATE, useClass: RxState }],
    }).compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(SettingsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
