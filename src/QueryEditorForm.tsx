import React, { useEffect, useState } from 'react';
import { QueryEditorProps, SelectableValue } from '@grafana/data';
import { DataSource } from './datasource';
import { InlineSegmentGroup } from '@grafana/ui';
import { ResourceSelector } from '@grafana/aws-sdk';
import SQLEditor from 'SQLEditor';
import {CtlDataSourceOptions, CtlQuery, defaultQuery} from "./types";
import {selectors} from "./tests/selectors";

type Props = QueryEditorProps<DataSource, CtlQuery, CtlDataSourceOptions> & {
  hideOptions?: boolean;
};

type QueryProperties = 'regions' |'edsId';

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

  const fetchRegions = async () => {
    const region: string[] = await props.datasource.getResource('regions');
    return region.map((region) =>  ({label: region, value: region}));
  }

  const fetchEventDatastores = async () => {
    const eds: string[] = await props.datasource.getResource('eventDataStores');
    return eds.map((eds) =>  ({label: eds, value: eds, id: eds}));
  }

  const onChange = (prop: QueryProperties) => (e: SelectableValue<string> | null) => {
    const newQuery = { ...props.query };
    const value = e?.value;
    switch (prop) {
      case 'regions':
        newQuery.connectionArgs = { ...newQuery.connectionArgs, region: value };
        break;
      case 'edsId':
        console.log("onChange, edsId", value)
        newQuery.connectionArgs = { ...newQuery.connectionArgs, edsId: value };
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
            labelWidth={20}
            className="width-16"
          />
          <ResourceSelector
              onChange={onChange('edsId')}
              fetch={fetchEventDatastores}
              value={queryWithDefaults.connectionArgs.edsId ??  null}
              default={props.datasource.defaultEdsId}
              tooltip="Use the selected EDS with the $__edsId macro"
              label={selectors.components.ConfigEditor.EventDataStore.input}
              data-testid={selectors.components.ConfigEditor.EventDataStore.wrapper}
              labelWidth={20}
              className="width-16"
          />
        </div>

        <div style={{ minWidth: '400px', marginLeft: '10px', flex: 1 }}>
          <SQLEditor query={queryWithDefaults} onChange={props.onChange} datasource={props.datasource} />
        </div>
      </InlineSegmentGroup>
    </>
  );
}
