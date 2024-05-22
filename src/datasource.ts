import { DataQueryRequest, DataQueryResponse, DataSourceInstanceSettings, ScopedVars } from '@grafana/data';
import { getTemplateSrv, TemplateSrv } from '@grafana/runtime';
import { CtlDataSourceOptions, CtlQuery } from './types';
import { filterSQLQuery, applySQLTemplateVariables } from '@grafana/aws-sdk';
import { DatasourceWithAsyncBackend } from '@grafana/async-query-data';
import { Observable } from 'rxjs';
import { cloneDeep } from 'lodash';
import { annotationSupport } from './annotationSupport';

export class DataSource extends DatasourceWithAsyncBackend<CtlQuery, CtlDataSourceOptions> {
  defaultRegion = '';
  defaultEdsId = '';

  constructor(
      instanceSettings: DataSourceInstanceSettings<CtlDataSourceOptions>,
      private readonly templateSrv: TemplateSrv = getTemplateSrv()
  ) {
    super(instanceSettings);
    this.defaultRegion = instanceSettings.jsonData.defaultRegion || '';
    this.defaultEdsId = instanceSettings.jsonData.eventDataStore?.id || '';
  }

  annotations = annotationSupport;

  filterQuery(target: CtlQuery) {
    return target.hide !== true && filterSQLQuery(target);
  }

  applyTemplateVariables = (query: CtlQuery, scopedVars: ScopedVars) =>
      applySQLTemplateVariables(query, scopedVars, getTemplateSrv);

  getVariables = () => this.templateSrv.getVariables().map((v) => `$${v.name}`);

  getRegions = () => this.getResource('regions');

  getEventDataStores = () => this.getResource('eventDataStores');

  buildQuery(options: DataQueryRequest<CtlQuery>, queries: CtlQuery[]): CtlQuery[] {
    const updatedQueries = queries.map((query) => {
      query.connectionArgs.region = this.templateSrv.replace(query.connectionArgs.region, options.scopedVars);
      query.edsId = this.templateSrv.replace(query.edsId, options.scopedVars);
      console.log("buildQuery - query", query)
      console.log("buildQuery - scopedVars", options.scopedVars)
      console.log("buildQuery - edsId", query.edsId)
      return query;
    });

    return updatedQueries;
  }

  query(options: DataQueryRequest<CtlQuery>): Observable<DataQueryResponse> {
    options = cloneDeep(options);

    const queries = options.targets.filter((item) => item.hide !== true);

    options.targets = this.buildQuery(options, queries);

    return super.query(options);
  }
}
