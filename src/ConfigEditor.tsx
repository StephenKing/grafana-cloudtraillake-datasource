import React, { useState } from 'react';
import { DataSourcePluginOptionsEditorProps, DataSourceSettings, SelectableValue } from '@grafana/data';
import { CtlDataSourceOptions, CtlDataSourceSecureJsonData, CtlDataSourceSettings } from './types';
import { getBackendSrv } from '@grafana/runtime';
import { ConfigSelect, ConnectionConfig } from '@grafana/aws-sdk';
import { selectors } from 'tests/selectors';

type Props = DataSourcePluginOptionsEditorProps<CtlDataSourceOptions, CtlDataSourceSecureJsonData>;

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
        const res: string[] = await getBackendSrv().get(resourcesURL + '/eventDataStores');
        return res.map((eds) => ({ label: eds, value: eds, id: eds }));
    };
    const onOptionsChange = (options: DataSourceSettings<CtlDataSourceOptions, CtlDataSourceSecureJsonData>) => {
        setSaved(false);
        props.onOptionsChange(options);
    };

    const onChangeEventDataStore = (e: SelectableValue<string> | null) => {
        const value = e?.value;
        props.onOptionsChange({
            ...props.options,
            jsonData: {
                ...props.options.jsonData,
                edsId: value
            },
        });
    };

    return (
        <div className="gf-form-group">
            <ConnectionConfig {...props} onOptionsChange={onOptionsChange} />
            <h3>CloudTrail Lake Details</h3>

            <ConfigSelect
                {...props}
                value={props.options.jsonData.edsId ?? ''}
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
