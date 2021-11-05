import { TestBed } from '@angular/core/testing';

import { AvailableVersionsService } from './available-versions.service';

describe('AvailableVersionsService', () => {
  let service: AvailableVersionsService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(AvailableVersionsService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
