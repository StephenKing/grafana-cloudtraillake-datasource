import React, { useState, FormEvent } from 'react';
import { DataSourcePluginOptionsEditorProps, DataSourceSettings, SelectableValue } from '@grafana/data';
import { CtlDataSourceOptions, CtlDataSourceSecureJsonData, CtlDataSourceSettings } from './types';
import { getBackendSrv } from '@grafana/runtime';
import { InlineInput, ConfigSelect, ConnectionConfig } from '@grafana/aws-sdk';
import { selectors } from 'tests/selectors';

type Props = DataSourcePluginOptionsEditorProps<CtlDataSourceOptions, CtlDataSourceSecureJsonData>;

export type ResourceType = 'catalog' | 'database' | 'workgroup';

export function ConfigEditor(props: Props) {
    const baseURL = `/api/datasources/${props.options.id}`;
    // const resourcesURL = `${baseURL}/resources`;
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

    const fetchTestOption2 = async () => {
        const test2: string[] = ["foo", "bar"]
        return test2;
    };
    const onOptionsChange = (options: DataSourceSettings<CtlDataSourceOptions, CtlDataSourceSecureJsonData>) => {
        setSaved(false);
        props.onOptionsChange(options);
    };

    const onChange = (resource: ResourceType) => (e: SelectableValue<string> | null) => {
        const value = e?.value ?? '';
        props.onOptionsChange({
            ...props.options,
            jsonData: {
                ...props.options.jsonData,
                [resource]: value,
            },
        });
    };

    const onChangeTestOption = (e: FormEvent<HTMLInputElement>) => {
        const value = e.currentTarget.value;
        props.onOptionsChange({
            ...props.options,
            jsonData: {
                ...props.options.jsonData,
                testOption: value,
            },
        });
    };

    return (
        <div className="gf-form-group">
            <ConnectionConfig {...props} onOptionsChange={onOptionsChange} />
            <h3>Athena Details</h3>

            <ConfigSelect
                {...props}
                value={props.options.jsonData.testOption2 ?? ''}
                onChange={onChange('workgroup')}
                fetch={fetchTestOption2}
                label={selectors.components.ConfigEditor.testOption2.input}
                data-testid={selectors.components.ConfigEditor.testOption2.wrapper}
                saveOptions={saveOptions}
            />
            <InlineInput
                {...props}
                value={props.options.jsonData.testOption ?? ''}
                onChange={onChangeTestOption}
                label={selectors.components.ConfigEditor.testOption.input}
                data-testid={selectors.components.ConfigEditor.testOption.wrapper}
                tooltip="Optional. Test"
                placeholder="example"
            />
        </div>
    );
}
