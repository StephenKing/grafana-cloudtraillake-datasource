import { DataSourcePlugin } from '@grafana/data';
import { DataSource } from './datasource';
import { CtlQuery, CtlDataSourceOptions } from './types';
import {QueryEditor} from "./QueryEditor";

export const plugin = new DataSourcePlugin<DataSource, CtlQuery, CtlDataSourceOptions>(DataSource)
    .setQueryEditor(QueryEditor);
