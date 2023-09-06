import React, { useEffect, useState } from 'react';
import { QueryEditorProps, SelectableValue } from '@grafana/data';
import { DataSource } from './datasource';
import { InlineSegmentGroup } from '@grafana/ui';
import { ResourceSelector } from '@grafana/aws-sdk';
import SQLEditor from 'SQLEditor';
import {CtlDataSourceOptions, CtlQuery, defaultQuery} from "./types";
import {selectors} from "./tests/selectors";
import {appendTemplateVariables} from "./utils";

type Props = QueryEditorProps<DataSource, CtlQuery, CtlDataSourceOptions> & {
  hideOptions?: boolean;
};

type QueryProperties = 'regions';

export function QueryEditorForm(props: Props) {
  const [] = useState(false);
  useEffect(() => {}, [props.datasource]);
  const queryWithDefaults = {
    ...defaultQuery,
    ...props.query,
    connectionArgs: {
      ...defaultQuery.connectionArgs,
      ...props.query.connectionArgs,
    },
  };

  const templateVariables = props.datasource.getVariables();

  const fetchRegions = () =>
    props.datasource.getRegions().then((regions) => appendTemplateVariables(templateVariables, regions));

  const onChange = (prop: QueryProperties) => (e: SelectableValue<string> | null) => {
    const newQuery = { ...props.query };
    const value = e?.value;
    switch (prop) {
      case 'regions':
        newQuery.connectionArgs = { ...newQuery.connectionArgs, region: value };
        break;
    }
    props.onChange(newQuery);
  };

  return (
    <>
      <InlineSegmentGroup>
        <div className="gf-form-group">
          <ResourceSelector
            onChange={onChange('regions')}
            fetch={fetchRegions}
            value={queryWithDefaults.connectionArgs.region ?? null}
            default={props.datasource.defaultRegion}
            label={selectors.components.ConfigEditor.region.input}
            data-testid={selectors.components.ConfigEditor.region.wrapper}
            labelWidth={11}
            className="width-12"
          />
        </div>

        <div style={{ minWidth: '400px', marginLeft: '10px', flex: 1 }}>
          <SQLEditor query={queryWithDefaults} onChange={props.onChange} datasource={props.datasource} />
        </div>
      </InlineSegmentGroup>
    </>
  );
}
