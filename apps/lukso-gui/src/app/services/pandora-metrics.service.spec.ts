import { TestBed } from '@angular/core/testing';
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { PandoraMetricsService } from './pandora-metrics.service';
import { RxState } from '@rx-angular/state';
import { GLOBAL_RX_STATE } from '../shared/rx-state';

describe('PandoraMetricsService', () => {
  let service: PandoraMetricsService;

  beforeEach(() => {
    TestBed.configureTestingModule({
      imports: [HttpClientTestingModule],
      providers: [{ provide: GLOBAL_RX_STATE, useClass: RxState }],
    });
    service = TestBed.inject(PandoraMetricsService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
