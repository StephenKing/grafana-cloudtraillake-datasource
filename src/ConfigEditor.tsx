import React, { useState } from 'react';
import { DataSourcePluginOptionsEditorProps, DataSourceSettings, SelectableValue } from '@grafana/data';
import { CtlDataSourceOptions, CtlDataSourceSecureJsonData, CtlDataSourceSettings } from './types';
import { getBackendSrv } from '@grafana/runtime';
import { ConfigSelect, ConnectionConfig } from '@grafana/aws-sdk';
import { selectors } from 'tests/selectors';

type Props = DataSourcePluginOptionsEditorProps<CtlDataSourceOptions, CtlDataSourceSecureJsonData>;

type EventDataStore = { name: string; id: string }

export function ConfigEditor(props: Props) {
    const baseURL = `/api/datasources/${props.options.id}`;
    const resourcesURL = `${baseURL}/resources`;
    const [saved, setSaved] = useState(!!props.options.jsonData.defaultRegion);
    const saveOptions = async () => {
        if (saved) {
            return;
        }
        await getBackendSrv()
            .put(baseURL, props.options)
            .then((result: { datasource: CtlDataSourceSettings }) => {
                props.onOptionsChange({
                    ...props.options,
                    version: result.datasource.version,
                });
            });
        setSaved(true);
    };

    const fetchEventDataStores = async () => {
        const res: EventDataStore[] = await getBackendSrv().get(resourcesURL + '/eventDataStores');
        return res.map((eds) => ({ label: eds.name , value: eds.id, description: eds.id }));
    };
    const onOptionsChange = (options: DataSourceSettings<CtlDataSourceOptions, CtlDataSourceSecureJsonData>) => {
        setSaved(false);
        props.onOptionsChange(options);
    };

    const onChangeEventDataStore = (e: SelectableValue<string> | null) => {
        const value = e?.value ?? '';
        const label = e?.label ?? '';
        props.onOptionsChange({
            ...props.options,
            jsonData: {
                ...props.options.jsonData,
                eventDataStore: { id: value, name: label },
            },
        });
    };

    return (
        <div className="gf-form-group">
            <ConnectionConfig {...props} onOptionsChange={onOptionsChange} />
            <h3>CloudTrail Lake Details</h3>

            <ConfigSelect
                {...props}
                value={props.options.jsonData.eventDataStore?.name ?? ''}
                onChange={onChangeEventDataStore}
                fetch={fetchEventDataStores}
                label={selectors.components.ConfigEditor.EventDataStore.input}
                isClearable={true}
                noOptionsMessage={"No Event Data Stores found"}
                data-testid={selectors.components.ConfigEditor.EventDataStore.wrapper}
                saveOptions={saveOptions}
            />
        </div>
    );
}
