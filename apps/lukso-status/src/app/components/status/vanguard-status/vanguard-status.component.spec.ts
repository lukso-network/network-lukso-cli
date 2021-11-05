import { ComponentFixture, TestBed } from '@angular/core/testing';

import { VanguardStatusComponent } from './vanguard-status.component';

describe('VanguardStatusComponent', () => {
  let component: VanguardStatusComponent;
  let fixture: ComponentFixture<VanguardStatusComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ VanguardStatusComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(VanguardStatusComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
