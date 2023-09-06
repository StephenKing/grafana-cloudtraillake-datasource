import React from 'react';
import { CtlQuery, CtlDataSourceOptions } from './types';
import { QueryEditorProps } from '@grafana/data';
import { DataSource } from 'datasource';
import { QueryEditorForm } from './QueryEditorForm';

export function AnnotationQueryEditor(props: QueryEditorProps<DataSource, CtlQuery, CtlDataSourceOptions>) {
  return <QueryEditorForm {...props} hideOptions />;
}
