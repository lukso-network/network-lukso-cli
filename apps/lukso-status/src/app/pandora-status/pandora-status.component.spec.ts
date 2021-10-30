import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PandoraStatusComponent } from './pandora-status.component';

describe('PandoraStatusComponent', () => {
  let component: PandoraStatusComponent;
  let fixture: ComponentFixture<PandoraStatusComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ PandoraStatusComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(PandoraStatusComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
