import {DataSourceSettings, SelectableValue} from '@grafana/data';
import {AwsAuthDataSourceJsonData, AwsAuthDataSourceSecureJsonData, SQLQuery} from '@grafana/aws-sdk';

export enum FormatOptions {
  TimeSeries,
  Table,
  Logs,
}

export const SelectableFormatOptions: Array<SelectableValue<FormatOptions>> = [
  {
    label: 'Time Series',
    value: FormatOptions.TimeSeries,
  },
  {
    label: 'Table',
    value: FormatOptions.Table,
  },
  {
    label: 'Logs',
    value: FormatOptions.Logs,
  },
];

export interface CtlQuery extends SQLQuery {
  format: FormatOptions;
  connectionArgs: {
    region?: string;
  };

  edsId?: string;
  queryID?: string;
}
export const defaultKey = '__default';
export const defaultQuery: Partial<CtlQuery> = {
  format: FormatOptions.Table,
  rawSQL: '',
  connectionArgs: {
    region: defaultKey,
  },
  edsId: defaultKey,
};
/**
 * These are options configured for each DataSource instance
 */
export interface CtlDataSourceOptions extends AwsAuthDataSourceJsonData {
  eventDataStore?: {
    name: string;
    id: string;
  }
}

export interface CtlDataSourceSecureJsonData extends AwsAuthDataSourceSecureJsonData {}

export type CtlDataSourceSettings = DataSourceSettings<CtlDataSourceOptions, CtlDataSourceSecureJsonData>;
