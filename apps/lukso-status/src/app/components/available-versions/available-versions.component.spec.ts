import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AvailableVersionsComponent } from './available-versions.component';

describe('AvailableVersionsComponent', () => {
  let component: AvailableVersionsComponent;
  let fixture: ComponentFixture<AvailableVersionsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ AvailableVersionsComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(AvailableVersionsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
