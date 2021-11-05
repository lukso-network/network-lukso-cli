import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ValidatorStatusComponent } from './validator-status.component';

describe('ValidatorStatusComponent', () => {
  let component: ValidatorStatusComponent;
  let fixture: ComponentFixture<ValidatorStatusComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ValidatorStatusComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(ValidatorStatusComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
