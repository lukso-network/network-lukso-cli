import { TestBed } from '@angular/core/testing';

import { KeygenService } from './keygen.service';

describe('KeygenService', () => {
  let service: KeygenService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(KeygenService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
