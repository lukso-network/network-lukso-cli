import { HttpClientTestingModule } from '@angular/common/http/testing';
import { ComponentFixture, TestBed } from '@angular/core/testing';

import { LaunchpadComponent } from './launchpad.component';

describe('LaunchpadComponent', () => {
  let component: LaunchpadComponent;
  let fixture: ComponentFixture<LaunchpadComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [LaunchpadComponent],
      imports: [HttpClientTestingModule],
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
