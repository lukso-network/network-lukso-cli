import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CreateKeysComponent } from './create-keys.component';

describe('CreateKeysComponent', () => {
  let component: CreateKeysComponent;
  let fixture: ComponentFixture<CreateKeysComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ CreateKeysComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(CreateKeysComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
