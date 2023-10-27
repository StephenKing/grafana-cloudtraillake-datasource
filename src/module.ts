import { DataSourcePlugin } from '@grafana/data';
import { DataSource } from './datasource';
import { CtlQuery, CtlDataSourceOptions } from './types';
import {QueryEditor} from "./QueryEditor";
import {ConfigEditor} from "./ConfigEditor";

export const plugin = new DataSourcePlugin<DataSource, CtlQuery, CtlDataSourceOptions>(DataSource)
    .setConfigEditor(ConfigEditor)
    .setQueryEditor(QueryEditor);
